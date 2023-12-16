package contextglwt

import (
	"context"
	"fmt"
	"net/http"
)

// Context allows us to cancel a resource if parent dies

type Store interface {
	Fetch(cxt context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := store.Fetch(r.Context())
		fmt.Fprint(w, data)
	}
}
