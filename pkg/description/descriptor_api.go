package description

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hrishin/pokemon-shakespeare/pkg/response"
)

//TODO: decompose into smaller structs
type speciesResponse struct {
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
		} `json:"language"`
		Version struct {
			Name string `json:"name"`
		} `json:"version"`
	} `json:"flavor_text_entries"`
}

//TODO: return string, err instead of reponse
//TODO: rename flavorFor to describe
func (s speciesResponse) flavorFor(lang, version string) *response.ServiceResponse {
	for _, fl := range s.FlavorTextEntries {
		//TODO: explain using comment
		if fl.Language.Name == lang && fl.Version.Name == version {
			//TODO: decompose into small string
			//TODO: fix string formatting
			return response.NewSuccess(strings.ReplaceAll(fl.FlavorText, "\n", " "))
		}
	}

	return response.NewError(errors.New("No pokemon description found"))
}

type Descriptor struct {
	Client *http.Client
	APIURL string
}

const (
	pokeAPI = "https://pokeapi.co/api/v2/"
)

func NewDescriptor() *Descriptor {
	return &Descriptor{
		Client: &http.Client{},
		APIURL: pokeAPI,
	}
}

// string, customeErr -> string, error code
func (d *Descriptor) DescribePokemon(name string) *response.ServiceResponse {
	//TODO: form the url package join
	req, err := http.NewRequest(http.MethodGet, d.APIURL+"pokemon-species/"+name, nil)
	if err != nil {
		//TODO: add logs here and other places
		return response.NewError(err)
	}

	resp, err := d.Client.Do(req)
	if err != nil {
		return response.NewError(err)
	}

	defer resp.Body.Close()
	// dues to small response size reading respoonse without using buffered I/O
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response.NewError(err)
	}

	//TODO: handle 500 case
	if resp.StatusCode >= 400 {
		return response.NewErrorCode(resp.StatusCode, errors.New(string(body)))
	}

	var species speciesResponse
	err = json.Unmarshal(body, &species)
	//TODO: error handling :P

	//TODO: constants e.g. pokeAPIVersion = ruby, dafaultLanguage=en
	return species.flavorFor("en", "ruby")
}
