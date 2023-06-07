package main

import (
	"fmt"

	tp_container "tp-doc/internal/container"
	tp_image "tp-doc/internal/images"
	tp_network "tp-doc/internal/network"
)

func main() {
	containers := tp_container.GetRunningContainer()
	for _, container := range containers {
		fmt.Printf("容器ID: %s,容器名称: %s,容器状态: %s,容器运行时间: %s\n", container.Id, container.Name, container.State, container.Status)
	}

	images := tp_image.GetAllImages()

	for _, image := range images {
		fmt.Printf("镜像ID: %s,容器父类: %s,容器标签: %s,镜像大小: %d\n", image.Id, image.ParentId, image.Tag, image.Size)
	}

	network := tp_network.GetAllNetwork()

	for _, net := range network {
		fmt.Printf("网络ID: %s,网络名称: %s,网络标签: %s,网络选项: %d\n", net.ID, net.Name, net.Driver, net.Options)
	}
}
