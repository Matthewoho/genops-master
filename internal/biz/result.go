package biz

import "genops-master/internal/types"

func Success(data interface{}) *types.Result {
	return &types.Result{
		Code:    OK,
		Message: "Success",
		Data:    data,
	}

}

func Fail(err *Error) *types.Result {
	return &types.Result{
		Code:    err.Code,
		Message: err.Message,
	}
}
