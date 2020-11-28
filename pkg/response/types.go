package response

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
