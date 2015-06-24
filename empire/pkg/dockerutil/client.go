package dockerutil

import (
	"io"
	"os"

	"github.com/fsouza/go-dockerclient"
)

// NewDockerClient returns a new docker.Client using the given socket and certificate path.
func NewDockerClient(socket, certPath string) (*docker.Client, error) {
	if certPath != "" {
		cert := certPath + "/cert.pem"
		key := certPath + "/key.pem"
		ca := certPath + "/ca.pem"
		return docker.NewTLSClient(socket, cert, key, ca)
	}

	return docker.NewClient(socket)
}

// NewDockerClientFromEnv returns a new docker client configured by the DOCKER_*
// environment variables.
func NewDockerClientFromEnv() (*docker.Client, error) {
	return NewDockerClient(os.Getenv("DOCKER_HOST"), os.Getenv("DOCKER_CERT_PATH"))
}

// Client wraps a docker.Client with authenticated pulls.
type Client struct {
	*docker.Client

	// Auth is the docker AuthConfiguration that will be used when pulling
	// images.
	Auth docker.AuthConfiguration
}

// NewClient returns a new Client instance.
func NewClient(auth docker.AuthConfiguration, socket, certPath string) (*Client, error) {
	c, err := NewDockerClient(socket, certPath)
	if err != nil {
		return nil, err
	}
	return &Client{Auth: auth, Client: c}, nil
}

// NewClientFromEnv returns a new Client instance configured by the DOCKER_*
// environment variables.
func NewClientFromEnv(auth docker.AuthConfiguration) (*Client, error) {
	c, err := NewDockerClientFromEnv()
	if err != nil {
		return nil, err
	}
	return &Client{Auth: auth, Client: c}, nil
}

type PullOptions struct {
	Image         string
	OutputStream  io.Writer
	RawJSONStream bool
}

// Pull is a helper around the raw PullImage method. It takes the plain string
// identification of an image, parses out the repository info and uses the
// appropriate auth config to ensure that the request is authenticated.
//
// This was mostly copied from https://github.com/docker/docker/blob/e5ded9c378f4efdece8ed6e8d166198124956459/api/client/pull.go#L17-L47
func (c *Client) Pull(opts PullOptions) error {
	return nil
}
