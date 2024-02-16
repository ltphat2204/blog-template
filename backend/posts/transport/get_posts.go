package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/posts/business"
	"github.com/ltphat2204/blog/backend/posts/storage"
	"gorm.io/gorm"
)

func GetPosts(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var pagination common.Pagination
		if err := c.ShouldBind(&pagination); err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}
		pagination.Normalize()

		queryParams := c.Request.URL.Query()
		condition := make(map[string]interface{})

		for key, value := range queryParams {
			if key != "limit" && key != "page" {
				condition[key] = value[0]
			}
		}

		s := storage.NewMySqlStorage(db)
		biz := business.GetPostsBusiness(s)

		posts, err := biz.GetPosts(c.Request.Context(), condition, &pagination)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(posts, pagination))
	}
}
