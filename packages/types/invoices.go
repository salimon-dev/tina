package types

import "github.com/google/uuid"

type TransactionStatusType uint8

const (
	TransactionStatusPending  TransactionStatusType = 1
	TransactionsStatusDone    TransactionStatusType = 2
	TransactionStatusRejected TransactionStatusType = 3
)

type Invoice struct {
	Id            uuid.UUID             `json:"id" gorm:"type:uuid;primaryKey"`
	UserId        uuid.UUID             `json:"user_id"`
	UserNexusId   uuid.UUID             `json:"user_nexus_id"`
	TransactionId uuid.UUID             `json:"transaction_id"`
	Amount        uint64                `json:"amount" gorm:"type:bigint"`
	Fee           uint64                `json:"fee" gorm:"type:bigint"`
	Status        TransactionStatusType `json:"status" gorm:"type:numeric"`
	Details       string                `json:"details" gorm:"size:256"`
}
type Transaction struct {
	Id          uuid.UUID             `json:"id"`
	SourceId    uuid.UUID             `json:"source_id"`
	TargetId    uuid.UUID             `json:"target_id"`
	Category    string                `json:"category"`
	Description string                `json:"description"`
	Amount      uint64                `json:"amount"`
	Fee         uint64                `json:"fee"`
	Status      TransactionStatusType `json:"status"`
	CreatedAt   int64                 `json:"created_at"`
	UpdatedAt   int64                 `json:"updated_at"`
}
