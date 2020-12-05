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

// Translator provides an API to translate the given words
// to Shakespeare's style words using funtranslations API
type Translator struct {
	client *http.Client
	APIURL string
	APIKey string
}

// NewTranslator is factory method to create the Translator instance
func NewTranslator() *Translator {
	return &Translator{
		client: &http.Client{},
		APIURL: funTransAPIURL,
		APIKey: os.Getenv("TRANSLATION_API_KEY"),
	}
}

func (t *Translator) requestShakespeare(text string) (*http.Request, error) {
	data := map[string]string{"text": text}
	post, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%sshakespeare.json", t.APIURL), bytes.NewBuffer(post))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	if t.APIKey != "" {
		req.Header.Set("X-Funtranslations-Api-Secret", t.APIKey)
	}

	return req, nil
}

// Translate translates the words in to Shakespear style words
func (t *Translator) Translate(text string) *response.ServiceResponse {
	req, err := t.requestShakespeare(text)
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
		log.Errorf("error executing http request for funtranlation API")
		t.logResponseError(body)
		err = fmt.Errorf("internal server error occurred (code: %d)", http.StatusInternalServerError)
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

func (t *Translator) logResponseError(body []byte) {
	var errResp errorResponse
	err := json.Unmarshal(body, &errResp)
	if err == nil {
		log.Errorf("response from funtranlation service (code :%d): %s", errResp.Error.Code, errResp.Error.Message)
	}
}
