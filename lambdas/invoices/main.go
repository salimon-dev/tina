package main

import (
	"context"
	"fmt"
	"tina/packages/db"
	"tina/packages/nexus"
	"tina/packages/types"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func handler(ctx context.Context) error {
	var users []types.User
	result := db.UsersModel().Where("usage > 0").Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	for _, user := range users {
		fmt.Printf("processing user %s", user.Username)
		err := HandleSingleUser(user)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func main() {
	godotenv.Load("/opt/.env")
	db.SetupDatabase()
	lambda.Start(handler)
}

func HandleSingleUser(user types.User) error {
	if user.Usage < user.UsageSoftLimit {
		fmt.Printf("user %s has not reached soft limit\n", user.Username)
		return nil
	}
	usageCredit := types.CreditFromUsage(user.Usage)
	// deduct what credit user already have with tina
	usageCredit -= user.Credit
	if usageCredit <= 0 {
		// user has enough credit to settle the usage. no need to create a new invoice
		user.Credit -= usageCredit
		user.Usage = 0
		result := db.DB.Save(&user)
		if result.Error != nil {
			fmt.Println(result.Error)
		}
		return nil
	} else {
		// spend credit from user and invoice the remaining debt
		user.Credit = 0
		result := db.DB.Save(&user)
		if result.Error != nil {
			fmt.Println(result.Error)
			return nil
		}
	}
	invoice, err := db.FindInvoice("user_id = ? AND status = ?", user.Id, types.TransactionStatusTypePending)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if invoice != nil {
		// there is an active invoice already submitted
		return nil
	}
	SubmitInvoice(usageCredit, &user)
	return nil
}

func SubmitInvoice(usageCredit uint64, user *types.User) {
	// send the invoice transaction request to nexus
	transaction, err := nexus.SendRequestTransaction(usageCredit, "general", "tina services payment", user.NexusId)
	if err != nil {
		fmt.Println(err)
		return
	}
	if transaction == nil {
		fmt.Println("no transaction created")
		return
	}
	// create a new invoice for this user
	invoice := &types.Invoice{
		Id:            uuid.New(),
		UserId:        user.Id,
		UserNexusId:   user.NexusId,
		TransactionId: transaction.Id,
		Amount:        transaction.Amount,
		Fee:           transaction.Fee,
		Status:        types.TransactionStatusTypePending,
		Details:       "services provided from tina",
	}
	err = db.InsertInvoice(invoice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("invoice %s for user %s created successfully\n", invoice.Id.String(), user.Username)
}
