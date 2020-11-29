package translation

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hrishin/pokemon-shakespeare/pkg/response"
)

type transolationResponse struct {
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

type translator struct {
	client *http.Client
	apiURL string
	//TODO: populate in env var
	apiKey string
}

func NewTranslator() *translator {
	return &translator{
		client: &http.Client{},
		apiURL: "https://api.funtranslations.com/translate/",
		apiKey: os.Getenv("TRANSLATION_API_KEY"),
	}

}

func (t *translator) rquestShakespeare(text string) (*http.Request, error) {
	data := map[string]string{"text": text}
	post, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, t.apiURL+"shakespeare.json", bytes.NewBuffer(post))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	if t.apiKey != "" {
		req.Header.Set("X-Funtranslations-Api-Secret", t.apiKey)
	}

	return req, nil
}

func (t *translator) Translate(text string) *response.ServiceResponse {
	req, err := t.rquestShakespeare(text)
	if err != nil {
		return response.NewError(err)
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return response.NewError(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response.NewError(err)
	}

	if resp.StatusCode >= 400 {
		var errorResponse errorResponse
		err := json.Unmarshal(body, &errorResponse)
		if err != nil {
			return response.NewError(err)
		}
		return response.NewErrorCode(resp.StatusCode, errors.New(errorResponse.Error.Message))
	}

	var transolationResp transolationResponse
	err = json.Unmarshal(body, &transolationResp)
	if err != nil {
		return response.NewError(err)
	}

	return response.NewSuccess(transolationResp.Contents.Translated)
}
