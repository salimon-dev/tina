package openai

import (
	"encoding/json"
	"errors"

	"github.com/salimon-dev/gomsg"
)

func ParseActionResult(message *gomsg.Message) (*CompletionMessage, error) {
	var result CompletionMessage

	result.Role = "tool"
	result.ToolCallId = message.Meta.ActionId
	content, err := json.Marshal(message.Result)

	if err != nil {
		return nil, err
	}
	result.Content = string(content)

	return &result, nil
}

func ParseActionCall(message *gomsg.Message) (*CompletionMessage, error) {
	var result CompletionMessage

	result.Role = "assistant"

	arguments, err := json.Marshal(message.Parameters)
	if err != nil {
		return nil, err
	}

	toolCall := CompletionToolCall{
		Type: "function",
		Id:   message.Meta.ActionId,
		Function: CompletionToolCallFunction{
			Name:      message.Type,
			Arguments: string(arguments),
		},
	}
	result.ToolCalls = []CompletionToolCall{toolCall}
	return &result, nil
}

func ParsePlainMessage(message *gomsg.Message) (*CompletionMessage, error) {
	var result CompletionMessage
	if message.From == "user" {
		result.Role = "user"
	} else {
		result.Role = "assistant"
	}
	result.Content = *message.Body
	return &result, nil
}

func ParseSingleMessage(message *gomsg.Message) (*CompletionMessage, error) {
	if message == nil {
		return nil, errors.New("message is nil")
	}
	if message.Type == "actionResult" {
		return ParseActionResult(message)
	}
	if message.Type != "plain" {
		return ParseActionCall(message)
	}
	// message is plain
	return ParsePlainMessage(message)
}

func ParseMessages(messages []gomsg.Message) ([]CompletionMessage, error) {
	result := make([]CompletionMessage, len(messages))
	for i, m := range messages {
		parsedMessage, err := ParseSingleMessage(&m)
		if err != nil {
			return nil, errors.New("failed to parse message")
		}
		result[i] = *parsedMessage
	}
	return result, nil
}
