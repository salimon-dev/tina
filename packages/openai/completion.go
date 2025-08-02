package openai

import (
	"encoding/json"
	"errors"
	"tina/packages/types"
)

const EMBED_LIMIT = 3

func SendCompletionRequest(messages []types.Message) (*CompletionParsedResponse, error) {

	completionMessages, err := ParseMessages(messages)
	if err != nil {
		return nil, err
	}

	params := CompletionParams{
		Messages:            completionMessages,
		MaxCompletionTokens: 500,
		Model:               "gpt-4o-mini",
		Temperature:         0.2,
	}

	// TODO: it will be useful in future
	// // embed messages to get embedding vector
	// messageLen := min(len(messages), EMBED_LIMIT)
	// embeddingMessages := make([]gomsg.Message, messageLen)
	// for i := 0; i < messageLen; i++ {
	// 	embeddingMessages[i] = messages[len(messages)-i-1]
	// }

	// embeddingData, err := json.Marshal(embeddingMessages)
	// vectors, err := SendEmbeddingRequest(string(embeddingData))

	// if err != nil {
	// 	return nil, err
	// }

	// action := GetBestAction(vectors)

	// if action != nil {
	// 	tools := make([]CompletionTool, 1)
	// 	tools[0] = CompletionTool{
	// 		Type: "function",
	// 		Function: CompletionFunction{
	// 			Name:        action.Name,
	// 			Description: action.Description,
	// 			Strict:      true,
	// 			Parameters:  action.Parameters,
	// 		},
	// 	}
	// 	params.Tools = tools
	// }

	// if len(vectors) == 0 {
	// 	return nil, errors.New("failed to send embedding request")
	// }

	paramsData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	response, err := SendHttpRequest("POST", "/v1/chat/completions", paramsData)
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

	result := CompletionParsedResponse{}
	result.Usage = uint64(completionResponse.Usage.TotalTokens)

	if choice.Message.Content != "" {
		result.Body = choice.Message.Content
	}

	return &result, nil

	// if choice.Message.ToolCalls == nil {
	// 	return nil, errors.New("no tool call found")
	// }
	// if len(choice.Message.ToolCalls) == 0 {
	// 	return nil, errors.New("no tool call found")
	// }
	// toolCall := choice.Message.ToolCalls[0]
	// messageType := toolCall.Function.Name
	// argumentsStr := toolCall.Function.Arguments
	// callId := toolCall.Id

	// var parameters gomsg.Parameters
	// err = json.Unmarshal([]byte(argumentsStr), &parameters)

	// return &gomsg.Message{
	// 	From: "tina",
	// 	Type: messageType,
	// 	Meta: &gomsg.Meta{
	// 		ActionId: callId,
	// 	},
	// 	Parameters: &parameters,
	// }, nil

}
