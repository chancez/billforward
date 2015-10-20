package billforward

import (
	"github.com/authclub/billforward/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestAccountCreate(t *testing.T) {
	tc := testClient(t)
	require.NotNil(t, tc)

	acct := tc.Accounts()
	require.NotNil(t, acct)

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

	account2, err := acct.GetAccountById(account.Id)
	require.NoError(t, err)
	require.NotNil(t, account)
	require.Equal(t, account.Id, account2.Id)
	require.Equal(t, account.Profile.FirstName, account2.Profile.FirstName)

	err = acct.DeleteAccountById(account.Id)
	require.NoError(t, err)
}

func TestAccountsList(t *testing.T) {
	tc := testClient(t)
	require.NotNil(t, tc)

	acct := tc.Accounts()
	require.NotNil(t, acct)

	called := false
	err := acct.ListAccounts(func(a *types.Account) (bool, error) {
		called = true
		require.NotNil(t, called)
		return true, nil
	})
	require.True(t, called)
	require.NoError(t, err)
}
