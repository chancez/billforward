package billforward

import (
	"fmt"
	"github.com/authclub/billforward/types"
)

const invoiceListByAccountUrl = "/v1/invoices/account/%s"

func (c *client) ListInvoicesByAccountId(accountId string, fn InvoiceIterateFn) error {
	url := fmt.Sprintf(invoiceListByAccountUrl, accountId)

	for {
		var invoices types.InvoicesResponse

		hr, err := c.req("GET", url, nil)
		if err != nil {
			return err
		}

		err = c.rt(hr, &invoices)
		if err != nil {
			return err
		}

		for _, a := range invoices.Results {
			next, err := fn(a)
			if err != nil {
				return err
			}

			if !next {
				return nil
			}
		}

		if invoices.NextPage == "" {
			break
		}

		if invoices.RecordsRequested == invoices.RecordsReturned {
			url = "/v1" + invoices.NextPage
			continue
		} else {
			break
		}
	}
	return nil
}
