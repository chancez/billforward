package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/authclub/billforward/client/accounts"

	"github.com/authclub/billforward/client"

	"github.com/go-swagger/go-swagger/httpkit"
	httpclient "github.com/go-swagger/go-swagger/httpkit/client"
)

type ReadWrapper struct {
	*bytes.Reader
}

func (r *ReadWrapper) Close() error {
	return nil
}

func main() {
	transport := httpclient.New("api-sandbox.billforward.net:443", "/v1", []string{"https"})
	transport.DefaultAuthentication = httpclient.BearerToken(os.Getenv("BILLFORWARD_API_KEY"))
	// Only necessary for the PDF response, but here in case someone needs it
	transport.Consumers["application/pdf"] = httpkit.ConsumerFunc(func(r io.Reader, data interface{}) error {
		f := data.(*httpkit.File)
		b, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		f.Data = &ReadWrapper{bytes.NewReader(b)}
		return nil
	})
	bfClient := client.New(transport, nil)

	acctsResp, err := bfClient.Accounts.GetAllAccounts(accounts.NewGetAllAccountsParams())
	if err != nil {
		log.Fatal(err)
	}
	for i, acct := range acctsResp.Payload.Results {
		log.Printf("account %d: %+v\n", i, acct)
		log.Printf("profile %d: %+v\n", i, acct.Profile)
	}

}
