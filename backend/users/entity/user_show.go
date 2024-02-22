package entity

import "time"

type UserShow struct {
	Username string `json:"username" gorm:"column:username;size:16;primaryKey;index:idx_username;not null"`
	CreatedAt    *time.Time `json:"created_at"     gorm:"column:created_at"`
	LastAccessAt *time.Time `json:"last_access_at" gorm:"column:last_access_at"`
}

func (UserShow) TableName() string { return "users" }
