package errcode

import (
	"common_service/pkg/translation"
	"net/http"
)

type ApiError struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

func NewApiError(code int, msg string) *ApiError {
	return &ApiError{
		code: code,
		msg:  msg,
	}
}

func (e *ApiError) Error() string {
	return e.msg
}

func (e *ApiError) WithDetails(details ...string) *ApiError {
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}
	return e
}

func (e *ApiError) WithValidErrs(errs map[string]string) *ApiError {
	e.details = []string{}
	for k, v := range errs {
		m := k + ": " + v
		e.details = append(e.details, m)
	}
	return e
}

func (e *ApiError) StatusCode() int {
	switch e.code {
	case InternalError.code:
		return http.StatusInternalServerError
	default:
		return http.StatusOK
	}
}

func (e *ApiError) JSON() map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = e.code
	//res["msg"], _ = global.Trans.T(e.msg)
	res["msg"] = translation.TransT(e.msg)

	if e.details == nil {
		res["details"] = make([]string, 0)
	} else {
		res["details"] = e.details
	}

	return res
}
