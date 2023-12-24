package httpserver

import (
	"fmt"
	"net/http"

	scalingacceptancetests "github.com/Rahul-NITD/scalingacceptancetests/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, scalingacceptancetests.Greet(name))
}
