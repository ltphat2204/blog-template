package business

import (
	"log"

	"github.com/ltphat2204/blog/backend/subcribers/entity"
	"gorm.io/gorm"
)

func CreateTableSubcribers(database *gorm.DB) error {
	if created := database.Migrator().HasTable("subcribers"); !created {
		if err := database.Migrator().CreateTable(&entity.Subcriber{}); err != nil {
			return err
		}
		log.Printf("Table `subcribers` created successfully!")
	}

	return nil
}