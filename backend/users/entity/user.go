package entity

import "time"

type User struct {
	CreatedAt    *time.Time `json:"created_at"     gorm:"column:created_at"`
	LastAccessAt *time.Time `json:"last_access_at" gorm:"column:last_access_at"`
	Username     string     `json:"username" gorm:"column:username;size:16;primaryKey;index:idx_username;not null"`
	Password     string     `json:"password" gorm:"column:password;not null"`
}

func (User) TableName() string { return "users" }

func CopyInformationFrom(b UserShow) User {
	var a User
	a.Username = b.Username
	a.CreatedAt = b.CreatedAt
	return a
}
