package types

import "github.com/google/uuid"

type Thread struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

type MessageType uint8

const (
	MessageTypeText MessageType = 1
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

type ThreadMemberType uint8

const (
	ThreadMemberTypeOwner  ThreadMemberType = 1
	ThreadMemberTypeAdmin  ThreadMemberType = 2
	ThreadMemberTypeNormal ThreadMemberType = 3
)

type ThreadMember struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	ThreadId  uuid.UUID `json:"thread_id"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}
