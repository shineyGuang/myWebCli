package user

import (
	"errors"
	"fmt"
	"myWebCli/controller/response"
	validator2 "myWebCli/controller/validator"
	"myWebCli/db/mysql"
	"myWebCli/logic/SignUp"
	"myWebCli/logic/login"
	"myWebCli/models"
	jwt "myWebCli/utils/Jwt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func TestWebHandler(c *gin.Context) {
	response.ResSuccess(c, response.TestSuccess, "Welcome to my web!")
}

//LoginHandler 登录
func LoginHandler(c *gin.Context) {
	// 获取请求参数，并且校验
	var loginUser models.LoginForm
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			response.ResErrorWithString(c, response.ParamsValidatorError, errs.Translate(validator2.Trans))
			return
		} else {
			response.ResError(c, response.ParseJsonFailed)
			return
		}
	}
	err := login.Login(&loginUser)
	if err != nil {
		if errors.Is(err, mysql.ErrorPasswordWrong) {
			response.ResErrorWithString(c, response.LoginFailed, mysql.ErrorPasswordWrong.Error())
			return
		} else {
			response.ResErrorWithString(c, response.LoginFailed, err.Error())
			return
		}
	}
	// 生成token
	atoken, rtoken, _ := jwt.GenToken(loginUser.UserId)
	response.ResSuccess(c, response.LoginSuccess, gin.H{
		"atoken":   atoken,
		"rtoken":   rtoken,
		"user_id":  loginUser.UserId,
		"userName": loginUser.UserName,
	})
	return
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

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		response.ResErrorWithString(c, response.AuthNilError, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		response.ResErrorWithString(c, response.AuthInvalidated, "Token格式不对")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
