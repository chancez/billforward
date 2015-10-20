package billforward

import (
	"fmt"
	"github.com/authclub/billforward/types"
)

const accountCreationUrl = "/v1/accounts"

type accountCreationBody struct {
	Profile *types.Profile `json:"profile"`
}

func (c *client) CreateAccount(profile *types.Profile) (*types.Account, error) {
	var account types.AccountsResponse

	hr, err := c.req("POST", accountCreationUrl,
		&accountCreationBody{profile})
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
const accountGetByIdUrl = "/v1/accounts/%s"

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

const accountDeleteByIdUrl = "/v1/accounts/%s"

func (c *client) DeleteAccountById(accountId string) error {
	hr, err := c.req("DELETE", fmt.Sprintf(accountDeleteByIdUrl, accountId), nil)
	if err != nil {
		return err
	}

	err = c.rt(hr, nil)
	if err != nil {
		return err
	}

	return nil
}

const accountListUrl = "/v1/accounts"

func (c *client) ListAccounts(filter *types.AccountsFilter, fn AccountIterateFn) error {
	url := accountListUrl

	// TODO(pquerna): query args library?
	if filter == nil {
		filter = &types.AccountsFilter{}
	}

	url += "?"
	if filter.IncludeRetired {
		url += "&include_retired=true"
	} else {
		url += "&include_retired=false"
	}

	for {
		var account types.AccountsResponse

		hr, err := c.req("GET", url, nil)
		if err != nil {
			return err
		}

		err = c.rt(hr, &account)
		if err != nil {
			return err
		}

		for _, a := range account.Results {
			next, err := fn(a)
			if err != nil {
				return err
			}

			if !next {
				return nil
			}
		}

		if account.NextPage == "" {
			break
		}

		if account.RecordsRequested == account.RecordsReturned {
			url = "/v1" + account.NextPage
			continue
		} else {
			break
		}
	}
	return nil
}
