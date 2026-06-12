package exception

func NewBadRequest(format string, a ...interface{}) *ApiException {
	return NewApiException(CODE_BAD_REQUEST, codeReason(CODE_BAD_REQUEST)).WithMessagef(format, a...)
}

func NewUnauthorisation(format string, a ...any) *ApiException {
	return NewApiException(CODE_UNAUTHORIZED, codeReason(CODE_UNAUTHORIZED)).WithMessagef(format, a...)
}

func NewConflictRequest(format string, a ...any) *ApiException {
	return NewApiException(CODE_CONFLICT, codeReason(CODE_CONFLICT)).WithMessagef(format, a...)
}

func NewNotFoundRequest(format string, a ...any) *ApiException {
	return NewApiException(CODE_NOT_FOUND, codeReason(CODE_NOT_FOUND)).WithMessagef(format, a...)
}
