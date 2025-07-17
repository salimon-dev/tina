package main

import (
	"context"
	"tina/packages/nexus"
	"tina/packages/openai"
	"tina/packages/types"
)

func HandleEvent(ctx context.Context, event *types.QueueEventInteract) error {
	messages, err := nexus.GetLastMessages(event.ThreadId)
	if err != nil {
		return err
	}
	response, err := openai.SendCompletionRequest(messages)
	if err != nil {
		return err
	}
	err = nexus.SendPlainMessage(event.ThreadId, response.Body)
	return err
}
