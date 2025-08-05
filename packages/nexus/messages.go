package nexus

import (
	"encoding/json"
	"errors"
	"fmt"
	"tina/packages/types"
)

func GetThreadMessages(threadId string, count uint8) ([]types.Message, error) {
	responseData, err := SendHttpRequest("GET", fmt.Sprintf("/member/messages?thread_id=%s&page_size=%d&sort=created_at&order=desc", threadId, count), nil)
	if err != nil {
		return nil, err
	}
	type Response struct {
		Data []types.Message `json:"data"`
	}
	var response Response
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func GetLastMessages(threadId string) ([]types.Message, error) {
	messages, err := GetThreadMessages(threadId, 3)
	if err != nil {
		return nil, err
	}
	if messages == nil {
		return nil, errors.New("no messages found")
	}
	if len(messages) == 0 {
		return nil, errors.New("no messages found")
	}
	msgLen := len(messages)
	result := make([]types.Message, msgLen)
	for i, message := range messages {
		result[msgLen-i-1] = message
	}
	return result, nil
}

func SendMessage(threadId string, body string, messageType types.MessageType) error {
	type Payload struct {
		ThreadId string            `json:"thread_id"`
		Body     string            `json:"body"`
		Type     types.MessageType `json:"type"`
	}
	payload := Payload{
		ThreadId: threadId,
		Body:     body,
		Type:     messageType,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	_, err = SendHttpRequest("POST", "/member/messages/send", payloadBytes)
	return err
}

func StartThread(targetId string, name string, category types.ThreadCategory) (*types.Thread, error) {
	type Payload struct {
		TargetId string               `json:"target_id"`
		Name     string               `json:"name"`
		Category types.ThreadCategory `json:"category"`
	}
	payload := Payload{
		TargetId: targetId,
		Name:     name,
		Category: category,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	responseData, err := SendHttpRequest("POST", "/member/threads/start", payloadBytes)
	if err != nil {
		return nil, err
	}
	var response types.Thread
	err = json.Unmarshal(responseData, &response)
	return &response, err
}
