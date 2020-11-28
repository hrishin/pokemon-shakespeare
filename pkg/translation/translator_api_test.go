package translation

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type MockResponse struct {
	URI        string
	Body       string
	StatusCode int
}

type MockResponses []MockResponse

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func MockClient(response MockResponses) *http.Client {
	return NewTestClient(func(req *http.Request) *http.Response {
		for _, r := range response {
			if strings.Contains(req.URL.String(), r.URI) == false {
				continue
			}

			return &http.Response{
				StatusCode: r.StatusCode,
				Body:       ioutil.NopCloser(strings.NewReader(r.Body)),
				Header:     make(http.Header),
			}
		}

		return &http.Response{
			StatusCode: 404,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Not found")),
			Header:     make(http.Header),
		}
	})
}

func Test_translate_text(t *testing.T) {
	given := "When several of these POKéMON gather, their electricity could build and cause lightning storms."
	mockResponse := MockResponses{
		MockResponse{
			StatusCode: 200,
			URI:        "/shakespeare.json",
			Body: `{
				"success": {
					"total": 1
				},
				"contents": {
					"translated": "At which hour several of these pokémon gather,  their electricity couldst buildeth and cause lightning storms."
				}
			}`,
		},
	}
	Translator := &Translator{
		Client: MockClient(mockResponse),
		URL:    "https://api.funtranslations.com/translate/",
	}

	got := Translator.Translate(given)

	if got.Error != nil {
		t.Errorf("wasnt expecting an error but got one : %v \n", got.Error)
	}
	want := "At which hour several of these pokémon gather,  their electricity couldst buildeth and cause lightning storms."

	if got.Content != want {
		t.Errorf("got: %s, want: %s", got.Content, want)
	}
}

func Test_translate_text_error(t *testing.T) {
	given := "When several of these POKéMON gather, their electricity could build and cause lightning storms."

	mockResponse := MockResponses{
		MockResponse{
			StatusCode: 429,
			URI:        "/shakespeare.json",
			Body: `{
				"error": {
					"message": "Too Many Requests: Rate limit of 5 requests per hour exceeded. Please wait for 59 minutes and 48 seconds."
				}
			}`,
		},
	}
	Translator := &Translator{
		Client: MockClient(mockResponse),
		URL:    "https://api.funtranslations.com/translate/",
	}
	got := Translator.Translate(given)

	wantErrorCode := 429
	wantErrorMessage := "Too Many Requests: Rate limit of 5 requests per hour exceeded. Please wait for 59 minutes and 48 seconds."

	if got.ErroCode != wantErrorCode {
		t.Errorf("expecting an error code but got none : %v \n", got.Error)
	}

	if got.Error.Error() != wantErrorMessage {
		t.Errorf("expecting an error but got none : %v \n", got.Error)
	}
}
