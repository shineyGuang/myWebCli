package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResError(c *gin.Context, code ResCode) {
	rd := &ResData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResErrorWithString(c *gin.Context, code ResCode, errMsg interface{}) {
	rd := &ResData{
		Code: code,
		Msg:  errMsg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResSuccess(c *gin.Context, code ResCode, data interface{}) {
	rd := &ResData{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
