package user

import (
	"errors"
	"myWebCli/controller/response"
	"myWebCli/db/mysql"
	"myWebCli/logic/SignUp"
	"myWebCli/models"

	"github.com/gin-gonic/gin"
)

func TestWebHandler(c *gin.Context) {
	response.ResSuccess(c, response.TestSuccess, "Welcome to my web!")
}

//LoginHandler 登录
func LoginHandler(c *gin.Context) {

}

//SignUpHandler 注册
func SignUpHandler(c *gin.Context) {
	// 获取请求参数，并且校验
	var fo models.RegisterForm
	if err := c.ShouldBindJSON(&fo); err != nil {
		response.ResErrorWithString(c, response.ParamsValidatorError, err.Error())
		return
	}
	// 注册用户
	err := SignUp.Register(&fo)
	if errors.Is(err, mysql.ErrorUserExit) {
		response.ResError(c, response.CodeUserExit)
		return
	}
	if err != nil {
		response.ResError(c, response.CodeServerBusy)
		return
	}
	response.ResSuccess(c, response.SignSuccess, "")
}
