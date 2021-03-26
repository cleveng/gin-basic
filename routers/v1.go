package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func V1Router(route *gin.Engine) {
	v1 := route.Group("v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "哈哈哈")
		})
	}
}
