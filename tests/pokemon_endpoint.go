package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hrishin/pokemon-shakespeare/pkg/pokemon"
)

var _ = Describe("Pokemon endpoint tests", func() {
	Context("Get Shakespeare's pokemons description", func() {
		It("for charizard pokemon", func() {
			rr, err := pokemonDescription("charizard")
			response := strings.TrimSuffix(rr.Body.String(), "\n")

			Expect(err).To(BeNil(), "Error")
			Expect(rr.Code).To(Equal(http.StatusOK), "Status code")
			Expect(response).To(Equal(`{"name":"charizard","description":"Charizard flies 'round the sky in search of powerful opponents. 't breathes fire of such most wondrous heat yond 't melts aught. However,  't nev'r turns its fiery breath on any opponent weaker than itself."}`), "Response message")
		})

		It("for a invlid pokemon", func() {
			rr, err := pokemonDescription("invlid")
			response := strings.TrimSuffix(rr.Body.String(), "\n")

			Expect(err).To(BeNil(), "Error")
			Expect(rr.Code).To(Equal(http.StatusNotFound), "Status code")
			Expect(response).To(Equal(`{"error":"failed to retrieve pokemon resource invlid (code: 404)","code":404}`), "Response message")
		})
	})
})

func pokemonDescription(pkmn string) (*httptest.ResponseRecorder, error) {
	endpoint := "/pokemon/"
	req, err := http.NewRequest(http.MethodGet, endpoint+pkmn, nil)
	if err != nil {
		return nil, err
	}

	router := mux.NewRouter()
	router.HandleFunc(endpoint+"{name}", pokemon.GetDescriptionHandler)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr, nil
}
