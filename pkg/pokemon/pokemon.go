package pokemon

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hrishin/pokemon-shakespeare/pkg/description"
	"github.com/hrishin/pokemon-shakespeare/pkg/response"
	"github.com/hrishin/pokemon-shakespeare/pkg/translation"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("pokemon")

func GetDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	//One of the reason to mux handler is to extract the path variables easily
	vars := mux.Vars(r)
	name := vars["name"]
	log.Infof("getting the transolation for %s poken", name)

	de := description.NewDescriptor()
	desc := de.DescribePokemon(name)
	if desc.Error != nil {
		desc.ToErrorResonse().WriteErrorTo(w)
		return
	}

	tr := translation.NewTranslator()
	trans := tr.Translate(desc.Content)
	if trans.Error != nil {
		trans.ToErrorResonse().WriteErrorTo(w)
		return
	}

	log.Infof("responding the transolation %s for %s poken", trans.Content, name)
	response.NewAPIResponse(name, trans.Content).SendReponseTO(w)
}
