package codes

import (
	"errors"
	"fmt"
)

func IsCodeError(err error) bool {
	var result bool
	if nil != err {
		var codeError *CodeError
		switch {
		case errors.As(err, &codeError):
			result = true
		default:
			result = false
		}
	} else {
		result = false
	}
	return result
}

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  error  `json:"err"`
}

func NewCodeErrorWitheRR(code *ResponseCode, err error) *CodeError {
	return &CodeError{Code: code.Code, Msg: code.Msg, Err: err}
}
func NewCodeError(code *ResponseCode) *CodeError {
	return &CodeError{Code: code.Code, Msg: code.Msg, Err: nil}
}

func (c *CodeError) Error() string {
	r := fmt.Sprintf("code : %d, msg : %s", c.Code, c.Msg)
	return r
}
