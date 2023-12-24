package httpserver

import (
	"fmt"
	"net/http"

	"github.com/Rahul-NITD/scalingacceptancetests"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, scalingacceptancetests.Greet(name))
}
