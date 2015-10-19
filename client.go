package billforward

import (
	"github.com/authclub/billforward/types"
	"golang.org/x/net/context"

	"net/url"
	// TOOD(pquerna): remove fmt.Errorf, use correct
	"fmt"
)

func New(c Config) (Client, error) {
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, err
	}

	if u.IsAbs() {
		return nil, fmt.Errorf("config.Endpoint must be an absolute url: %s", c.Endpoint)
	}

	config := Config(c)
	config.Endpoint = u.String()

	return &client{
		config: c,
		// TODO(pquerna): should we allow the client to pass the context?
		ctx: context.Background(),
	}, nil
}

// Each major sub-system is a separate Interface which makes it easier to mock in test cases.
type Client interface {
	Accounts() AccountClient
	Invoices() InvoiceClient

	httpClientDo
}

type AccountClient interface {
	CreateAccount(profile *types.Profile) (*types.Account, error)
	GetAccountById(accountId string) (*types.Account, error)
}

type InvoiceClient interface {
	//	FetchInvoiceByAccount(accountId string) ([]*types.Invoice, error)
}

type client struct {
	config Config
	ctx    context.Context
}

func (c *client) Accounts() AccountClient {
	return c
}

func (c *client) Invoices() InvoiceClient {
	return c
}
