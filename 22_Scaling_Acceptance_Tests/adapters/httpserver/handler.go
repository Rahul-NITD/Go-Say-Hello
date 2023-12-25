package httpserver

import (
	"fmt"
	"net/http"

	scalingacceptancetests "github.com/Rahul-NITD/scalingacceptancetests/domain/interactions"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", GreetResponse())
	mux.HandleFunc("/curse", CurseResponse())
	return mux
}

func GreetResponse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, scalingacceptancetests.Greet(name))
	}
}

func CurseResponse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, scalingacceptancetests.Curse(name))
	}
}
