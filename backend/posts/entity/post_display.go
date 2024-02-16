package entity

import (
	"time"
)

type PostDisplay struct {
	ID          uint      `json:"id"          gorm:"primary_key;auto_increment"`
	Title       string    `json:"title"       gorm:"type:varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci;unique;not null"`
	Banner      string    `json:"banner"      gorm:"type:text; not null"`
	Description string    `json:"description" gorm:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci;not null"`
	Content     string    `json:"content"     gorm:"type:longtext  CHARACTER SET utf8 COLLATE utf8_general_ci;not null"`
	View 		int		  `json:"view"        gorm:"type:int;default:0"`
	CreatedAt   time.Time `json:"created_at"  gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at"  gorm:"autoUpdateTime"`
}

func (PostDisplay) TableName() string { return "posts" }
