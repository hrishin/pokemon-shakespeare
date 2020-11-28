package translation

import (
	"testing"

	"github.com/hrishin/pokemon-shakespeare/pkg/test"
)

func Test_translate_text(t *testing.T) {
	given := "When several of these POKéMON gather, their electricity could build and cause lightning storms."
	mockResponse := test.MockResponse{
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
	}

	Translator := &Translator{
		Client: test.MockClient(mockResponse),
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

	mockResponse := test.MockResponse{
		StatusCode: 429,
		URI:        "/shakespeare.json",
		Body: `{
			"error": {
				"message": "Too Many Requests: Rate limit of 5 requests per hour exceeded. Please wait for 59 minutes and 48 seconds."
			}
		}`,
	}

	Translator := &Translator{
		Client: test.MockClient(mockResponse),
		URL:    "https://api.funtranslations.com/translate/",
	}
	got := Translator.Translate(given)

	wantErrorCode := 429
	if got.ErroCode != wantErrorCode {
		t.Errorf("expecting an error code %d but got none : %d \n", wantErrorCode, got.Error)
	}

	wantErrorMessage := "Too Many Requests: Rate limit of 5 requests per hour exceeded. Please wait for 59 minutes and 48 seconds."
	if got.Error.Error() != wantErrorMessage {
		t.Errorf("expecting an error %s but got %s \n", wantErrorMessage, got.Error.Error())
	}
}