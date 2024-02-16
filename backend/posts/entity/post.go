package entity

import (
	"time"
)

type Post struct {
	ID          uint      `json:"id"          gorm:"primary_key;auto_increment"`
	Title       string    `json:"title"       binding:"required" gorm:"type:varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci;unique;not null"`
	Banner      string    `json:"banner"      binding:"required" gorm:"type:text; not null"`
	Description string    `json:"description" binding:"required" gorm:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci;not null"`
	Content     string    `json:"content"     binding:"required" gorm:"type:longtext  CHARACTER SET utf8 COLLATE utf8_general_ci;not null"`
	CreatedAt   time.Time `json:"created_at"  gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at"  gorm:"autoUpdateTime"`
	View 		int		  `json:"view"        gorm:"type:int;default:0"`
	IsDeleted   bool      `json:"is_deleted"  gorm:"is_deleted;type:boolean;default:false"`
	IsDraft     bool      `json:"is_draft"    gorm:"is_draft;type:boolean;default:true"`
}
