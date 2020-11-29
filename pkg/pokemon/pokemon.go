package pokemon

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hrishin/pokemon-shakespeare/pkg/description"
	"github.com/hrishin/pokemon-shakespeare/pkg/response"
	"github.com/hrishin/pokemon-shakespeare/pkg/translation"
)

func GetByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	desc := description.NewDescriptor()
	resp := desc.Describe(name)
	if resp.Error != nil {
		resp.WriteErrorTo(w)
		return
	}

	trans := translation.NewTranslator()
	resp = trans.Translate(resp.Content)
	if resp.Error != nil {
		resp.WriteErrorTo(w)
		return
	}

	response.NewAPIResponse(name, resp.Content).SendReponseTO(w)
}
