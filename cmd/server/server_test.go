package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPServer (t *testing.T) {

	ts := httptest.NewServer(NewServer().Handler)
	defer ts.Close()

	newReq := func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return r
	  }

	testCases := map[string]struct {
		request *http.Request
		expectedBody string
	}{
		"Health endpoint is up": {
			request: newReq(http.MethodGet, ts.URL + "/health", nil),
			expectedBody: "Ok",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func (t *testing.T) {
			resp, err := http.DefaultClient.Do(tc.request)
			
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(b) != tc.expectedBody {
				t.Errorf("want '%s', got '%s'", tc.expectedBody, string(b))
			}
		})
	}
}