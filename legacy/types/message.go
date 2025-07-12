package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func (r *Message) BeforeCreate(tx *gorm.DB) (err error) {
	r.CreatedAt = time.Now().Unix()
	r.UpdatedAt = time.Now().Unix()
	return nil
}
func (r *Message) BeforeUpdate(tx *gorm.DB) (err error) {
	r.UpdatedAt = time.Now().Unix()
	return nil
}
