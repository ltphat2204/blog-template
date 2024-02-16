package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ltphat2204/blog/backend/common"
	postBiz "github.com/ltphat2204/blog/backend/posts/business"
	postTrans "github.com/ltphat2204/blog/backend/posts/transport"
	subTrans "github.com/ltphat2204/blog/backend/subcribers/transport"
	subBiz "github.com/ltphat2204/blog/backend/subcribers/business"
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

	if err := postBiz.CreateTablePosts(database); err != nil {
		log.Fatal(err.Error())
	}

	if err := subBiz.CreateTableSubcribers(database); err != nil {
		log.Fatal(err.Error())
	}
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

	app.Run()
}
