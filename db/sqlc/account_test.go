package db

import (
	"context"
	"testing"
	"time"

	"github.com/chavocito/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomAmount(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.Balance)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NotEmpty(t, account1, "Account1 is not an empty object")
	require.NotEmpty(t, account2, "Account2 is not an empty object")
	require.NoError(t, err, "The account creation query does not throw any errors")
	require.Equal(t, account1, account2, "Accounts 1 and 2 are the same object")
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdatedAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	updateOwner := util.RandomOwner()
	updateAmount := util.RandomAmount()

	arg := UpdateAccountsParams{
		ID:      account1.ID,
		Owner:   updateOwner,
		Balance: updateAmount,
	}

	err := testQueries.UpdateAccounts(context.Background(), arg)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEqual(t, account1, account2)
	require.NotEmpty(t, account1)
	require.NotEqual(t, account1.Owner, account2.Owner)
	require.NotEqual(t, account1.Balance, account2.Balance)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.NotEqual(t, account1, account2)
	require.Empty(t, account2)
	require.NotEqual(t, account1.Balance, account2.Balance)
}
