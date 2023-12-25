package httpserver

import (
	"fmt"
	"net/http"

	scalingacceptancetests "github.com/Rahul-NITD/scalingacceptancetests/domain/interactions"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", Response(scalingacceptancetests.Greet))
	mux.HandleFunc("/curse", Response(scalingacceptancetests.Curse))
	return mux
}

func Response(f func(string) string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, f(name))
	}
}
