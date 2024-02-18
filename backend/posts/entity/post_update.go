package entity

import (
	"github.com/ltphat2204/blog/backend/categories/entity"
)

type PostUpdate struct {
	ID          uint            `json:"id"          gorm:"primary_key;auto_increment"`
	Title       string          `json:"title"       gorm:"type:varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci;unique;not null"`
	Banner      string          `json:"banner"      gorm:"type:text; not null"`
	Description string          `json:"description" gorm:"type:text CHARACTER SET utf8 COLLATE utf8_general_ci;not null"`
	Content     string          `json:"content"     gorm:"type:longtext  CHARACTER SET utf8 COLLATE utf8_general_ci;not null"`
	View        int             `json:"view"        gorm:"type:int;default:0"`
	CategoryID  uint            `json:"category_id" gorm:"index;not null"`
	Category    entity.Category `json:"category"    gorm:"foreignKey:CategoryID;references:ID;"`
	IsDeleted   *bool           `json:"is_deleted"  gorm:"type:boolean; is_deleted; default: false"`
	IsDraft     *bool           `json:"is_draft"    gorm:"type:boolean; is_draft; default: true"`
}

func (PostUpdate) TableName() string { return "posts" }
