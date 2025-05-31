package errors

import "github.com/bingodfok/freshguard-boot/pkg/model/resp"

type BizError struct {
	Code int
	Msg  string
}

func NewBizError(code int, msg string) *BizError {
	return &BizError{code, msg}
}

func NewBizErrorCode(code resp.ResultCode) *BizError {
	return &BizError{
		Code: code.Code,
		Msg:  code.Message,
	}
}

func (e *BizError) Error() string {
	return e.Msg
}
