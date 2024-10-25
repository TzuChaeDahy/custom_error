package custom_error

import (
	"net/http"
	"time"
)

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type RestErr struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Code      int       `json:"code"`
	Causes    []Cause   `json:"causes,omitempty"`
}

func (re *RestErr) Error() string {
	return re.Message
}

func NewRestErr(message string, code int, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Timestamp: time.Now().UTC(),
		Code:    code,
		Causes:  causes,
	}
}

func NewRestBadRequestError(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Timestamp: time.Now().UTC(),
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewRestUnauthorizedError(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Timestamp: time.Now().UTC(),
		Code:    http.StatusUnauthorized,
		Causes:  causes,
	}
}

func NewRestNotFoundError(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Timestamp: time.Now().UTC(),
		Code:    http.StatusNotFound,
		Causes:  causes,
	}
}

func NewRestInternalServerError(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Timestamp: time.Now().UTC(),
		Code:    http.StatusInternalServerError,
		Causes:  causes,
	}
}
