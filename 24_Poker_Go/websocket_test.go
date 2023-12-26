package poker_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	poker "github.com/Rahul-NITD/Poker"
	"github.com/gorilla/websocket"
)

func TestWebsocket(t *testing.T) {
	t.Run("receive winner over websocket", func(t *testing.T) {
		store := NewSTUBStorage()
		winner := "Rahul"
		server := httptest.NewServer(poker.NewServer(&store))
		defer server.Close()
		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			t.Fatalf("could not open ws connection, %v", err)
		}
		defer conn.Close()

		if err := conn.WriteMessage(websocket.TextMessage, []byte(winner)); err != nil {
			t.Fatalf("could not send message over ws connection, %v", err)
		}
		assertPlayerWin(t, &store, "Rahul", 3)

	})
}
