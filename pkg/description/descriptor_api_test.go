package description

import (
	"testing"
)

func Test_get_pokemon_description(t *testing.T) {
	descript := NewDescriptor()
	got, err := descript.Describe("pikachu")
	want := "When several of these POKéMON gather, their\felectricity could build and cause lightning storms."

	if err != nil {
		t.Errorf("wasnt expecting an error got one : %v\n", err)
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func Test_get_invalid_pokemon_name(t *testing.T) {
	descript := NewDescriptor()
	_, err := descript.Describe("doesnt exist")
	if err == nil {
		t.Error("expecting an error but error is nil\n")
	}
}
