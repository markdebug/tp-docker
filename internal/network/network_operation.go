package network

import (
	"context"
	docker_client "tp-doc/internal/docker_client"

	"github.com/docker/docker/api/types"
)

type DockerNetwork struct {
	Name       string
	ID         string
	Scope      string
	Driver     string
	EnableIPv6 bool
	Internal   bool
	Attachable bool
	Ingress    bool
	ConfigOnly bool
	Options    map[string]string
	Labels     map[string]string
}

func GetAllNetwork() []DockerNetwork {
	cli := docker_client.GetDockerClient()
	netwrok, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic(err)
	}
	result := []DockerNetwork{}
	for _, net := range netwrok {
		tmp := DockerNetwork{Name: net.Name, ID: net.ID, Driver: net.Driver, Scope: net.Scope, EnableIPv6: net.EnableIPv6, Internal: net.Internal, Ingress: net.Ingress, Options: net.Options, Labels: net.Labels}
		result = append(result, tmp)
	}
	return result
}
