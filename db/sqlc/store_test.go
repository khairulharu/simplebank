package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	arg := TransferTxParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        100,
	}

	transfer, err := testStore.TransferTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	fromEntry, err := testQueries.GetEntryByID(context.Background(), fromAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fromEntry)

	toEntry, err := testQueries.GetEntryByID(context.Background(), toAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, toEntry)

	require.Equal(t, fromEntry.AccountID, fromAccount.ID)
	require.Equal(t, toEntry.AccountID, toAccount.ID)
	require.Equal(t, fromEntry.Amount, -arg.Amount)
	require.Equal(t, toEntry.Amount, arg.Amount)
}
