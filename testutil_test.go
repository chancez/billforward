package billforward

import (
	"github.com/authclub/billforward/types"
	"github.com/stretchr/testify/require"

	"os"
	"testing"
)

func getTestClient() (Client, error) {
	ep := os.Getenv("BILLFORWARD_TEST_ENDPOINT")
	if ep == "" {
		ep = EndpointSandbox
	}
	if ep == EndpointProduction {
		panic("test attempted to use production endpoint")
	}

	at := os.Getenv("BILLFORWARD_TEST_ACCESS_TOKEN")
	if at == "" {
		return nil, nil
	}
	c, err := New(Config{
		Endpoint:       ep,
		AccessToken:    at,
		PrintCurlDebug: true,
	})

	return c, err
}

func testClient(t *testing.T) Client {
	c, err := getTestClient()
	require.NoError(t, err)
	if c == nil {
		t.Skip("BILLFORWARD_TEST_ACCESS_TOKEN is unset, skipping....")
		t.SkipNow()
		return nil
	}
	require.NotNil(t, c)
	return c
}

func cleanupAccount(account *types.Account) {
	c, _ := getTestClient()
	if c != nil {
		acct := c.Accounts()
		acct.RetireAccountById(account.Id)
	}
}
