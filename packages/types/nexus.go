package types

import "github.com/google/uuid"

type Thread struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"type:string;size:64"`
	CreatedAt int64     `json:"created_at" gorm:"type:bigint"`
	UpdatedAt int64     `json:"updated_at" gorm:"type:bigint"`
}

type MessageType uint8

const (
	MessageTypeText MessageType = 1
)

type Message struct {
	Id        uuid.UUID   `json:"id" gorm:"type:uuid;primaryKey"`
	Body      string      `json:"body" gorm:"type:text"`
	UserId    uuid.UUID   `json:"user_id" gorm:"type:uuid"`
	ThreadId  uuid.UUID   `json:"thread_id" gorm:"type:uuid"`
	Type      MessageType `json:"type" gorm:"type:numeric"`
	CreatedAt int64       `json:"created_at" gorm:"type:bigint"`
	UpdatedAt int64       `json:"updated_at" gorm:"type:bigint"`
}

type ThreadMemberType uint8

const (
	ThreadMemberTypeOwner  ThreadMemberType = 1
	ThreadMemberTypeAdmin  ThreadMemberType = 2
	ThreadMemberTypeNormal ThreadMemberType = 3
)

type ThreadMember struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserId    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	ThreadId  uuid.UUID `json:"thread_id" gorm:"type:uuid"`
	CreatedAt int64     `json:"created_at" gorm:"type:bigint"`
	UpdatedAt int64     `json:"updated_at" gorm:"type:bigint"`
}
