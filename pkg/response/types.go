package response

import (
	"encoding/json"
	"net/http"
)

type ServiceResponse struct {
	Content  string
	Error    error
	ErroCode int
}

func NewError(err error) *ServiceResponse {
	return &ServiceResponse{
		Error: err,
	}
}

func NewErrorCode(code int, err error) *ServiceResponse {
	return &ServiceResponse{
		ErroCode: code,
		Error:    err,
	}
}

func NewSuccess(content string) *ServiceResponse {
	return &ServiceResponse{Content: content}
}

func (sr *ServiceResponse) ToErrorResonse() *ErrorRsponse {
	statuCode := http.StatusInternalServerError
	if sr.ErroCode != 0 {
		statuCode = sr.ErroCode
	}
	return &ErrorRsponse{Error: sr.Error.Error(), Code: statuCode}
}

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

func (ar *APIResponse) SendReponseTO(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ar)
}

type ErrorRsponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func (er *ErrorRsponse) WriteErrorTo(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(er.Code)
	json.NewEncoder(w).Encode(er)
}
