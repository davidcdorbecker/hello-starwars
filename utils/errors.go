package utils

type error struct {
	ErrorMessage    string `json:"message"`
	ErrorCodeStatus int    `json:"code"`
}

type Error interface {
	Message() string
	CodeStatus() int
}

func NewRestError(message string, codeStatus int) Error {
	return &error{
		ErrorMessage:    message,
		ErrorCodeStatus: codeStatus,
	}
}

func (e *error) Message() string {
	return e.ErrorMessage
}

func (e *error) CodeStatus() int {
	return e.ErrorCodeStatus
}
