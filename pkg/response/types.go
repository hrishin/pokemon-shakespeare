package response

import (
	"encoding/json"
	"net/http"
)

// ServiceResponse to represent the response from
// the different services integrations
type ServiceResponse struct {
	Content   string
	Error     error
	ErrorCode int
}

// NewError creates the ServiceResponse with
// error value to represent only an error
func NewError(err error) *ServiceResponse {
	return &ServiceResponse{
		Error: err,
	}
}

// NewErrorCode creates the ServiceResponse with
// error value and code to represent only an error message and
// the error code from service integration responses.
func NewErrorCode(code int, err error) *ServiceResponse {
	return &ServiceResponse{
		ErrorCode: code,
		Error:     err,
	}
}

// NewSuccess creates the ServiceResponse to
// to represnet the response body from service integration
// response.
func NewSuccess(content string) *ServiceResponse {
	return &ServiceResponse{Content: content}
}

// ToErrorResponse create the ErrorResponse type to
// from the error fields of ServiceResponse. The ErrorResponse
// is used to represent the API error response.
func (sr *ServiceResponse) ToErrorResponse() *ErrorResponse {
	statusCode := http.StatusInternalServerError
	if sr.ErrorCode != 0 {
		statusCode = sr.ErrorCode
	}
	return &ErrorResponse{Error: sr.Error.Error(), Code: statusCode}
}

// APIResponse repsent the API response to return from
// Pokemon endpoints. Its fields are intended for JSON
// marshalling and unmarshalling purpose.
type APIResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewAPIResponse(name, description string) *APIResponse {
	return &APIResponse{
		Name:        name,
		Description: description,
	}
}

// SendResponseTO sends the APIResponse contents as a HTTP response
// in the JSON form with the appropriate HTTP status code.
func (ar *APIResponse) SendResponseTO(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ar)
}

// ErrorResponse repsent the API error to return from
// Pokemon endpoints. Its fields are intended for JSON
// marshalling and unmarshalling purpose.
type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// ErrorResponse sends the ErrorResponse contents as a HTTP response
// in the JSON form with the appropriate HTTP status code.
func (er *ErrorResponse) WriteErrorTo(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(er.Code)
	json.NewEncoder(w).Encode(er)
}
