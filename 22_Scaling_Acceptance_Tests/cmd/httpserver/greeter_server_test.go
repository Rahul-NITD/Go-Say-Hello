package main_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/Rahul-NITD/scalingacceptancetests/adapters"
	"github.com/Rahul-NITD/scalingacceptancetests/adapters/httpserver"
	"github.com/Rahul-NITD/scalingacceptancetests/specs"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		DockerFilePath = "./DockerFile"
		baseURL        = "http://localhost:" + port
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{Timeout: 1 * time.Second}}
		build_bin      = "httpserver"
	)
	adapters.StartDockerServer(t, port, build_bin, DockerFilePath)
	specs.GreeterSpecification(t, driver)
}
