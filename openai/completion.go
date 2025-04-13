package openai

import (
	"encoding/json"
	"errors"
	"salimon/tina/types"
)

const EMBED_LIMIT = 3

func ParseMessages(messages []types.Message) []CompletionMessage {
	result := make([]CompletionMessage, len(messages))
	for i, m := range messages {
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

	// embed messages to get embedding vector
	messageLen := min(len(messages), EMBED_LIMIT)
	embeddingMessages := make([]types.Message, messageLen)
	for i := 0; i < messageLen; i++ {
		embeddingMessages[i] = messages[len(messages)-i-1]
	}

	embeddingData, err := json.Marshal(embeddingMessages)
	vectors, err := SendEmbeddingRequest(string(embeddingData))

	if err != nil {
		return nil, err
	}

	action := GetBestAction(vectors)

	if action != nil {
		tools := make([]CompletionTool, 1)
		tools[0] = CompletionTool{
			Type: "function",
			Function: CompletionFunction{
				Name:        action.Type,
				Description: action.Description,
				Strict:      true,
				Parameters:  action.Parameters,
			},
		}
		params.Tools = tools
	}

	if len(vectors) == 0 {
		return nil, errors.New("failed to send embedding request")
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
