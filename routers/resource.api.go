package routers

import (
	"cloudPlatformDemo/controller"
	"github.com/gin-gonic/gin"
)

func ResourceRouters(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/costs/list", controller.GetCostList)
		v1.POST("/costs/create")
	}
}

