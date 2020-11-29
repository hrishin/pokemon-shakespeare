package description

import (
	"net/http"
	"testing"

	"github.com/hrishin/pokemon-shakespeare/pkg/httpmock"
)

//TODO: table drive test
func Test_get_pokemon_description(t *testing.T) {
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
		Client: httpmock.MockClient(mockResponse),
		APIURL: "https://pokeapi.co/api/v2/",
	}
	got := descriptor.DescribePokemon(given)
	want := "When several of these POKéMON gather, their\felectricity could build and cause lightning storms."

	//TODO: decide fatal vs error
	if got.Error != nil {
		t.Errorf("wasnt expecting an error got one : %v\n", got.Error)
	}

	if got.Content != want {
		t.Errorf("got %s, want %s", got.Content, want)
	}
}

func Test_description_for_invalid_pokemon_name(t *testing.T) {
	t.Parallel()
	given := "invalid_pokemon"
	mockResponse := httpmock.MockResponse{
		StatusCode: 404,
		URI:        "/pokemon-species/" + given,
		Body:       "Not Found",
	}

	descriptor := &Descriptor{
		Client: httpmock.MockClient(mockResponse),
		APIURL: "https://pokeapi.co/api/v2/",
	}
	got := descriptor.DescribePokemon(given)

	wantErrorCode := http.StatusNotFound
	if got.ErroCode != wantErrorCode {
		t.Errorf("expecting an error code %d but got none : %d \n", wantErrorCode, got.Error)
	}

	wantErrorMessage := "Not Found"
	if got.Error.Error() != wantErrorMessage {
		t.Errorf("expecting an error %s but got %s \n", wantErrorMessage, got.Error.Error())
	}
}
