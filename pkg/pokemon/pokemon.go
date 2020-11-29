package pokemon

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hrishin/pokemon-shakespeare/pkg/description"
	"github.com/hrishin/pokemon-shakespeare/pkg/response"
	"github.com/hrishin/pokemon-shakespeare/pkg/translation"
)

func GetByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	//TODO: name validation

	de := description.NewDescriptor()
	desc := de.DescribePokemon(name)
	if desc.Error != nil {
		desc.WriteErrorTo(w)
		return
	}

	tr := translation.NewTranslator()
	trans := tr.Translate(desc.Content)
	if trans.Error != nil {
		trans.WriteErrorTo(w)
		return
	}

	response.NewAPIResponse(name, trans.Content).SendReponseTO(w)
}
