package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type API struct {
	Client  *http.Client
	BaseURL string
}

func NewPokemonAPI() *API {
	return &API{
		Client:  &http.Client{},
		BaseURL: "https://pokeapi.co/api/v2/",
	}
}

type pokemon struct {
	Species *struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"species"`
}

type descrptions struct {
	Flavors []struct {
		Text     string `json:"flavor_text"`
		Language struct {
			Name string `json:"name"`
		} `json:"language"`
		Version struct {
			Name string `json:"name"`
		} `json:"version"`
	} `json:"flavor_text_entries"`
}

func (p *API) Describe(name string) (string, error) {
	r, err := p.Client.Get(fmt.Sprintf("%s/pokemon/%s", p.BaseURL, name))
	if err != nil {
		return "", err
	}

	defer r.Body.Close()
	var pokemon pokemon
	err = json.NewDecoder(r.Body).Decode(&pokemon)
	if err != nil {
		return "", err
	}

	resp, err := p.Client.Get(pokemon.Species.Url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var description descrptions
	err = json.NewDecoder(resp.Body).Decode(&description)
	if err != nil {
		return "", err
	}

	for _, fl := range description.Flavors {
		if fl.Language.Name == "en" {
			return fl.Text, nil
		}
	}
	return "", errors.New("No description found")
}
