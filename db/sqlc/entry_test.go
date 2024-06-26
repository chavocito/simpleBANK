package db

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var entryArg = CreateEntriesParams{
	AccountID: 1,
	Amount:    200,
}

func TestDeleteEntry(t *testing.T) {
	var entry Entry
	var ctx = context.Background()

	entries, err := testQueries.ListEntries(ctx)
	if err != nil {
		log.Fatal("error getting accounts list:", err)
	}
	for _, ent := range entries {
		if ent.AccountID == entryArg.AccountID {
			entry = ent
		}
	}

	err = testQueries.DeleteAccount(ctx, entry.ID)
	if err != nil {
		return
	}

	entry, err = testQueries.GetEntry(ctx, entry.ID)

	//require.Empty(t, entry)
	require.Empty(t, err)
	//require.Zero(t, entry.Amount)
	//require.Empty(t, entry.AccountID)
}
