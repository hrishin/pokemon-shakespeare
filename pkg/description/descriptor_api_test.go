package description

import (
	"net/http"
	"testing"

	"github.com/hrishin/pokemon-shakespeare/pkg/httpmock"
)

func Test_fetch_valid_resource_description(t *testing.T) {
	t.Parallel()
	given := "pikachu"
	mockResponse := httpmock.MockResponse{
		StatusCode: 200,
		URI:        "/pokemon-species/pikachu",
		Body: `{
			"flavor_text_entries": [{
				"flavor_text": "When several of these POKéMON gather, their\felectricity could build and cause lightning storms.",
				"language": {
					"name": "en"
				},
				"version": {
					"name": "ruby"
				}
			}]
		}`,
	}

	descriptor := &Descriptor{
		client: httpmock.MockClient(mockResponse),
		APIURL: pokeAPIURL,
	}
	got := descriptor.DescribePokemon(given)
	want := "When several of these POKéMON gather, their\felectricity could build and cause lightning storms."

	if got.Error != nil {
		t.Fatalf("wasnt expecting an error got one : %v\n", got.Error)
	}

	if got.Content != want {
		t.Fatalf("got %s, want %s", got.Content, want)
	}
}

func Test_fetch_invalid_resource_description(t *testing.T) {
	t.Parallel()
	given := "invalid_pokemon"
	mockResponse := httpmock.MockResponse{
		StatusCode: 404,
		Body:       "Not Found",
	}

	descriptor := &Descriptor{
		client: httpmock.MockClient(mockResponse),
		APIURL: pokeAPIURL,
	}
	got := descriptor.DescribePokemon(given)

	wantErrorCode := http.StatusNotFound
	if got.ErrorCode != wantErrorCode {
		t.Errorf("expecting an error code %d but got none : %d \n", wantErrorCode, got.Error)
	}

	wantErrorMessage := "failed to retrieve pokemon resource invalid_pokemon (code: 404)"
	if got.Error.Error() != wantErrorMessage {
		t.Errorf("expecting an error %s but got %s \n", wantErrorMessage, got.Error.Error())
	}
}
