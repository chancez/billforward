package billforward

import (
	"github.com/authclub/billforward/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestListInvoicesInvalidAccountId(t *testing.T) {
	tc := testClient(t)
	require.NotNil(t, tc)

	invoices := tc.Invoices()
	called := false
	err := invoices.ListInvoicesByAccountId("invalid-id", func(i *types.Invoice) (bool, error) {
		called = true
		return false, nil
	})

	if false {
		// TODO(pquerna): this should error, but they don't appear to validate account Ids and return an empty list instead :(
		require.Error(t, err)
	}
	require.False(t, called)
}

func TestListInvoices(t *testing.T) {
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

	invoices := tc.Invoices()
	called := false
	err = invoices.ListInvoicesByAccountId(account.Id, func(i *types.Invoice) (bool, error) {
		called = true
		return false, nil
	})
	require.NoError(t, err)
	require.False(t, called)

}
