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

func (sr *ServiceResponse) WriteErrorTo(w http.ResponseWriter) {
	statuCode := http.StatusInternalServerError
	if sr.ErroCode != 0 {
		statuCode = sr.ErroCode
	}
	w.WriteHeader(statuCode)
	w.Write([]byte(sr.Error.Error()))
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

func (ar *APIResponse) SendReponseTO(w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(ar)
}
