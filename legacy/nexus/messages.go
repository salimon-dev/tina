package nexus

import (
	"encoding/json"
	"errors"
	"fmt"
	"salimon/tina/helpers"
	"salimon/tina/types"

	"github.com/google/uuid"
)

func GetThreadMessages(threadId uuid.UUID, count uint8) ([]types.Message, error) {
	responseData, err := helpers.SendNexusRequest(fmt.Sprintf("/member/messages?thread_id=%s&page=1&page_size=%d&sort=created_at&order=desc", threadId, count), "GET", nil)
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

func GetLastMessages(threadId uuid.UUID) ([]types.Message, error) {
	messages, err := GetThreadMessages(threadId, 3)
	if err != nil {
		return nil, errors.New("failed to fetch messages")
	}
	if messages == nil || len(messages) == 0 {
		return nil, errors.New("no messages found")
	}
	msgLen := len(messages)
	result := make([]types.Message, msgLen)
	for i, message := range messages {
		result[msgLen-i-1] = message
	}
	return result, nil
}

func SendPlainMessage(threadId uuid.UUID, body string) error {
	type Payload struct {
		ThreadId uuid.UUID `json:"thread_id"`
		Body     string    `json:"body"`
	}
	payload := Payload{
		ThreadId: threadId,
		Body:     body,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	_, err = helpers.SendNexusRequest("/member/messages/send", "POST", payloadBytes)
	return err
}
