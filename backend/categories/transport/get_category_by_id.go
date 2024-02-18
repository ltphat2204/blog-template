package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/blog/backend/categories/business"
	"github.com/ltphat2204/blog/backend/categories/entity"
	"github.com/ltphat2204/blog/backend/categories/storage"
	"github.com/ltphat2204/blog/backend/common"
	"gorm.io/gorm"
)

func GetCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		idString, _ := c.Params.Get("id")
		id, _ := strconv.Atoi(idString)

		s := storage.NewMySqlStorage(db)
		biz := business.GetCategoryBusiness(s)

		post, err := biz.GetCategory(c.Request.Context(), uint(id))
		if err == entity.ErrCategoryNotFound {
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
