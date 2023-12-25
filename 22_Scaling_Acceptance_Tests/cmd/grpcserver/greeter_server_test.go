package main_test

import (
	"fmt"
	"testing"

	"github.com/Rahul-NITD/scalingacceptancetests/adapters"
	"github.com/Rahul-NITD/scalingacceptancetests/adapters/grpcserver"
	"github.com/Rahul-NITD/scalingacceptancetests/specs"
)

func TestGreeterServer(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	var (
		port           = "50051"
		dockerFilePath = "./DockerFile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
		build_bin      = "grpcserver"
	)

	adapters.StartDockerServer(t, port, build_bin, dockerFilePath)
	specs.GreeterSpecification(t, &driver)
	specs.CurseSpecification(t, &driver)
}
