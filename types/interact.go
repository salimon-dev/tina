package types

type MessageType string

type Message struct {
	From string `json:"from" validate:"required"`
	Type string `json:"type" validate:"required"`
	Body string `json:"body" validate:"required"`
}

type InteractSchema struct {
	Data []Message `json:"data" validate:"required,dive"`
}
