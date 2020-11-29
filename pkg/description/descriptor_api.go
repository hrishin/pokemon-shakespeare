package description

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hrishin/pokemon-shakespeare/pkg/response"
)

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

func (s *speciesResponse) flavorFor(lang, version string) *response.ServiceResponse {
	for _, fl := range s.FlavorTextEntries {
		if fl.Language.Name == lang && fl.Version.Name == version {
			return response.NewSuccess(strings.ReplaceAll(fl.FlavorText, "\n", " "))
		}
	}

	return response.NewError(errors.New("No pokemon description found"))
}

type Descriptor struct {
	Client *http.Client
	APIURL string
}

func NewDescriptor() *Descriptor {
	return &Descriptor{
		Client: &http.Client{},
		APIURL: "https://pokeapi.co/api/v2/",
	}
}

func (d *Descriptor) DescribePokemon(name string) *response.ServiceResponse {
	req, err := http.NewRequest(http.MethodGet, d.APIURL+"pokemon-species/"+name, nil)
	if err != nil {
		return response.NewError(err)
	}

	resp, err := d.Client.Do(req)
	if err != nil {
		return response.NewError(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response.NewError(err)
	}

	if resp.StatusCode >= 400 {
		return response.NewErrorCode(resp.StatusCode, errors.New(string(body)))
	}

	var species speciesResponse
	err = json.Unmarshal(body, &species)

	return species.flavorFor("en", "ruby")
}
