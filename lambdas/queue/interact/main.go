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
		var event types.QueueEventInteract
		err := json.Unmarshal([]byte(record.Body), &event)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch event.Type {
		case types.QueueEventInteractTypeThreadUpdate:
			err = HandleThreadUpdateEvent(ctx, &event)
		case types.QueueEventInteractTypeTransaction:
			err = HandleTransactionEvent(ctx, &event)
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
