package billforward

import (
	"github.com/authclub/billforward/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestListProducts(t *testing.T) {
	tc := testClient(t)
	require.NotNil(t, tc)

	acct := tc.Accounts()

	profile := &types.Profile{
		FirstName: "Alice",
		LastName:  "Example",
		Email:     "alice.example@tests.billforward.net",
		Dob:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}
	account, err := acct.CreateAccount(profile)
	require.NotNil(t, account)
	require.NoError(t, err)
	require.NotEmpty(t, account.Id)

	defer cleanupAccount(account)

	products := tc.Products()
	called := false
	err = products.ListProducts(nil, func(i *types.Product) (bool, error) {
		called = true
		return false, nil
	})
	require.NoError(t, err)
	require.False(t, called)

}
