package main

import (
	"context"
	"tina/packages/db"
	"tina/packages/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// migrate db schemas
	db.DB.AutoMigrate(types.User{})
	return events.APIGatewayV2HTTPResponse{StatusCode: 200, Body: "migration complete"}, nil
}

func main() {
	godotenv.Load("/opt/.env")
	db.SetupDatabase()
	lambda.Start(handler)
}
