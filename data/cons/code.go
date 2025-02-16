package cons

import "errors"

var (
	InternalErr = errors.New("服务器内部错误")
)

const (
	InternalServerError = 50000
	InternalErrorMsg    = "服务器内部错误"

	BasicSuccessCode = 10000
	BasicSuccessMsg  = "success"

	VerifyCodeFailCode    = 40007
	VerifyCodeFailMsg     = "验证码错误"
	VerifyCodeExpiredCode = 40008
	VerifyCodeExpiredMsg  = "验证码已过期或未发送"

	BadRequestCode = 40000
	BadRequestMsg  = "请求参数错误"
)

const (
	UserIsExistMsg           = "用户已存在"
	UserIsExistCode          = 40009
	UserNotFoundMsg          = "用户不存在"
	UserNotFoundCode         = 40010
	UserWaitToRegisterMsg    = "已发送验证码,等待注册"
	UserWaitToRegisterCode   = 40011
	UserPasswordNotRightMsg  = "密码错误"
	UserPasswordNotRightCode = 40012
)
