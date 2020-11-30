package response

import (
	"encoding/json"
	"net/http"
)

// ServiceResponse to represent the response from
// the different services intefrations
type ServiceResponse struct {
	Content  string
	Error    error
	ErroCode int
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
// the error code from service intgration responses.
func NewErrorCode(code int, err error) *ServiceResponse {
	return &ServiceResponse{
		ErroCode: code,
		Error:    err,
	}
}

// NewSuccess creates the ServiceResponse to
// to represnet the response body from service integration
// response.
func NewSuccess(content string) *ServiceResponse {
	return &ServiceResponse{Content: content}
}

// ToErrorResonse create the ErrorRsponse type to
// from the error fields of ServiceResponse. The ErrorRsponse
// is used to represent the API error response.
func (sr *ServiceResponse) ToErrorResonse() *ErrorRsponse {
	statuCode := http.StatusInternalServerError
	if sr.ErroCode != 0 {
		statuCode = sr.ErroCode
	}
	return &ErrorRsponse{Error: sr.Error.Error(), Code: statuCode}
}

// APIResponse repsent the API response to return from
// Pokemkon endpoints. Its fields are intended for JSON
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

// SendReponseTO sends the APIResponse contents as a HTTP response
// in the JSON form with the appropriate HTTP status code.
func (ar *APIResponse) SendReponseTO(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ar)
}

// ErrorRsponse repsent the API error to return from
// Pokemkon endpoints. Its fields are intended for JSON
// marshalling and unmarshalling purpose.
type ErrorRsponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// ErrorRsponse sends the ErrorRsponse contents as a HTTP response
// in the JSON form with the appropriate HTTP status code.
func (er *ErrorRsponse) WriteErrorTo(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(er.Code)
	json.NewEncoder(w).Encode(er)
}
