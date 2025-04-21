package openai

import (
	"encoding/json"
	"errors"
	"salimon/tina/types"
)

func ParseActionResult(message *types.Message) (*CompletionMessage, error) {
	var result CompletionMessage
	bodyStr := message.Body
	type MetaType struct {
		CallId string `json:"call_id"`
	}
	type ActionResultType struct {
		Meta      MetaType        `json:"meta"`
		Arguments json.RawMessage `json:"arguments"`
	}

	var body ActionResultType
	err := json.Unmarshal([]byte(bodyStr), &body)
	if err != nil {
		return nil, err
	}
	result.Role = "tool"

	content, err := json.Marshal(body.Arguments)
	if err != nil {
		return nil, err
	}
	result.Content = string(content)
	result.ToolCallId = body.Meta.CallId
	return &result, nil
}

func ParseActionCall(message *types.Message) (*CompletionMessage, error) {
	var result CompletionMessage

	bodyStr := message.Body
	type MetaType struct {
		CallId string `json:"call_id"`
	}
	type ActionResultType struct {
		Meta      MetaType        `json:"meta"`
		Arguments json.RawMessage `json:"arguments"`
	}

	var body ActionResultType
	err := json.Unmarshal([]byte(bodyStr), &body)
	if err != nil {
		return nil, err
	}

	argumentsStr, err := json.Marshal(body.Arguments)
	if err != nil {
		return nil, err
	}

	result.Role = "assistant"
	toolCall := CompletionToolCall{
		Type: "function",
		Id:   body.Meta.CallId,
		Function: CompletionToolCallFunction{
			Name:      message.Type,
			Arguments: string(argumentsStr),
		},
	}
	result.ToolCalls = []CompletionToolCall{toolCall}
	return &result, nil
}

func ParsePlainMessage(message *types.Message) (*CompletionMessage, error) {
	var result CompletionMessage
	if message.From == "user" {
		result.Role = "user"
	} else {
		result.Role = "assistant"
	}
	result.Content = message.Body
	return &result, nil
}

func ParseSingleMessage(message *types.Message) (*CompletionMessage, error) {
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

func ParseMessages(messages []types.Message) ([]CompletionMessage, error) {
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
