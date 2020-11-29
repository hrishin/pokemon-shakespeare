package description

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hrishin/pokemon-shakespeare/pkg/response"
)

type speciesResponse struct {
	FlavorTextEntries []flavorTextEntry `json:"flavor_text_entries"`
}

type flavorTextEntry struct {
	FlavorText string `json:"flavor_text"`
	Language   name   `json:"language"`
	Version    name   `json:"version"`
}

func (fl flavorTextEntry) macth(language, version string) bool {
	if fl.Language.Name == language && fl.Version.Name == version {
		return true
	}
	return false
}

func (fl flavorTextEntry) formatted() string {
	return strings.ReplaceAll(fl.FlavorText, "\n", " ")
}

type name struct {
	Name string `json:"name"`
}

func (s speciesResponse) descriptionFor(language, version string) (string, error) {
	for _, fl := range s.FlavorTextEntries {
		if fl.macth(language, version) {
			return fl.formatted(), nil
		}
	}
	return "", errors.New("No pokemon description found")
}

type Descriptor struct {
	Client *http.Client
	APIURL string
}

const (
	pokeAPI        = "https://pokeapi.co/api/v2/"
	defaultEnglish = "en"
	pokeAPIVersion = "ruby"
)

func NewDescriptor() *Descriptor {
	return &Descriptor{
		Client: &http.Client{},
		APIURL: pokeAPI,
	}
}

// string, customeErr -> string, error code
func (d *Descriptor) DescribePokemon(resource string) *response.ServiceResponse {
	url = fmt.Sprintf(d.APIURL + "pokemon-species/" + resource)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		//TODO: add logs here and other places
		return response.NewError(err)
	}

	resp, err := d.Client.Do(req)
	if err != nil {
		return response.NewError(err)
	}

	defer resp.Body.Close()
	//dues to relatively small response size reading respoonse without doing the buffered I/O (buffio)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response.NewError(err)
	}

	if resp.StatusCode >= 500 {
		err := fmt.Errorf("internal server error (code: %d)", resp.StatusCode)
		return response.NewErrorCode(resp.StatusCode, err)
	} else if resp.StatusCode >= 400 {
		err := fmt.Errorf("failed to retrieve pokemon resource %s (code: %d)", resource, resp.StatusCode)
		return response.NewErrorCode(resp.StatusCode, err)
	}

	var species speciesResponse
	err = json.Unmarshal(body, &species)
	if err != nil {
		return response.NewError(err)
	}

	desc, err := species.descriptionFor(defaultEnglish, pokeAPIVersion)
	if err != nil {
		return response.NewError(err)
	}

	return response.NewSuccess(desc)
}
