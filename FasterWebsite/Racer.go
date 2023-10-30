package fasterwebsite

import (
	"net/http"
	"time"
)

const (
	ErrServerTimeout = RacerErr("neither server responded within timeout")
	NIL              = RacerErr("no error, everything's fine!")
)

type RacerErr string

func (r RacerErr) Error() string {
	return string(r)
}

func Racer(url1, url2 string) (winner string, err RacerErr) {

	select {
	case <-ping(url1):
		return url1, NIL
	case <-ping(url2):
		return url2, NIL
	case <-time.After(10 * time.Millisecond):
		return "", ErrServerTimeout
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func(url string) {
		http.Get(url)
		close(ch)
	}(url)
	return ch
}
