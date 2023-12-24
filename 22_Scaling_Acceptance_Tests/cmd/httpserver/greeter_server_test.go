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
		DockerFilePath = "./Dockerfile"
		baseURL        = "http://localhost:" + port
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{Timeout: 1 * time.Second}}
	)
	adapters.StartDockerServer(t, port, DockerFilePath)
	specs.GreeterSpecification(t, driver)
}
