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

func (fl flavorTextEntry) format() string {
	//TODO: decompose into small string
	//TODO: fix string formatting
	return strings.ReplaceAll(fl.FlavorText, "\n", " ")
}

type name struct {
	Name string `json:"name"`
}

func (s speciesResponse) descriptionFor(language, version string) (string, error) {
	for _, fl := range s.FlavorTextEntries {
		if fl.macth(language, version) {
			return fl.format(), nil
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
	if err != nil {
		return response.NewError(err)
	}

	desc, err := species.descriptionFor(defaultEnglish, pokeAPIVersion)
	if err != nil {
		return response.NewError(err)
	}

	return response.NewSuccess(desc)
}
