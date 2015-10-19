package billforward

import (
	"golang.org/x/net/context"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Transport interface {
	http.RoundTripper
	CancelRequest(req *http.Request)
}

var DefaultTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
}

type httpAction interface {
	HTTPRequest() *http.Request
}

type httpClientDo interface {
	Do(context.Context, httpAction) (*http.Response, []byte, error)
}

type httpAct struct {
	req *http.Request
}

func (act *httpAct) HTTPRequest() *http.Request {
	return act.req
}

func (c *client) req(method string, path string, body interface{}) (httpAction, error) {
	var bodyBytes []byte
	var err error

	if body != nil {
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	url := c.config.Endpoint
	req, err := http.NewRequest(method, url, bytes.NewReader(bodyBytes))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Cotent-Type", "application/json; charset=utf-8")

	if c.config.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.AccessToken)
	}

	return &httpAct{
		req: req,
	}, nil
}

func (c *client) rt(action httpAction, respObj interface{}) error {
	resp, bodyBytes, err := c.Do(c.ctx, action)

	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		err = json.Unmarshal(bodyBytes, respObj)
		if err != nil {
			// TODO(pquerna): wrap json errors, they are worthless.
			return err
		}
	} else {
		return c.handleApiError(action, resp, bodyBytes)
	}

	return nil
}

type remoteError struct {
	Method  string
	URL     string
	Code    int
	Type    string `json:"type"`
	Message string `json:"message"`
}

// TODO(pquerna): improve
func (e remoteError) Error() string {
	return e.Message
}

func (c *client) handleApiError(action httpAction, resp *http.Response, body []byte) error {
	req := action.HTTPRequest()
	rerr := remoteError{}

	err := json.Unmarshal(body, &rerr)
	if err != nil {
		return fmt.Errorf("Received malformed %d error from %s to %s:\n%s\n", resp.StatusCode, req.Method, req.RequestURI, string(body))
	}

	if rerr.Type == "" {
		return fmt.Errorf("Received malformed %d error from %s to %s:\n%s\n", resp.StatusCode, req.Method, req.RequestURI, string(body))
	}

	rerr.Method = req.Method
	rerr.URL = req.RequestURI
	rerr.Code = resp.StatusCode

	return rerr
}

type roundTripResponse struct {
	resp *http.Response
	err  error
}

func (c *client) attempt(ctx context.Context, transport Transport, req *http.Request) (*http.Response, []byte, error) {
	// based on etcd-client's simpleHTTPClient.Do():
	// 	<https://github.com/coreos/etcd/blob/master/client/client.go#L374>

	hctx, hcancel := context.WithCancel(ctx)
	if c.config.TimeoutPerRequest > 0 {
		hctx, hcancel = context.WithTimeout(ctx, c.config.TimeoutPerRequest)
	}
	defer hcancel()

	reqcancel := requestCanceler(transport, req)

	rtchan := make(chan roundTripResponse, 1)
	go func() {
		resp, err := transport.RoundTrip(req)
		rtchan <- roundTripResponse{resp: resp, err: err}
		close(rtchan)
	}()

	var resp *http.Response
	var err error

	select {
	case rtresp := <-rtchan:
		resp, err = rtresp.resp, rtresp.err
	case <-hctx.Done():
		// cancel and wait for request to actually exit before continuing
		reqcancel()
		rtresp := <-rtchan
		resp = rtresp.resp
		switch {
		case ctx.Err() != nil:
			err = ctx.Err()
		case hctx.Err() != nil:
			err = fmt.Errorf("client: endpoint %s exceeded header timeout", c.config.Endpoint)
		default:
			panic("failed to get error from context")
		}
	}

	// race conditions between channels above
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {
		return nil, nil, err
	}

	var body []byte
	done := make(chan struct{})
	go func() {
		body, err = ioutil.ReadAll(resp.Body)
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		resp.Body.Close()
		<-done
		return nil, nil, ctx.Err()
	case <-done:
	}

	return resp, body, err
}

func (c *client) Do(ctx context.Context, action httpAction) (*http.Response, []byte, error) {
	attempts := c.config.Attempts
	if attempts == 0 {
		attempts = 5
	}

	// TOOD(pquerna): better ua
	ua := "billforward-go/0.1.0"
	if c.config.AppendUserAgent != "" {
		ua = ua + " " + c.config.AppendUserAgent
	}

	transport := c.config.Transport
	if transport == nil {
		transport = DefaultTransport
	}

	var resp *http.Response
	var err error
	var body []byte

	for i := 0; i < attempts; i++ {
		req := action.HTTPRequest()
		req.Header.Set("User-Agent", ua)
		resp, body, err = c.attempt(ctx, transport, req)
		if err != nil {
			return nil, nil, err
		}

		if resp.StatusCode >= 500 {
			continue
		}

		return resp, body, nil
	}

	if err != nil {
		return nil, nil, err
	}

	return nil, nil, fmt.Errorf("client: reached max retries: %d", attempts)
}
