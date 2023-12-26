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
	t.Run("receive winner over websocket", func(t *testing.T) {
		store := NewSTUBStorage()
		winner := "Rahul"
		server := httptest.NewServer(poker.NewServer(&store))
		defer server.Close()

		conn := DialWS(t, server)

		defer conn.Close()

		WriteMessageWS(t, conn, winner)
		time.Sleep(10 * time.Millisecond)
		assertPlayerWin(t, &store, "Rahul", 3)

	})
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
