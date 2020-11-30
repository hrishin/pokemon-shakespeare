package httpmock

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

// MockResponse is to mock the HTTP response
// to use with the MockClient
type MockResponse struct {
	URI        string
	Body       string
	StatusCode int
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// MockClient create the mock http.client type
// to mock HTTP response with a given MockResponse type
func MockClient(response MockResponse) *http.Client {
	fn := func(req *http.Request) *http.Response {
		if strings.Contains(req.URL.String(), response.URI) == false {
			return &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(bytes.NewBufferString("Not found")),
				Header:     make(http.Header),
			}
		}

		return &http.Response{
			StatusCode: response.StatusCode,
			Body:       ioutil.NopCloser(strings.NewReader(response.Body)),
			Header:     make(http.Header),
		}
	}

	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
