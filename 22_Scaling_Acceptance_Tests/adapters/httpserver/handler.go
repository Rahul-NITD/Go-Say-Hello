package httpserver

import (
	"fmt"
	"net/http"
	"strings"

	scalingacceptancetests "github.com/Rahul-NITD/scalingacceptancetests/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	switch {
	case strings.Contains(r.URL.Path, "/curse"):
		fmt.Fprint(w, scalingacceptancetests.Curse(name))
	default:
		fmt.Fprint(w, scalingacceptancetests.Greet(name))
	}

}
