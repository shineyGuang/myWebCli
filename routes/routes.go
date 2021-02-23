package routes

import (
	"myWebCli/controller/index"
	"myWebCli/controller/user"
	"myWebCli/log"

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

	// 此处填写路由
	r.GET("/", user.TestWebHandler)
	r.POST("/login", user.LoginHandler)     // 登录
	r.POST("/signUp", user.SignUpHandler)   // 注册
	r.GET("/index", index.HomeIndexHandler) // 首页展示
	return r
}
