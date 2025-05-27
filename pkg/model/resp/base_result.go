package resp

import "time"

type Result struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
}

func CommonResult(code int, msg string, data interface{}) *Result {
	return &Result{Code: code, Msg: msg, Timestamp: time.Now().UnixMilli(), Data: data}
}

func CodeResult(code ResultCode, data interface{}) *Result {
	return &Result{Code: code.Code, Msg: code.Message, Timestamp: time.Now().UnixMilli(), Data: data}
}

func CodeMsgResult(code ResultCode, msg string) *Result {
	return &Result{Code: code.Code, Msg: msg, Timestamp: time.Now().UnixMilli()}
}

func EmptyDataResult(code ResultCode) *Result {
	return &Result{Code: code.Code, Msg: code.Message, Timestamp: time.Now().UnixMilli()}
}

func Success(data interface{}) *Result {
	return CodeResult(SuccessCode, data)
}

func (r *Result) Ok() bool {
	return r.Code == 0
}
