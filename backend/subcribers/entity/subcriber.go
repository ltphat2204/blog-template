package entity

import (
	"time"
)

type Subcriber struct {
	Email     string    `json:"email"       gorm:"primary_key;type:varchar(36);unique;not null"`
	Name      string    `json:"name"        gorm:"type:varchar(20); not null"`
	CreatedAt time.Time `json:"created_at"  gorm:"autoCreateTime"`
}
