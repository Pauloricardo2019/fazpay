package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	ID             uint64    `gorm:"primaryKey; column:id"`
	FirstName      string    `validate:"required" gorm:"column:first_name"`
	LastName       string    `validate:"required" gorm:"column:last_name" `
	Email          string    `validate:"required" gorm:"column:email" `
	HashedPassword string    `gorm:"column:hashed_password"`
	Password       string    `gorm:"-"`
	CreatedAt      time.Time `gorm:"autoCreateTime; column:created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime:milli; column:updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Validate() error {
	return validator.New().Struct(u)
}
