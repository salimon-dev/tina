package main

import (
	"context"
	"fmt"
	"tina/packages/types"
)

func HandleTransactionEvent(ctx context.Context, event *types.TransactionEvent) error {
	fmt.Printf("new transaction event with id %s", event.Transaction.Id)
	return nil
}
