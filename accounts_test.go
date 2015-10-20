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
	require.Nil(t, account)
	require.Error(t, err)
}
