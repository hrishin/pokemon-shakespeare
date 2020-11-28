package translation

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type response struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Translated string `json:"translated"`
	} `json:"contents"`
}

type errorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type ServiceResponse struct {
	Content  string
	Error    error
	ErroCode int
}

func NewResponseError(err error) *ServiceResponse {
	return &ServiceResponse{
		Error: err,
	}
}

func NewResponseErrorCode(code int, err error) *ServiceResponse {
	return &ServiceResponse{
		ErroCode: code,
		Error:    err,
	}
}

func NewRespnse(content string) *ServiceResponse {
	return &ServiceResponse{Content: content}
}

type Translator struct {
	Client *http.Client
	URL    string
	Key    string
}

func NewAPI(APIKey string) *Translator {
	return &Translator{
		Client: &http.Client{},
		URL:    "https://api.funtranslations.com/translate/",
		Key:    APIKey,
	}
}

func (a *Translator) rquestShakespeare(text string) (*http.Request, error) {
	data := map[string]string{"text": text}
	post, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, a.URL+"shakespeare.json", bytes.NewBuffer(post))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (a *Translator) Translate(text string) *ServiceResponse {
	req, err := a.rquestShakespeare(text)
	if err != nil {
		return NewResponseError(err)
	}

	resp, err := a.Client.Do(req)
	if err != nil {
		return NewResponseError(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NewResponseError(err)
	}

	if resp.StatusCode >= 400 {
		var errorResponse errorResponse
		err := json.Unmarshal(body, &errorResponse)
		if err != nil {
			return NewResponseError(err)
		}
		return NewResponseErrorCode(resp.StatusCode, errors.New(errorResponse.Error.Message))
	}

	var response response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return NewResponseError(err)
	}

	return NewRespnse(response.Contents.Translated)
}
