package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	ID             uint64    `gorm:"primaryKey; column:id"`
	FullName       string    `gorm:"column:full_name"`
	Email          string    `gorm:"column:email"`
	Login          string    `gorm:"column:login"`
	HashedPassword string    `gorm:"column:hashed_password"`
	Password       string    `gorm:"-"`
	LastLogin      time.Time `gorm:"column:last_login"`
	CreatedAt      time.Time `gorm:"autoCreateTime; column:created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime:milli; column:updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Validate() error {
	return validator.New().Struct(u)
}
