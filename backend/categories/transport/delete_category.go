package transport

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/categories/business"
	"github.com/ltphat2204/blog/backend/categories/entity"
	"github.com/ltphat2204/blog/backend/categories/storage"
	"gorm.io/gorm"
)

func DeleteCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		idString, _ := c.Params.Get("id")
		id, _ := strconv.Atoi(idString)

		s := storage.NewMySqlStorage(db)
		biz := business.DeleteCategoryBusiness(s)

		err := biz.DeleteCategory(c.Request.Context(), uint(id))
		if err == entity.ErrCategoryNotFound {
			c.JSON(http.StatusNotFound, common.NewSimpleErrorResponse("Id not found"))
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(http.StatusInternalServerError, "Can not process", err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(map[string]interface{}{
			"id": id,
			"message": "Deleted successfully!",
		}))
	}
}
