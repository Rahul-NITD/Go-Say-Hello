package main_test

import (
	"fmt"
	"testing"

	"github.com/Rahul-NITD/scalingacceptancetests/adapters"
	"github.com/Rahul-NITD/scalingacceptancetests/adapters/grpcserver"
	"github.com/Rahul-NITD/scalingacceptancetests/specs"
)

func TestGreeterServer(t *testing.T) {

	var (
		port           = "50051"
		dockerFilePath = "./cmd/grpcserver/DockerFile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specs.GreeterSpecification(t, &driver)
}
