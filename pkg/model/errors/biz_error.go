package errors

import "github.com/bingodfok/freshguard-boot/pkg/model/resp"

type BizError struct {
	code int
	msg  string
}

func NewBizError(code int, msg string) *BizError {
	return &BizError{code, msg}
}

func NewBizErrorCode(code resp.ResultCode) *BizError {
	return &BizError{
		code: code.Code,
		msg:  code.Message,
	}
}

func (e *BizError) Error() string {
	return e.msg
}
