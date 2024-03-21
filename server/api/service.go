package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) MountCommonService(group *gin.RouterGroup) *Controller {
	group.GET("/alive", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
	return c
}

func (c *Controller) MountV1Api(group *gin.RouterGroup) *Controller {
	// Return companies in same sectors
	group.POST("/land/pnu", func(ctx *gin.Context) {
		servicePnuSelect(ctx, c)
	})
	return c
}
