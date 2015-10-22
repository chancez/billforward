package billforward

import (
	"github.com/authclub/billforward/types"
)

const productsListUrl = "/v1/products"

func (c *client) ListProducts(filter *types.ProductsFilter, fn ProductIterateFn) error {
	url := productsListUrl

	// TODO(pquerna): query args library?
	if filter == nil {
		filter = &types.ProductsFilter{}
	}

	url += "?"
	if filter.IncludeRetired {
		url += "&include_retired=true"
	} else {
		url += "&include_retired=false"
	}

	for {
		var products types.ProductsResponse

		hr, err := c.req("GET", url, nil)
		if err != nil {
			return err
		}

		err = c.rt(hr, &products)
		if err != nil {
			return err
		}

		for _, a := range products.Results {
			next, err := fn(a)
			if err != nil {
				return err
			}

			if !next {
				return nil
			}
		}

		if products.NextPage == "" {
			break
		}

		if products.RecordsRequested == products.RecordsReturned {
			url = "/v1" + products.NextPage
			continue
		} else {
			break
		}
	}
	return nil
}
