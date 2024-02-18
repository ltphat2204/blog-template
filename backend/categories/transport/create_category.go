package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/categories/business"
	"github.com/ltphat2204/blog/backend/categories/entity"
	"github.com/ltphat2204/blog/backend/categories/storage"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var category entity.Category

		if err := c.ShouldBind(&category); err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}

		s := storage.NewMySqlStorage(db)
		biz := business.CreateCategoryBusiness(s)
		res, err := biz.CreateCategory(c.Request.Context(), &category)

		if  err != nil {
			c.JSON(http.StatusNotAcceptable, common.NewErrorResponse(http.StatusNotAcceptable, "Can not create category", err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(res))
	}
}
