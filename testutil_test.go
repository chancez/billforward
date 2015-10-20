package billforward

import (
	"github.com/stretchr/testify/require"

	"os"
	"testing"
)

func testClient(t *testing.T) Client {
	ep := os.Getenv("BILLFORWARD_TEST_ENDPOINT")
	if ep == "" {
		ep = EndpointSandbox
	}
	if ep == EndpointProduction {
		panic("test attempted to use production endpoint")
	}

	at := os.Getenv("BILLFORWARD_TEST_ACCESS_TOKEN")
	if at == "" {
		t.Skip("BILLFORWARD_TEST_ACCESS_TOKEN is unset, skipping....")
		t.SkipNow()
		return nil
	}
	c, err := New(Config{
		Endpoint:       ep,
		AccessToken:    at,
		PrintCurlDebug: true,
	})

	require.NoError(t, err)
	require.NotNil(t, c)

	return c
}
