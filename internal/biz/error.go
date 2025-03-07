package biz

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func NewError(code int64, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	return e.Message
}
