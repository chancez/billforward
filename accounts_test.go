package billforward

import (
	"github.com/authclub/billforward/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAccountCreate(t *testing.T) {
	tc := testClient(t)
	require.NotNil(t, tc)

	acct := tc.Accounts()
	require.NotNil(t, acct)

	profile := &types.Profile{}
	account, err := acct.CreateAccount(profile)
	require.Nil(t, account)
	require.Error(t, err)
}
