package poker_test

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	poker "github.com/Rahul-NITD/Poker"
	"github.com/gorilla/websocket"
)

func TestWebsocket(t *testing.T) {
	t.Run("receive winner over websocket, accept blind messages", func(t *testing.T) {
		// wantedBlindAlert := "Blind is 100"
		store := NewSTUBStorage()
		winner := "Rahul"
		game := poker.NewTexasHoldem(&SpyAlerter{}, &store)
		server := httptest.NewServer(poker.NewServer(&store, game))
		defer server.Close()

		conn := DialWS(t, server)

		defer conn.Close()

		WriteMessageWS(t, conn, "3")
		WriteMessageWS(t, conn, winner)
		done := retryUntil(500*time.Millisecond, func() bool {
			got, _ := store.GetScore(winner)
			return got == 3
		})
		if !done {
			t.Error("Score not updated")
		}
	})
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}

func DialWS(t testing.TB, server *httptest.Server) *websocket.Conn {
	t.Helper()
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("could not open ws connection, %v", err)
	}
	return conn
}

func WriteMessageWS(t testing.TB, conn *websocket.Conn, message string) {
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection, %v", err)
	}
}
