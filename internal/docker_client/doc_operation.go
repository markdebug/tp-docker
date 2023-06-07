package container

import (
	"github.com/docker/docker/client"
)

var dockerClient *client.Client

func GetDockerClient() *client.Client {
	if dockerClient == nil {
		cli, err := client.NewClientWithOpts(client.WithVersion("1.41"))
		if err != nil {
			panic(err)
		}
		dockerClient = cli
	}
	return dockerClient
}
