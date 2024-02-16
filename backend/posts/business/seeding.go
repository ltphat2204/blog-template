package business

import (
	"log"

	"github.com/ltphat2204/blog/backend/posts/entity"
	"gorm.io/gorm"
)

func CreateTablePosts(database *gorm.DB) error {
	if created := database.Migrator().HasTable("posts"); !created {
		if err := database.Migrator().CreateTable(&entity.Post{}); err != nil {
			return err
		}
		log.Printf("Table `posts` created successfully!")
	}

	return nil
}