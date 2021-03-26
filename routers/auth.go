package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRouter(route *gin.Engine) {
	auth := route.Group("auth")
	{
		auth.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "哈哈哈2")
		})
	}
}
