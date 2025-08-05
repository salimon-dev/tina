package main

import (
	"context"
	"fmt"
	"time"
	"tina/packages/nexus"
	"tina/packages/openai"
	"tina/packages/types"

	"github.com/google/uuid"
)

func HandleTransactionEvent(ctx context.Context, event *types.TransactionEvent) error {
	message := fmt.Sprintf("user %s sent you %d tokens credit over salimon network with category %s and description: %s\n thank from user in a new message", event.Transaction.SourceUsername, event.Transaction.Amount, event.Transaction.Category, event.Transaction.Description)
	messages := []types.Message{
		{
			Id:        uuid.New(),
			Body:      message,
			Type:      types.MessageTypeText,
			UserId:    event.Transaction.SourceId,
			Username:  event.Transaction.SourceUsername,
			ThreadId:  uuid.New(),
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}
	response, err := openai.SendCompletionRequest(messages)
	if err != nil {
		return err
	}
	fmt.Println(response.Body)

	// start a payment thread with user
	thread, err := nexus.StartThread(event.Transaction.SourceId.String(), "", types.ThreadCategroyPayment)
	if err != nil {
		return err
	}
	err = nexus.SendPlainMessage(thread.Id.String(), response.Body)
	return err
}
