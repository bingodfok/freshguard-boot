package resp

var (
	// 基础响应状态码
	SuccessCode      = NewResultCode(200, "Success")
	BadRequestCode   = NewResultCode(400, "Bad Request")
	UnauthorizedCode = NewResultCode(401, "Unauthorized")
	ServerErrorCode  = NewResultCode(500, "Server Error")
	// 特殊响应状态码
	CaptchaErrorCode = NewResultCode(1001, "验证码错误")
	CaptchaExpCode   = NewResultCode(1002, "验证码过期")
	UserNotExistCode = NewResultCode(1003, "用户不存在")

	FridgeExistCode = NewResultCode(2001, "冰箱已经存在")
)

type ResultCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResultCode(code int, message string) ResultCode {
	return ResultCode{
		Code:    code,
		Message: message,
	}
}
