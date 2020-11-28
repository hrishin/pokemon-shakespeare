package description

import (
	"testing"
)

func Test_get_pokemon_description(t *testing.T) {
	want := "When several of\nthese POKéMON\ngather, their\felectricity could\nbuild and cause\nlightning storms."
	got, err := Describe("pikachu")

	if err != nil {
		t.Errorf("wasnt expecting an error got one : %v\n", err)
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func Test_get_invalid_pokemon_name(t *testing.T) {
	_, err := Describe("doesnt exist")
	if err == nil {
		t.Error("expecting an error but error is nil\n")
	}
}
