package main

import (
	"context"
	"fmt"
	"tina/packages/types"
)

func HandleTransactionEvent(ctx context.Context, event *types.TransactionEvent) error {
	message := fmt.Sprintf("I %s sent you %d tokens credit over salimon network with category %s and description: %s", event.Transaction.SourceUsername, event.Transaction.Amount, event.Transaction.Category)
	fmt.Println(message)
	return nil
}
