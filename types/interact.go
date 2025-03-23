package types

type MessageType string

const (
	MessageTypePlain MessageType = "plain"
)

type Message struct {
	From string      `json:"from" validate:"required"`
	Type MessageType `json:"type" validate:"required"`
	Body string      `json:"body" validate:"required"`
}

type InteractSchema struct {
	Data []Message `json:"data" validate:"required"`
}
