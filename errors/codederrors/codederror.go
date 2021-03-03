package codederrors

import (
	"fmt"
)

type CodedError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Err     error  `json:"-"`
}

func (codedError CodedError) Error() string {
	return fmt.Sprintf("Error : %s - %s", codedError.Code, codedError.Message)
}

func New(code string, message string) *CodedError {
	return &CodedError{Code: code, Message: message}
}

func Wrap(err error, code string, message string) *CodedError {
	return &CodedError{Code: code, Message: message, Err: err}
}
