package db

import (
	"context"
	"log"
	"testing"

	"github.com/chavocito/simple_bank/util"
	"github.com/stretchr/testify/require"
)

var arg = CreateAccountParams{
	Owner:    util.RandomOwner(),
	Balance:  util.RandomAmount(),
	Currency: util.RandomCurrency(),
}

func TestCreateAccount(t *testing.T) {
	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.Balance)
	require.NotZero(t, account.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	var account Account
	var ctx = context.Background()

	accounts, err := testQueries.ListAccounts(ctx)
	if err != nil {
		log.Fatal("error getting accounts list:", err)
	}
	for _, acc := range accounts {
		if acc.Owner == arg.Owner {
			account = acc
		}
	}

	err = testQueries.DeleteAccount(ctx, account.ID)
	if err != nil {
		return
	}

	acc, err := testQueries.GetAccount(ctx, account.ID)

	require.Empty(t, acc)
	require.Zero(t, acc.Balance)
	require.Empty(t, acc.Owner)
	require.Empty(t, acc.Currency)
}
