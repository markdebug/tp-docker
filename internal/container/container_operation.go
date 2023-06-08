package container

import (
	"context"
	"log"

	docker_client "tp-doc/internal/docker_client"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Container struct {
	Id         string
	Name       string
	ImageId    string
	ImageName  string
	Status     string
	State      string
	CreateTime int64
}

func GetRunningContainer() []Container {
	cli := docker_client.GetDockerClient()
	// 获取本地已启动的Docker容器，如果要查看全部容器，可以配置types.ContainerListOptions{}参数
	containers := queryRunningContainer(cli)
	result := []Container{}
	for _, container := range containers {
		if container.ID == "" {
			log.Println("容器id为空 跳过")
			continue
		}
		tmp := Container{Id: container.ID, Name: container.Names[0], ImageId: container.ImageID, ImageName: container.Image, Status: container.Status, State: container.State, CreateTime: container.Created}
		result = append(result, tmp)
	}

	return result

}

func StopRunningContainer(id string) bool {
	log.Println("Stop container id is: ", id)
	if err := docker_client.GetDockerClient().ContainerStop(context.Background(), id, container.StopOptions{}); err != nil {
		return false
	}
	return true
}

func StartContainer(id string) bool {

	err := docker_client.GetDockerClient().ContainerStart(context.Background(), id, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	return true
}

func RenameContainer(id string, name string) bool {

	err := docker_client.GetDockerClient().ContainerRename(context.Background(), id, name)
	if err != nil {
		panic(err)
	}
	return true
}

func KillContainer(id string, signle string) bool {
	err := docker_client.GetDockerClient().ContainerKill(context.Background(), id, signle)
	if err != nil {
		panic(err)
	}
	return true
}

func InspectContainer(id string) types.ContainerJSON {

	json, err := docker_client.GetDockerClient().ContainerInspect(context.Background(), id)
	if err != nil {
		panic(err)
	}

	return json
}

func RestartContainer(id string) bool {

	err := docker_client.GetDockerClient().ContainerRestart(context.Background(), id, container.StopOptions{})
	if err != nil {
		panic(err)
	}
	return true
}

func queryRunningContainer(client *client.Client) []types.Container {
	// 获取本地已启动的Docker容器，如果要查看全部容器，可以配置types.ContainerListOptions{}参数
	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	return containers
}
