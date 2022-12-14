package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SuperNerdIT/telepathic/configs"
)

func TestHTTPServer(t *testing.T) {
	cfg := configs.MakeConfigs()
	handler := NewServer(cfg).Handler

	ts := httptest.NewServer(handler)
	defer ts.Close()

	newReq := func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	testCases := map[string]struct {
		request      *http.Request
		expectedBody string
	}{
		"Health endpoint is up": {
			request:      newReq(http.MethodGet, ts.URL+"/health", nil),
			expectedBody: "{\"Ok\":true}", // remove and check status different to 404 for make the test more generic
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tc.request)

			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			str := strings.TrimSpace(string(b))
			if str != tc.expectedBody {
				t.Errorf("want '%s', got '%s'", tc.expectedBody, str)
			}
		})
	}
}
