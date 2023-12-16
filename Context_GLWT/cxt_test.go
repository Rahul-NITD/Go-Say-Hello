package contextglwt_test

import (
	contextglwt "GoSayHello/Context_GLWT"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	res       string
	cancelled bool
}

func (s *StubStore) Cancel() {
	s.cancelled = true
}

func (s *StubStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.res
}

func TestServer(t *testing.T) {
	t.Run("Run server to return data", func(t *testing.T) {
		s := "This is the stub response"
		str := &StubStore{s, false}
		svr := contextglwt.Server(str)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()

		svr.ServeHTTP(resp, req)

		if resp.Body.String() != s {
			t.Errorf("got %s != %s", resp.Body.String(), s)
		}

		if str.cancelled {
			t.Error("Should not be cancelled")
		}

	})

	t.Run("Cancelled after 5 ms", func(t *testing.T) {
		str := &StubStore{"This is the stub response", false}
		svr := contextglwt.Server(str)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		cancelContext, cancelfunc := context.WithCancel(req.Context())
		req = req.WithContext(cancelContext)
		time.AfterFunc(5*time.Millisecond, cancelfunc)
		resp := httptest.NewRecorder()

		svr.ServeHTTP(resp, req)

		if !str.cancelled {
			t.Error("request not cancelled")
		}
	})

}
