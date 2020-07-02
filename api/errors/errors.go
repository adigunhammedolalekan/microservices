package errors

type ApiError struct {
	Code    int
	Message string
}

func New(code int, message string) error {
	return &ApiError{Code: code, Message: message}
}

func (e *ApiError) Error() string {
	return e.Message
}
