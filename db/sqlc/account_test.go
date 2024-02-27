package db

import (
	"context"
	"testing"

	"github.com/mahdikarami0111/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomString(6),
		Currency: util.RandomString(3),
		Balance:  util.RandomInt(50, 1000),
	}
	acc, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Currency, acc.Currency)
	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomString(6),
		Currency: util.RandomString(3),
		Balance:  util.RandomInt(50, 1000),
	}
	acc, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Currency, acc.Currency)
	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}

func TestGetAccount(t *testing.T) {
	var id int64
	id = 17
	account, err := testQueries.GetAccount(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, account)
}
