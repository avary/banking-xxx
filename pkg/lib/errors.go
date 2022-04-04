package lib

import (
	"fmt"
	"net/http"
)

type RestErr interface {
	AsMessage() string
	AsStatus() int
	AsCauses() []interface{}
}

type restErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Causes  []interface{} `json:"causes"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - causes: %v",
		e.Message, e.Status, e.Causes)
}

func (e restErr) AsMessage() string {
	return e.Message
}

func (e restErr) AsStatus() int {
	return e.Status
}

func (e restErr) AsCauses() []interface{} {
	return e.Causes
}

func NewRestError(message string, status int, causes []interface{}) RestErr {
	return restErr{
		Message: message,
		Status:  status,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		Message: message,
		Status:  http.StatusNotFound,
	}
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		Message: message,
		Status:  http.StatusUnauthorized,
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		Message: message,
		Status:  http.StatusInternalServerError,
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
