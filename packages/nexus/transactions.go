package nexus

import (
	"encoding/json"
	"tina/packages/types"

	"github.com/google/uuid"
)

func SendRequestTransaction(amount uint64, category string, description string, userId uuid.UUID) (*types.Transaction, error) {
	type Payload struct {
		UserId      uuid.UUID `json:"user_id"`
		Amount      uint64    `json:"amount"`
		Description string    `json:"description"`
		Category    string    `json:"category"`
	}

	payload := Payload{
		UserId:      userId,
		Amount:      amount,
		Description: description,
		Category:    category,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	responseData, err := SendHttpRequest("POST", "/member/transactions/request", payloadBytes)

	if err != nil {
		return nil, err
	}
	var response types.Transaction
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
