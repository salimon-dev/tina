package types

import "github.com/google/uuid"

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
