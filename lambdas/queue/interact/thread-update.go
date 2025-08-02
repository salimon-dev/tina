package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"tina/packages/db"
	"tina/packages/nexus"
	"tina/packages/openai"
	"tina/packages/types"

	"github.com/google/uuid"
)

func HandleThreadUpdateEvent(ctx context.Context, event *types.QueueEventInteract) error {
	messages, err := nexus.GetLastMessages(event.ThreadId)
	if err != nil {
		return err
	}
	response, err := openai.SendCompletionRequest(messages)
	if err != nil {
		return err
	}
	fmt.Println("passed the completion")
	err = nexus.SendPlainMessage(event.ThreadId, response.Body)
	if err != nil {
		return err
	}

	username, nexusId := FindUserInfoFromMessages(messages)
	if username == "" || nexusId.String() == "" {
		return errors.New("username or nexusId is empty")
	}
	err = db.UpdateUserUsage(username, nexusId, response.Usage, types.UserStatusActive)
	return err
}

func FindUserInfoFromMessages(messages []types.Message) (string, uuid.UUID) {
	var username string
	var nexusId uuid.UUID
	for _, message := range messages {
		// TODO: isolate getEnv methods
		if message.Username != os.Getenv("NEXUS_USERNAME") {
			username = message.Username
			nexusId = message.UserId
		}
	}
	return username, nexusId
}
