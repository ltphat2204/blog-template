package entity

import "time"

type UserAccess struct {
	LastAccessAt *time.Time `json:"last_access_at" gorm:"column:last_access_at"`
}

func (UserAccess) TableName() string { return "users" }
