package types

import "github.com/google/uuid"

type ThreadCategory uint8

const (
	ThreadCategoryChat    ThreadCategory = 1
	ThreadCategroyPayment ThreadCategory = 2
)

type Thread struct {
	Id        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Category  ThreadCategory `json:"category"`
	MemberIds []string       `json:"member_ids"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
}

type MessageType uint8

const (
	MessageTypeText        MessageType = 1
	MessageTypeTransaction MessageType = 2
)

type Message struct {
	Id        uuid.UUID   `json:"id"`
	Body      string      `json:"body"`
	UserId    uuid.UUID   `json:"user_id"`
	Username  string      `json:"username"`
	ThreadId  uuid.UUID   `json:"thread_id"`
	Type      MessageType `json:"type"`
	CreatedAt int64       `json:"created_at"`
	UpdatedAt int64       `json:"updated_at"`
}

type ThreadMember struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	ThreadId  uuid.UUID `json:"thread_id"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}
type TransactionStatusType int8

const (
	TransactionStatusTypePending  TransactionStatusType = 1
	TransactionStatusTypeDone     TransactionStatusType = 2
	TransactionStatusTypeRejected TransactionStatusType = 3
)

type Transaction struct {
	Id             uuid.UUID             `json:"id"`
	SourceId       uuid.UUID             `json:"source_id"`
	TargetId       uuid.UUID             `json:"target_id"`
	Category       string                `json:"category"`
	Description    string                `json:"description"`
	Amount         uint64                `json:"amount"`
	Fee            uint64                `json:"fee"`
	Status         TransactionStatusType `json:"status"`
	CreatedAt      int64                 `json:"created_at"`
	UpdatedAt      int64                 `json:"updated_at"`
	SourceUsername string                `json:"source_username"`
	TargetUsername string                `json:"target_username"`
}
