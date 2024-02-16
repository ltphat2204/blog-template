package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/posts/business"
	"github.com/ltphat2204/blog/backend/posts/entity"
	"github.com/ltphat2204/blog/backend/posts/storage"
	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var post entity.Post

		if err := c.ShouldBind(&post); err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}

		s := storage.NewMySqlStorage(db)
		biz := business.CreatePostBusiness(s)
		res, err := biz.CreatePost(c.Request.Context(), &post)

		if  err != nil {
			c.JSON(http.StatusNotAcceptable, common.NewErrorResponse(http.StatusNotAcceptable, "Can not create post", err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(res))
	}
}
