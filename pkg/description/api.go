package description

import (
	"errors"
	"fmt"

	pokeapi "github.com/mtslzr/pokeapi-go"
)

func Describe(name string) (string, error) {
	species, err := pokeapi.PokemonSpecies(name)
	if err != nil {
		return "", fmt.Errorf("Error occured in getting the pokemon : %v", err)
	}

	for _, sp := range species.FlavorTextEntries {
		if sp.Language.Name == "en" {
			return sp.FlavorText, nil
		}
	}

	return "", errors.New("No pokemon description found")
}
