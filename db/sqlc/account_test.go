package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
	"simplebank/util"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, newAccount.ID, account.ID)
	require.Equal(t, newAccount.Owner, account.Owner)
	require.Equal(t, newAccount.Balance, account.Balance)
	require.Equal(t, newAccount.Currency, account.Currency)
	require.WithinDuration(t, newAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	var newAccounts []Account
	for i := 0; i < 10; i++ {
		newAccounts = append(newAccounts, createRandomAccount(t))
	}
	arg := ListAccountsParams{
		Owner:  newAccounts[0].Owner,
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	for _, a := range accounts {
		require.NotEmpty(t, a)
		require.Equal(t, newAccounts[0].Owner, a.Owner)
	}
}

func TestUpdateAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	arg := AddAccountBalanceParams{ID: newAccount.ID, Amount: util.RandomMoney()}
	account, err := testQueries.AddAccountBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, newAccount.ID, account.ID)
	require.Equal(t, newAccount.Owner, account.Owner)
	require.Equal(t, newAccount.Balance+arg.Amount, account.Balance)
	require.Equal(t, newAccount.Currency, account.Currency)
	require.WithinDuration(t, newAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)

	account, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}
