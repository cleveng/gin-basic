package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode) // DebugMode 调试  ReleaseMode 上线
	router := gin.Default()
	router.Use(location.Default())
	store, _ := sessions.NewRedisStore(10, "tcp", "127.0.0.1:6379", "", []byte("jsfQ9idPxzU"))
	router.Use(sessions.Sessions("SESSION_ID", store))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 15 * time.Second,
	}))

	//router.LoadHTMLGlob("views/**/*")
	router.StaticFS("/public", http.Dir("public")) // 设置附件资源目录
	//router.Static("/assets", "public/assets")

	// v1接口 router
	V1Router(router)

	// auth接口 router
	AuthRouter(router)

	//router.GET("error", Common.NoRouteError)
	//router.NoRoute(Common.NoRouteError)

	return router
}
