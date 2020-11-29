package pokemon

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hrishin/pokemon-shakespeare/pkg/description"
	"github.com/hrishin/pokemon-shakespeare/pkg/response"
	"github.com/hrishin/pokemon-shakespeare/pkg/translation"
)

func GetDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: exaplain why we used mux to handle such path variable
	vars := mux.Vars(r)
	name := vars["name"]
	//TODO: name validation
	//Scope: no empty

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
