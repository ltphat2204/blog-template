package transport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/ltphat2204/blog/backend/common"
	"github.com/ltphat2204/blog/backend/subcribers/business"
	"github.com/ltphat2204/blog/backend/subcribers/entity"
	"github.com/ltphat2204/blog/backend/subcribers/storage"
	"net/http"
)

func CreateSubcriber(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var subcriber entity.Subcriber

		if err := c.ShouldBind(&subcriber); err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}

		s := storage.NewMySqlStorage(db)
		biz := business.CreateSubcriberBusiness(s)
		res, err := biz.CreateSubcriber(c.Request.Context(), &subcriber)

		if  err != nil {
			c.JSON(http.StatusNotAcceptable, common.NewErrorResponse(http.StatusNotAcceptable, "Can not subcribe", err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(res))
	}
}
