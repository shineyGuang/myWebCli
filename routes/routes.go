package routes

import (
	"myWebCli/controller/user"
	"myWebCli/log"
	"myWebCli/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	// 选择发布模式，dev 还是 release
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()                                // 使用自定义开启gin框架
	r.Use(log.GinLogger(), log.GinRecovery(true)) // 使用自定义中间件，logger 和 错误规避
	v1 := r.Group("/api/v1")
	// 此处填写路由

	v1.POST("/login", user.LoginHandler)               // 登录
	v1.POST("/signUp", user.SignUpHandler)             // 注册
	v1.GET("/refresh_token", user.RefreshTokenHandler) // 刷新token

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/", user.TestWebHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
