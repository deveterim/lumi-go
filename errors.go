package lumi

import "fmt"

type ErrorResponse struct {
	StatusCode int    `json:"statusCode,omitempty"`
	Message    string `json:"message"`
	Name       string `json:"name,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("mailifica: status %d: %s", e.StatusCode, e.Message)
}
