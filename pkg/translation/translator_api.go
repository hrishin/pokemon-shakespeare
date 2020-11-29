package translation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hrishin/pokemon-shakespeare/pkg/response"
	"github.com/op/go-logging"
)

const (
	funTransAPIURL  = "https://api.funtranslations.com/translate/"
	httpLimitExceed = 429
)

var log = logging.MustGetLogger("translator")

type translationResponse struct {
	Success  success  `json:"success"`
	Contents contents `json:"contents"`
}

type success struct {
	Total int `json:"total"`
}

type contents struct {
	Translated string `json:"translated"`
}

type errorResponse struct {
	Error err `json:"error"`
}

type err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type translator struct {
	client *http.Client
	apiURL string
	apiKey string
}

func NewTranslator() *translator {
	return &translator{
		client: &http.Client{},
		apiURL: funTransAPIURL,
		apiKey: os.Getenv("TRANSLATION_API_KEY"),
	}
}

func (t *translator) rquestShakespeare(text string) (*http.Request, error) {
	data := map[string]string{"text": text}
	post, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%sshakespeare.json", t.apiURL), bytes.NewBuffer(post))
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
		log.Errorf("error creating http request : %v", err)
		return response.NewError(err)
	}

	resp, err := t.client.Do(req)
	if err != nil {
		log.Errorf("error executing http request for the text %s : %v", text, err)
		return response.NewError(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("error reading http response for the text %s : %v", text, err)
		return response.NewError(err)
	}

	if resp.StatusCode >= 400 {
		log.Error("error executing http request")
		var errResp errorResponse
		err = json.Unmarshal(body, &errResp)
		if err == nil {
			log.Errorf("response from funtranlation service (code :%d): %s", errResp.Error.Code, errResp.Error.Message)
		}
		err = fmt.Errorf("internal server error (code: %d)", http.StatusInternalServerError)
		return response.NewErrorCode(http.StatusInternalServerError, err)
	}

	var transResp translationResponse
	err = json.Unmarshal(body, &transResp)
	if err != nil {
		log.Errorf("error unmarshalling translation response : %v", err)
		return response.NewError(err)
	}

	return response.NewSuccess(transResp.Contents.Translated)
}
