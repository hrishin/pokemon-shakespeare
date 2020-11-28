package description

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

type speciesResponse struct {
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Version struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"flavor_text_entries"`
}

func (s *speciesResponse) flavorFor(lang, version string) (string, error) {
	for _, fl := range s.FlavorTextEntries {
		if fl.Language.Name == lang && fl.Version.Name == version {
			return strings.ReplaceAll(fl.FlavorText, "\n", " "), nil
		}
	}

	return "", errors.New("No pokemon description found")
}

func (d *Descriptor) Describe(name string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, d.APIURL+"pokemon-species/"+name, nil)
	if err != nil {
		return "", err
	}

	resp, err := d.Client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		return string(body), errors.New(fmt.Sprintf("%d", resp.StatusCode))
	}

	var species speciesResponse
	err = json.Unmarshal(body, &species)

	return species.flavorFor("en", "red")
}
