package entity

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey;auto_increment"`
	Name string `json:"name" gorm:"type:varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci;unique;not null"`
}
