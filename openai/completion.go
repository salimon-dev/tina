package openai

import (
	"encoding/json"
	"errors"
	"salimon/tina/types"
)

func ParseMessages(message []types.Message) []CompletionMessage {
	result := make([]CompletionMessage, len(message))
	for i, m := range message {
		role := m.From
		content := m.Body
		if role != "user" {
			role = "assistant"
		}
		result[i] = CompletionMessage{
			Role:    role,
			Content: content,
		}
	}
	return result
}

func SendCompletionRequest(messages []types.Message) (*types.Message, error) {
	completionMessages := ParseMessages(messages)
	params := CompletionParams{
		Messages:            completionMessages,
		Model:               "gpt-4o-mini",
		MaxCompletionTokens: 256,
		Temperature:         0.2,
	}

	paramsData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	response, err := SendRequest("POST", "/v1/chat/completions", paramsData)

	if err != nil {
		return nil, err
	}

	var completionResponse CompletionResponse
	err = json.Unmarshal(response, &completionResponse)
	if err != nil {
		return nil, err
	}

	if len(completionResponse.Choices) == 0 {
		return nil, errors.New("no choice found")
	}

	choice := completionResponse.Choices[0]
	return &types.Message{
		From: "tina",
		Type: types.MessageTypePlain,
		Body: choice.Message.Content,
	}, nil
}
