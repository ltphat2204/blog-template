package entity

import "time"

type UserUpdate struct {
	CreatedAt    *time.Time `json:"created_at"     gorm:"column:created_at"`
	LastAccessAt *time.Time `json:"last_access_at" gorm:"column:last_access_at"`
	Password string `json:"password" gorm:"column:password;not null"`
}

func (UserUpdate) TableName() string { return "users" }
