package main

import (
	"testing"
)

func Test_get_pokemon_description(t *testing.T) {
	pokemon := NewPokemonAPI()

	want := "When several of\nthese POKéMON\ngather, their\felectricity could\nbuild and cause\nlightning storms."
	got, err := pokemon.Describe("pikachu")
	if err != nil {
		t.Fatal("Wasnt expecting an error but got one ", err)
	}

	if got != want {
		t.Fatalf("Want %s \n, got %s\n", want, got)
	}
}
