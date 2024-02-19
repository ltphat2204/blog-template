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

func UpdateCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		idString, _ := c.Params.Get("id")
		id, _ := strconv.Atoi(idString)

		var post entity.Category

		if err := c.ShouldBind(&post); err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, "Can not process", err.Error()))
			return
		}

		s := storage.NewMySqlStorage(db)
		biz := business.UpdateCategoryBusiness(s)

		err := biz.UpdateCategory(c.Request.Context(), uint(id), &post)
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
