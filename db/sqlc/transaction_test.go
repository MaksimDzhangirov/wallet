package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/MaksimDzhangirov/wallet/util"
)

func createRandomTransaction(t *testing.T, account Account) Transaction {
	arg := CreateTransactionParams{
		AccountID:   account.ID,
		Amount:      util.RandomMoney(),
		Description: util.RandomString(50),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.AccountID, transaction.AccountID)
	require.Equal(t, arg.Amount, transaction.Amount)

	require.NotZero(t, transaction.ID)
	require.NotZero(t, transaction.CreatedAt)

	return transaction
}

func TestCreateTransaction(t *testing.T) {
	account := createRandomAccount(t)
	createRandomTransaction(t, account)
}

func TestGetTransaction(t *testing.T) {
	account := createRandomAccount(t)
	transaction1 := createRandomTransaction(t, account)
	transaction2, err := testQueries.GetTransaction(context.Background(), transaction1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transaction2)

	require.Equal(t, transaction1.ID, transaction2.ID)
	require.Equal(t, transaction1.AccountID, transaction2.AccountID)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
	require.WithinDuration(t, transaction1.CreatedAt, transaction2.CreatedAt, time.Second)
}

