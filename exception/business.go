package exception

func NewBadRequest(format string, a ...interface{}) *ApiException {
	return NewApiException(CODE_BAD_REQUEST, codeReason(CODE_BAD_REQUEST)).WithMessagef(format, a...)
}
