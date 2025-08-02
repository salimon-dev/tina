package main

import (
	"context"
	"fmt"
	"tina/packages/types"
)

func HandleTransactionEvent(ctx context.Context, event *types.QueueEventInteract) error {
	fmt.Printf("new transaction event with id %s", event.TransactionId)
	return nil
}
