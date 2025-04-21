package openai

import (
	"encoding/json"
	"errors"
	"salimon/tina/types"
)

const EMBED_LIMIT = 3

func SendCompletionRequest(messages []types.Message) (*types.Message, error) {
	completionMessages, err := ParseMessages(messages)
	if err != nil {
		return nil, err
	}

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
				Name:        action.Name,
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

	if choice.Message.Content != "" {
		return &types.Message{
			From: "tina",
			Type: "plain",
			Body: choice.Message.Content,
		}, nil
	}

	if choice.Message.ToolCalls == nil {
		return nil, errors.New("no tool call found")
	}
	if len(choice.Message.ToolCalls) == 0 {
		return nil, errors.New("no tool call found")
	}
	toolCall := choice.Message.ToolCalls[0]
	messageType := toolCall.Function.Name
	argumentsStr := toolCall.Function.Arguments
	callId := toolCall.Id

	var arguments json.RawMessage
	err = json.Unmarshal([]byte(argumentsStr), &arguments)

	if err != nil {
		return nil, errors.New("invalid function arguments")
	}

	type MetaType struct {
		CallId string `json:"call_id"`
	}
	type Body struct {
		Meta      MetaType        `json:"meta"`
		Arguments json.RawMessage `json:"arguments"`
	}

	body := Body{
		Meta:      MetaType{CallId: callId},
		Arguments: arguments,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("invalid function arguments")
	}

	return &types.Message{
		From: "tina",
		Type: messageType,
		Body: string(bodyBytes),
	}, nil
}
