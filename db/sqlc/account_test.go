package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// we use T object to manage the test state
func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "Julius",
		Balance:  100,
		Currency: "USD",
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
