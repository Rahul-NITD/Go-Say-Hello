package context

import (
	"fmt"
	"net/http"
)

// Context allows us to cancel a resource if parent dies

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}
