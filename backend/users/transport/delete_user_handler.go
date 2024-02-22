package transport

import (
	"net/http" // http status

	"github.com/gin-gonic/gin" // request
	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/users/business" // business layer
	"github.com/ltphat2204/blog/backend/users/storage"  // storage layer
	"gorm.io/gorm"                                      // interact with mysql
)

func DeleteUser(db *gorm.DB) func(*gin.Context) {
	// gin handler
	return func(c *gin.Context) {
		// Get username
		username := c.Param("username")

		// Create storage
		s := storage.NewMySqlStorage(db)

		// Initialize business layer with created storage
		biz := business.DeleteUserBusiness(s)

		// Can not delete a user
		if err := biz.DeleteUser(c.Request.Context(), username); err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}

		// Create new user successfully
		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(map[string]string{
			"message": "User " + username + " deleted successfully!",
		}))
	}
}
