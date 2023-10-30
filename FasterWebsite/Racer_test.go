package fasterwebsite

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var assertStrings = func(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

var assertErrors = func(t testing.TB, got, want RacerErr) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRacer(t *testing.T) {
	t.Run("Test if returns faster website", func(t *testing.T) {
		slowServer := makeDelayedServer(9 * time.Millisecond)
		fastServer := makeDelayedServer(8 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		fast := fastServer.URL
		slow := slowServer.URL
		got, err := Racer(fast, slow)
		want := fast
		assertStrings(t, got, want)
		assertErrors(t, err, NIL)
	})
	t.Run("Test if Server reaches timeout", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(15 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		fast := fastServer.URL
		slow := slowServer.URL
		_, err := Racer(fast, slow)
		want := ErrServerTimeout
		assertErrors(t, err, want)
	})

}

func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(duration)
				w.WriteHeader(http.StatusOK)
			},
		),
	)
}
