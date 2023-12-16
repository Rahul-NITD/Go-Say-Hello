package context_test

import (
	context "GoSayHello/Context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	res string
}

// fetch implements context.Store.
func (s *StubStore) Fetch() string {
	return s.res
}

func TestServer(t *testing.T) {
	str := &StubStore{"This is the stub response"}
	svr := context.Server(str)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	svr.ServeHTTP(resp, req)

	if resp.Body.String() != "This is the stub response" {
		t.Errorf("got %s != %s", resp.Body.String(), "This is the stub response")
	}

}
