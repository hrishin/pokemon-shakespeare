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
			name:       "for charizard",
			endpoint:   "/pokemon/",
			pokmon:     "charizard",
			wantStatus: 200,
		},
		{
			name:       "for a invalid pokemon",
			endpoint:   "/pokemon/",
			pokmon:     "foobar",
			wantStatus: 404,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.endpoint+tc.pokmon, nil)
			if err != nil {
				t.Fatal(err)
			}
			router := mux.NewRouter()
			router.HandleFunc(tc.endpoint+"{name}", GetDescriptionHandler)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.wantStatus {
				t.Errorf("unexpected error code returned: got %v want %v", status, http.StatusOK)
			}
		})
	}
}
