package scalingacceptancetests

import (
	"io"
	"net/http"
)

type Driver struct {
	BaseURL string
	Client  *http.Client
}

func (d Driver) Greet() (string, error) {
	res, err := d.Client.Get(d.BaseURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	msg, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(msg), nil
}
