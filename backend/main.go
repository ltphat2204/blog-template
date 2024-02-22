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
	usersModel "github.com/ltphat2204/blog/backend/users/entity"
	usersTrans "github.com/ltphat2204/blog/backend/users/transport"
	usersBiz "github.com/ltphat2204/blog/backend/users/business"
	"github.com/ltphat2204/blog/backend/middlewares"
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
	database.AutoMigrate(&usersModel.User{})
	if err := usersBiz.CreateAdmin(database); err != nil {
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

	categoryRoute := app.Group("/categories")
	{
		categoryRoute.GET("", categoriesTrans.GetCategories(database))
		categoryRoute.POST("", categoriesTrans.CreateCategory(database))
		categoryRoute.GET("/:id", categoriesTrans.GetCategory(database))
	}

	categoryRoute.Group("/users")
	{
		categoryRoute.GET("", middlewares.AuthMiddleware, middlewares.AdminMiddleWare, usersTrans.GetAllUsers(database))
		categoryRoute.POST("", middlewares.AuthMiddleware, middlewares.AdminMiddleWare, usersTrans.CreateUser(database))

		categoryRoute.GET("/:username", middlewares.AuthMiddleware, middlewares.AdminMiddleWare, usersTrans.GetUserByUsername(database))
		categoryRoute.PATCH("/:username", middlewares.AuthMiddleware, usersTrans.UpdatePasswordByUsername(database))
		categoryRoute.DELETE("/:username", middlewares.AuthMiddleware, middlewares.AdminMiddleWare, usersTrans.DeleteUser(database))

		categoryRoute.GET("/token", usersTrans.GetTokenFromUser(database))
		categoryRoute.GET("/verify", middlewares.AuthMiddleware, usersTrans.VerifyToken(database))
	}

	app.Run()
}
