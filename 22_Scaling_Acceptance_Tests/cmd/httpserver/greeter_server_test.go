package main_test

import (
	"context"
	"testing"

	"github.com/Rahul-NITD/scalingacceptancetests"
	"github.com/Rahul-NITD/scalingacceptancetests/specs"
	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../.", // Verify that this path is correct and includes the Dockerfile
			Dockerfile: "./cmd/httpserver/Dockerfile",
		},
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)

	// Clean up the container after the test
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	// Get the container's host port
	port, err := container.MappedPort(ctx, "8080")
	assert.NoError(t, err)

	// Use the host port to set the BaseURL
	driver := scalingacceptancetests.Driver{BaseURL: "http://localhost:" + port.Port()}
	specs.GreeterSpecification(t, driver)
}
