package main

import "fmt"

// ErrorMessage error struct
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewError new error
func NewError(code int, message string) ErrorMessage {
	return ErrorMessage{
		Code:    code,
		Message: message,
	}
}

// implements error interface
func (e ErrorMessage) Error() string {
	return fmt.Sprintf("code=[%d] message=[%s]", e.Code, e.Message)
}
