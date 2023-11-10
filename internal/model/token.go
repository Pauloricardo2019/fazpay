package model

import (
	"time"
)

const (
	ActionLogin = "login"
)

type Token struct {
	ID        uint64    `gorm:"primaryKey; column:id"`
	UserID    uint64    `gorm:"column:user_id"`
	User      User      `gorm:"foreignKey:user_id;references:id"`
	Value     string    `gorm:"column:value"`
	CreatedAt time.Time `gorm:"autoCreateTime; column:created_at"`
	ExpiresAt time.Time `gorm:"autoUpdateTime:milli; column:expires_at"`
}

func (Token) TableName() string {
	return "tokens"
}
