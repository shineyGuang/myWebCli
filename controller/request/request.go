package request

import (
	"myWebCli/middlewares"

	"github.com/gin-gonic/gin"
)

// GetCurrentUser 获取当前登录用的userId
func GetCurrentUser(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(middlewares.ContextUserIDKey)
	if !ok {
		err = middlewares.ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = middlewares.ErrorUserNotLogin
	}
	return
}
