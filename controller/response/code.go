package response

type ResCode uint64

const (
	LoginSuccess         ResCode = 1000
	LoginFailed          ResCode = 1001
	SignSuccess          ResCode = 1002
	SignFailed           ResCode = 1003
	ParseJsonFailed      ResCode = 4001
	ParamsValidatorError ResCode = 5001
	TokenError           ResCode = 6001
	AuthNilError         ResCode = 1004
	AuthValidateError    ResCode = 1005
	AuthInvalidated      ResCode = 1006
	CodeServerBusy       ResCode = 9999
	TestSuccess          ResCode = 0001
	CodeUserExit         ResCode = 9001
)

var msgFlags = map[ResCode]string{
	LoginSuccess:         "success",
	LoginFailed:          "登录失败",
	SignSuccess:          "success",
	SignFailed:           "注册失败",
	ParseJsonFailed:      "数据解析失败",
	ParamsValidatorError: "参数校验失败",
	TokenError:           "token解析错误",
	AuthNilError:         "请求头中auth为空",
	AuthValidateError:    "请求头中auth格式有误",
	AuthInvalidated:      "无效的Token",
	CodeServerBusy:       "服务器繁忙",
	TestSuccess:          "测试成功",
	CodeUserExit:         "用户已存在",
}

func (c ResCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
