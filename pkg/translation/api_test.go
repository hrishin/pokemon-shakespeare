package translation

import (
	"testing"
)

func Test_translate_text(t *testing.T) {
	given := "When several of these POKéMON gather, their electricity could build and cause lightning storms."

	api := NewAPI("")
	got := api.Translate(given)

	if got.Error != nil {
		t.Errorf("wasnt expecting an error but got one : %v \n", got.Error)
	}
	want := "At which hour several of these pokémon gather,  their electricity couldst buildeth and cause lightning storms."

	if got.Content != want {
		t.Errorf("got: %s, want: %s", got.Content, want)
	}
}

func Test_translate_text_error(t *testing.T) {
	given := "When several of these POKéMON gather, their electricity could build and cause lightning storms."

	api := NewAPI("")
	got := api.Translate(given)

	if got.Error == nil {
		t.Errorf("expecting an error but got none : %v \n", got.Error)
	}

	if got.ErroCode == 0 {
		t.Errorf("expecting an error code but got none : %v \n", got.Error)
	}
}
