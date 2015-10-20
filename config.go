package billforward

import (
	"time"
)

const (
	EndpointSandbox    = "https://app-sandbox.billforward.net/"
	EndpointProduction = "https://app.billforward.net/"
)

type Config struct {
	// Full URL to BillForward API
	Endpoint string

	// Transport, if nil, DefaultTransport will be used.
	Transport Transport

	// Secret Access Token.
	AccessToken string

	// String to appended to User-Agent header.
	AppendUserAgent string

	// Timeout on a per-request basis.
	TimeoutPerRequest time.Duration

	// Number times to attempt a request, if 0, defaults to 5.
	Attempts int

	// Print ever request to stderr
	PrintCurlDebug bool
}
