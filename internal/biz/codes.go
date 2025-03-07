package biz

const OK = 200

var (
	DBERROR              = NewError(10000, "数据库错误")
	AlredyRegister       = NewError(10100, "用户已存在")
	InvalidParam         = NewError(10200, "参数错误")
	InvalidUser          = NewError(10300, "用户名或密码错误")
	GenerateCaptchaError = NewError(10400, "获取验证码错误")
	InvalidCaptcha       = NewError(10500, "验证码错误")
	RegisterError        = NewError(10600, "注册失败")
	GenerateTokenError   = NewError(10700, "生成token错误")
	RedisError           = NewError(10800, "Redis错误")
)
