package biz

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

func NewError(code int64, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

func (e *Error) Error() string {
	return e.Message
}
