package errors

import (
	"net/http"
)

//RestErroAPI retorno das mensages de erro conforme o padrão rest
type RestErroAPI struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

//NewBadRequestError retorno das mensages de erro conforme o padrão rest
func NewBadRequestError(message string) *RestErroAPI {
	return &RestErroAPI{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

//NewInternalServerError retorno das mensages de erro conforme o padrão rest
func NewInternalServerError(message string) *RestErroAPI {
	return &RestErroAPI{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Error:      "internal_server_error",
	}
}

//NewNotFoundErro retorno das mensages de erro conforme o padrão rest
func NewNotFoundErro(message string) *RestErroAPI {
	return &RestErroAPI{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "not_found",
	}
}

//NewValidationError retorno das mensages de erro conforme o padrão rest
func NewValidationError(message string) *RestErroAPI {
	return &RestErroAPI{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
		Error:      "validation_error",
	}
}

//NewUnexpectedError retorno das mensages de erro conforme o padrão rest
func NewUnexpectedError(message string) *RestErroAPI {
	return &RestErroAPI{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Error:      "database_error",
	}
}
