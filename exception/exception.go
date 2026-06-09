package exception

import "fmt"

type ApiException struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func NewApiException(code int, reason string) *ApiException {
	return &ApiException{
		Code:   code,
		Reason: reason,
	}
}

func (e *ApiException) WithMessagef(format string, a ...any) *ApiException {
	e.Message = fmt.Sprintf(format, a...)
	return e
}

func (e *ApiException) Error() string {
	msg := e.Reason
	if e.Message != "" {
		msg += ": " + e.Message
	}
	return msg
}
