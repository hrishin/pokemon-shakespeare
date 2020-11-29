package description

import (
	"testing"

	"github.com/hrishin/pokemon-shakespeare/pkg/test"
)

func Test_get_pokemon_description(t *testing.T) {
	given := "invalid_pokemon"
	mockResponse := test.MockResponse{
		StatusCode: 200,
		URI:        "/pokemon-species/" + given,
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
		Client: test.MockClient(mockResponse),
		APIURL: "https://pokeapi.co/api/v2/",
	}
	got := descriptor.DescribePokemon(given)
	want := "When several of these POKéMON gather, their\felectricity could build and cause lightning storms."

	if got.Error != nil {
		t.Errorf("wasnt expecting an error got one : %v\n", got.Error)
	}

	if got.Content != want {
		t.Errorf("got %s, want %s", got.Content, want)
	}
}

func Test_description_for_invalid_pokemon_name(t *testing.T) {
	given := "invalid_pokemon"
	mockResponse := test.MockResponse{
		StatusCode: 404,
		URI:        "/pokemon-species/" + given,
		Body:       "Not Found",
	}

	descriptor := &Descriptor{
		Client: test.MockClient(mockResponse),
		APIURL: "https://pokeapi.co/api/v2/",
	}
	got := descriptor.DescribePokemon(given)

	wantErrorCode := 404
	if got.ErroCode != wantErrorCode {
		t.Errorf("expecting an error code %d but got none : %d \n", wantErrorCode, got.Error)
	}

	wantErrorMessage := "Not Found"
	if got.Error.Error() != wantErrorMessage {
		t.Errorf("expecting an error %s but got %s \n", wantErrorMessage, got.Error.Error())
	}
}
