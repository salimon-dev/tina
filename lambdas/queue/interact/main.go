package main

import (
	"context"
	"encoding/json"
	"fmt"
	"tina/packages/db"
	"tina/packages/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func handler(ctx context.Context, request events.SQSEvent) error {
	for _, record := range request.Records {
		fmt.Printf("SQS message body: %s\n", record.Body)
		var event types.BaseEvent
		err := json.Unmarshal([]byte(record.Body), &event)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch event.Action {
		// case "THREAD":
		// 	var threadEvent types.ThreadEvent
		// 	err = json.Unmarshal([]byte(record.Body), &threadEvent)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		continue
		// 	}
		// 	err = HandleThreadEvent(ctx, &threadEvent)
		case "MESSAGE":
			var messageEvent types.MessageEvent
			err = json.Unmarshal([]byte(record.Body), &messageEvent)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = HandleMessageEvent(ctx, &messageEvent)
		case "TRANSACTION":
			var transactionEvent types.TransactionEvent
			err = json.Unmarshal([]byte(record.Body), &transactionEvent)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = HandleTransactionEvent(ctx, &transactionEvent)
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return nil
}

func main() {
	godotenv.Load("/opt/.env")
	db.SetupDatabase()
	lambda.Start(handler)
}
