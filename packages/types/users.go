package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStatus uint8

const (
	UserStatusActive   UserStatus = 1
	UserStatusInActive UserStatus = 2
)

type User struct {
	Id            uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	Username      string     `json:"username" gorm:"size:32;unique;not null"`
	NexusId       uuid.UUID  `json:"nexus_id" gorm:"type:uuid"`
	Status        UserStatus `json:"status" gorm:"type:numeric"`
	Usage         uint64     `json:"usage" gorm:"type:bigint"`
	CreditDebt    uint64     `json:"credit_debt" gorm:"type:bigint"`
	DebtSoftLimit uint64     `json:"debt_soft_limit" gorm:"type:bigint"`
	DebtHardLimit uint64     `json:"debt_hard_limit" gorm:"type:bigint"`
	CreatedAt     int64      `json:"created_at" gorm:"type:bigint"`
	UpdatedAt     int64      `json:"updated_at" gorm:"type:bigint"`
}

func (r *User) BeforeCreate(tx *gorm.DB) (err error) {
	r.CreatedAt = time.Now().Unix()
	r.UpdatedAt = time.Now().Unix()
	return nil
}
func (r *User) BeforeUpdate(tx *gorm.DB) (err error) {
	r.UpdatedAt = time.Now().Unix()
	return nil
}
