package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	categoriesModel "github.com/ltphat2204/blog/backend/categories/entity"
	categoriesTrans "github.com/ltphat2204/blog/backend/categories/transport"
	"github.com/ltphat2204/blog/backend/common"
	postsModel "github.com/ltphat2204/blog/backend/posts/entity"
	postTrans "github.com/ltphat2204/blog/backend/posts/transport"
	subcribersModel "github.com/ltphat2204/blog/backend/subcribers/entity"
	subTrans "github.com/ltphat2204/blog/backend/subcribers/transport"
	"gorm.io/gorm"
)

var database *gorm.DB

func init() {
	mode := os.Getenv("GIN_MODE")

	if mode != "release" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("No setting in .env.local found in Debug mode!")
		}
	}

	database = common.ConnectDB()

	database.AutoMigrate(&subcribersModel.Subcriber{})
	database.AutoMigrate(&categoriesModel.Category{})
	database.AutoMigrate(&postsModel.Post{})
}

func main() {
	app := gin.Default()
	app.Use(cors.Default())

	app.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running!"})
	})

	postRoute := app.Group("/posts")
	{
		postRoute.GET("", postTrans.GetPosts(database))
		postRoute.POST("", postTrans.CreatePost(database))
		postRoute.GET("/:id", postTrans.GetPost(database))
		postRoute.PATCH("/:id", postTrans.UpdatePost(database))
		postRoute.DELETE("/:id", postTrans.DeletePost(database))
	}

	subcriberRoute := app.Group("/subcribers")
	{
		subcriberRoute.GET("", subTrans.GetSubcriber(database))
		subcriberRoute.POST("", subTrans.CreateSubcriber(database))
	}

	categoryRoute := app.Group("/categories")
	{
		categoryRoute.GET("", categoriesTrans.GetCategories(database))
		categoryRoute.POST("", categoriesTrans.CreateCategory(database))
		categoryRoute.GET("/:id", categoriesTrans.GetCategory(database))
	}

	app.Run()
}
