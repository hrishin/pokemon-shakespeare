package pokemon

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_get_description(t *testing.T) {
	tt := []struct {
		name         string
		endpoint     string
		pokmon       string
		wantStatus   int
		wantResponse string
	}{
		{
			name:         "for charizard",
			endpoint:     "/pokemon/",
			pokmon:       "charizard",
			wantStatus:   200,
			wantResponse: `{ "description": "Charizard flies 'round the sky in search of powerful opponents. 't breathes fire of such most wondrous heat yond 't melts aught. However,  't nev'r turns its fiery breath on any opponent weaker than itself.", "name": "charizard" }`,
		},
		{
			name:         "for a invalid pokemon",
			endpoint:     "/pokemon/",
			pokmon:       "foobar",
			wantStatus:   404,
			wantResponse: "Not Found",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.endpoint+tc.pokmon, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc(tc.endpoint+"{name}", GetByName)
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.wantStatus {
				t.Errorf("unexpected error code returned: got %v want %v", status, http.StatusOK)
			}

			if rr.Body.String() != tc.wantResponse {
				t.Errorf("handler returned unexpected body: got %v want %x", rr.Body.String(), tc.wantResponse)
			}
		})
	}
}
