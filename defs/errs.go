package defs

import "net/http"

type Result struct {
	Code   string      `json:"code"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

type Response struct {
	HttpSC int
	Result Result
}

var (
	ErrorRequestBodyParseFailed = Response{HttpSC: http.StatusBadRequest, Result: Result{Code: "001", ErrMsg: "Request body is not correct"}}
	ErrorNotAuthUser            = Response{HttpSC: http.StatusUnauthorized, Result: Result{Code: "002", ErrMsg: "User authentication failed."}}
	ErrorRedisError             = Response{HttpSC: http.StatusInternalServerError, Result: Result{Code: "003", ErrMsg: "Redis ops failed"}}
	ErrorInternalFaults         = Response{HttpSC: http.StatusInternalServerError, Result: Result{Code: "004", ErrMsg: "Internal service error"}}
	ErrorTooManyRequests        = Response{HttpSC: http.StatusTooManyRequests, Result: Result{Code: "005", ErrMsg: "Too many Request"}}
	ErrorParameterValidate      = Response{HttpSC: http.StatusBadRequest, Result: Result{Code: "006", ErrMsg: "validate parameters failed"}}
	ErrorStorageError           = Response{HttpSC: http.StatusInternalServerError, Result: Result{Code: "007", ErrMsg: "storage error"}}
)
