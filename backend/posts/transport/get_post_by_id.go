package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/posts/business"
	"github.com/ltphat2204/blog/backend/posts/entity"
	"github.com/ltphat2204/blog/backend/posts/storage"
	"gorm.io/gorm"
)

func GetPost(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		idString, _ := c.Params.Get("id")
		id, _ := strconv.Atoi(idString)

		s := storage.NewMySqlStorage(db)
		biz := business.GetPostBusiness(s)

		post, err := biz.GetPost(c.Request.Context(), uint(id))
		if err == entity.ErrPostNotFound {
			c.JSON(http.StatusNotFound, common.NewSimpleErrorResponse("Id not found"))
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, "Can not process", err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(post))
	}
}
