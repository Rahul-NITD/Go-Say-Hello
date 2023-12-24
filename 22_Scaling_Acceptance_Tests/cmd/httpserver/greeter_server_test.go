package main_test

import (
	scalingacceptancetests "GoSayHello/22_Scaling_Acceptance_Tests"
	"GoSayHello/22_Scaling_Acceptance_Tests/specs"
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {

	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../.",
			Dockerfile: "./cmd/httpsserver/Dockerfile",
		},
		ExposedPorts: []string{"8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	driver := scalingacceptancetests.Driver{BaseURL: "http://localhost:8080"}
	specs.GreeterSpecification(t, driver)
}
