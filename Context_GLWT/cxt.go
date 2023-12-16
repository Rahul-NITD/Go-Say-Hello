package contextglwt

import (
	"fmt"
	"net/http"
)

// Context allows us to cancel a resource if parent dies

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cxt := r.Context()
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-cxt.Done():
			store.Cancel()
		}

	}
}
