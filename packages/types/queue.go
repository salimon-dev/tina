package types

type QueueEventInteractType string

const (
	QueueEventInteractTypeThreadUpdate QueueEventInteractType = "THREAD_UPDATE"
	QueueEventInteractTypeTransaction  QueueEventInteractType = "TRANSACTION"
)

type QueueEventInteract struct {
	Type          QueueEventInteractType `json:"type"`
	TransactionId string                 `json:"transaction_id"`
	ThreadId      string                 `json:"thread_id"`
	Origin        string                 `json:"origin"`
}
