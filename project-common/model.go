package common

import "net/http"

type BusinessCode int

type Result struct {
	Code BusinessCode `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}

// 响应成功
func (res *Result) Ok(data any) *Result {
	res.Code = http.StatusOK
	res.Msg = "suc"
	res.Data = data
	return res
}

// 响应失败
func (res *Result) Err(code BusinessCode, msg string) *Result {
	res.Code = code
	res.Msg = msg
	return res
}
