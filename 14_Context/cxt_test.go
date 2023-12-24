package contextglwt_test

import (
	contextglwt "GoSayHello/14_Context"
	"context"
	"errors"
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

func (s *StubStore) Fetch(cxt context.Context) (string, error) {
	data := make(chan string, 1)

	// mocking 10ms delay for each char
	go func(cxt context.Context) {
		var result string
		for _, c := range s.res {
			select {
			case <-cxt.Done():
				s.Cancel()
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}(cxt)

	select {
	case <-cxt.Done():
		s.Cancel()
		return "", cxt.Err()
	case res := <-data:
		return res, nil
	}

}

func AssertString(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %s != %s", got, want)
	}
}

func AssertCancelled(t testing.TB, got, want bool) {
	if got != want {
		t.Errorf("got %t != %t", got, want)
	}
}

type SpyResponseRecorder struct {
	written bool
}

func (s *SpyResponseRecorder) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("TODO")
}

func (s *SpyResponseRecorder) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyResponseRecorder) Header() http.Header {
	s.written = true
	return nil
}

func TestServer(t *testing.T) {
	t.Run("Run server to return data", func(t *testing.T) {
		s := "This is the stub response"
		str := &StubStore{s, false}
		svr := contextglwt.Server(str)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()

		svr.ServeHTTP(resp, req)

		AssertString(t, resp.Body.String(), s)
		AssertCancelled(t, str.cancelled, false)

	})

	t.Run("Cancelled after 5 ms", func(t *testing.T) {
		str := &StubStore{"This is the stub response", false}
		svr := contextglwt.Server(str)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		cancelContext, cancelfunc := context.WithCancel(req.Context())
		req = req.WithContext(cancelContext)
		time.AfterFunc(5*time.Millisecond, cancelfunc)
		resp := &SpyResponseRecorder{}

		svr.ServeHTTP(resp, req)

		if resp.written {
			t.Error("kyu likh dia?")
		}

		AssertCancelled(t, str.cancelled, true)
	})

}
