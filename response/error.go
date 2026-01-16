package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func InternalServerError(message string) Response {
	return errors(http.StatusInternalServerError, message)
}

func NotFound(message string) Response {
	return errors(http.StatusNotFound, message)
}

func BadRequest(message string) Response {
	return errors(http.StatusBadRequest, message)
}

func Unauthorized(message string) Response {
	return errors(http.StatusUnauthorized, message)
}

func Forbidden(message string) Response {
	return errors(http.StatusForbidden, message)
}

func errors(status int, message string) Response {
	return &ErrorResponse{
		Status:  status,
		Message: message,
	}
}

func (e *ErrorResponse) StatusCode() int {
	return e.Status
}

func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) GetData() interface{} {
	return nil
}
