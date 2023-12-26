package poker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

type PokerStorage interface {
	GetScore(player string) (int, error)
	RecordWin(player string) error
	GetLeague() []Player
}

type Player struct {
	Name string
	Wins int
}

type PokerServer struct {
	ScoreStorage PokerStorage
	game         Game
	http.Handler
}

func NewServer(storage PokerStorage, game Game) *PokerServer {
	server := new(PokerServer)
	server.ScoreStorage = storage
	router := http.NewServeMux()
	server.game = game
	router.Handle("/players/", http.HandlerFunc(server.playersRouteHandler))
	router.Handle("/league", http.HandlerFunc(server.leagueRouteHandler))
	router.Handle("/game", http.HandlerFunc(server.gameHandler))
	router.Handle("/ws", http.HandlerFunc(server.wsHandler))
	server.Handler = router
	return server
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (server *PokerServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws := NewPlayerWSServer(w, r)

	numberOfPlayersMsg := ws.WaitForMessage()
	numberOfPlayers, _ := strconv.Atoi(numberOfPlayersMsg)
	server.game.Start(numberOfPlayers, ws)

	winner := ws.WaitForMessage()
	server.game.Finish(winner)
}

type playerServerWS struct {
	*websocket.Conn
}

func (ws *playerServerWS) Write(p []byte) (n int, err error) {
	err = ws.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func NewPlayerWSServer(w http.ResponseWriter, r *http.Request) *playerServerWS {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("ws connection err, %v\n", err)
	}
	return &playerServerWS{conn}
}

func (p *playerServerWS) WaitForMessage() string {
	_, msg, err := p.ReadMessage()
	if err != nil {
		log.Printf("ws read err, %v\n", err)
	}
	return string(msg)
}

func (server *PokerServer) playersRouteHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		getScore(w, player, server.ScoreStorage)
	case http.MethodPost:
		w.WriteHeader(http.StatusAccepted)
		recordWin(w, player, server.ScoreStorage)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (server *PokerServer) leagueRouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getLeague(w, server.ScoreStorage)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (server *PokerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("game.html")

	if err != nil {
		http.Error(w, fmt.Sprintf("problem loading template %s", err.Error()), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func getLeague(w http.ResponseWriter, storage PokerStorage) {
	playersLeague := storage.GetLeague()
	b := bytes.Buffer{}
	json.NewEncoder(&b).Encode(playersLeague)
	w.Write(b.Bytes())
}

func getScore(w http.ResponseWriter, player string, storage PokerStorage) {
	score, err := storage.GetScore(player)
	switch err {
	case nil:
		fmt.Fprintf(w, "%d", score)
	case ERRORPlayerNotFound:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Player Not Found")
	default:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Player Not Found")
	}
}

func recordWin(w http.ResponseWriter, player string, storage PokerStorage) {
	err := storage.RecordWin(player)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
