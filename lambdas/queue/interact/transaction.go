package main

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tina/packages/config"
	"tina/packages/db"
	"tina/packages/nexus"
	"tina/packages/openai"
	"tina/packages/types"

	"github.com/google/uuid"
)

func HandleTransactionEvent(ctx context.Context, event *types.TransactionEvent) error {
	switch event.Type {
	case "REQUEST":
		return HandleRequest(ctx, event)
	case "SEND":
		return HandleSend(ctx, event)
	case "ACCEPT":
		return HandleAccept(ctx, event)
	case "REJECT":
		return HandleReject(ctx, event)
	default:
		return fmt.Errorf("unknown transaction event type %s", event.Type)
	}
}

func HandleRequest(ctx context.Context, event *types.TransactionEvent) error {
	fmt.Println("request")
	return nil
}

func HandleSend(ctx context.Context, event *types.TransactionEvent) error {
	fmt.Println("send")
	transaction := &event.Transaction
	if transaction.SourceUsername == config.GetUsername() {
		// for now we just handle direct payment tranfers from users
		return nil
	}

	err := HandleSuccessTransaction(ctx, transaction)
	if err != nil {
		return err
	}

	metaStr := "name: tina, salimon is a network where you are providing service for users"
	contentStr := fmt.Sprintf("%s sent you %d credits for following reason:\n%s\n\nplease provide your feedback here.\n",
		event.Transaction.SourceUsername,
		event.Transaction.Amount,
		event.Transaction.Description)
	body := fmt.Sprintf("%s\n%s", metaStr, contentStr)
	return SendMesssageToUser(event.Transaction.SourceId, body)
}

func HandleAccept(ctx context.Context, event *types.TransactionEvent) error {
	fmt.Println("accept")
	transaction := &event.Transaction
	if transaction.SourceUsername == config.GetUsername() {
		// for now we just handle direct payment tranfers from users
		return nil
	}

	err := HandleSuccessTransaction(ctx, transaction)
	if err != nil {
		return err
	}

	metaStr := "name: tina, salimon is a network where you are providing service for users"
	contentStr := fmt.Sprintf("%s accepted your credit request with amount %d credits. your transaction description was:\n%s\n\nplease provide your feedback here.\n",
		event.Transaction.SourceUsername,
		event.Transaction.Amount,
		event.Transaction.Description)
	body := fmt.Sprintf("%s\n%s", metaStr, contentStr)
	return SendMesssageToUser(event.Transaction.SourceId, body)
}

func HandleReject(ctx context.Context, event *types.TransactionEvent) error {
	fmt.Println("reject")

	transaction := &event.Transaction
	if transaction.SourceUsername == config.GetUsername() {
		// for now we just handle direct payment tranfers from users
		return nil
	}
	metaStr := "name: tina, salimon is a network where you are providing service for users"
	contentStr := fmt.Sprintf("%s rejected your transaction request with amount %d credits. your transaction description was:\n%s\n\nplease provide your feedback here.\n",
		event.Transaction.SourceUsername,
		event.Transaction.Amount,
		event.Transaction.Description)
	body := fmt.Sprintf("%s\n%s", metaStr, contentStr)
	return SendMesssageToUser(event.Transaction.SourceId, body)
}

func GetResponse(body string) (string, error) {
	messages := []types.Message{
		{
			Id:        uuid.New(),
			Body:      body,
			Type:      types.MessageTypeText,
			UserId:    uuid.New(),
			Username:  "username",
			ThreadId:  uuid.New(),
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}
	response, err := openai.SendCompletionRequest(messages)
	if err != nil {
		return "", err
	}
	return response.Body, nil
}

func SendMesssageToUser(userId uuid.UUID, message string) error {
	responseBody, err := GetResponse(message)
	if err != nil {
		return err
	}
	// start a payment thread with user
	thread, err := nexus.StartThread(userId.String(), "", types.ThreadCategroyPayment)
	if err != nil {
		return err
	}
	return nexus.SendMessage(thread.Id.String(), responseBody, types.MessageTypeText)
}

// if the transaction is successful, entity will manage the linked invoice or add it up to user credit
func HandleSuccessTransaction(ctx context.Context, tx *types.Transaction) error {
	// fetch user info first
	user, err := db.FindUser("nexus_id = ?", tx.SourceId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("no such user")
	}
	invoice, err := db.FindInvoice("transaction_id = ? AND status = ?", tx.Id, types.TransactionStatusTypePending)
	if err != nil {
		return err
	}
	if invoice != nil {
		// there is an invoice for it
		invoice.Status = types.TransactionStatusTypeDone
		result := db.DB.Save(invoice)
		if result.Error != nil {
			return result.Error
		}
	}

	paidUsage := types.UsageFromCredit(tx.Amount)
	if user.Usage < paidUsage {
		// user paid more than what he used so entity stroes extra credit
		extraUsage := paidUsage - user.Usage
		extraCredit := types.CreditFromUsage(extraUsage)
		user.Credit += extraCredit
		user.Usage = 0
	} else {
		user.Usage -= paidUsage
	}
	result := db.DB.Save(user)
	return result.Error
}
