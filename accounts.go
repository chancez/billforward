package billforward

import (
	"fmt"
	"github.com/authclub/billforward/types"
)

const accountCreationUrl = "v1/accounts"

func (c *client) CreateAccount(profile *types.Profile) (*types.Account, error) {
	var account types.AccountsResponse

	hr, err := c.req("POST", accountCreationUrl, profile)
	if err != nil {
		return nil, err
	}

	err = c.rt(hr, &account)
	if err != nil {
		return nil, err
	}

	if len(account.Results) != 1 {
		return nil, fmt.Errorf("CreateAccount: returned more than 1 result: %d", len(account.Results))

	}

	return account.Results[0], nil
}

// TOOD(pquerna): named parameters?
const accountGetByIdUrl = "v1/accounts/%s"

func (c *client) GetAccountById(accountId string) (*types.Account, error) {
	var account types.AccountsResponse

	hr, err := c.req("GET", fmt.Sprintf(accountGetByIdUrl, accountId), nil)
	if err != nil {
		return nil, err
	}

	err = c.rt(hr, &account)
	if err != nil {
		return nil, err
	}

	if len(account.Results) != 1 {
		return nil, fmt.Errorf("GetAccountById: returned more than 1 result: %d", len(account.Results))
	}

	return account.Results[0], nil
}
