package images

import (
	"context"
	docker_client "tp-doc/internal/docker_client"

	"github.com/docker/docker/api/types"
)

type Image struct {
	Id       string
	ParentId string
	Size     int64
	Tag      map[string]string
	Created  int64
}

func GetAllImages() []Image {
	cli := docker_client.GetDockerClient()
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	result := []Image{}
	for _, img := range images {
		tmp := Image{Id: img.ID, ParentId: img.ParentID, Size: img.Size}
		result = append(result, tmp)
	}
	return result
}

func ImageRemove(id string) []types.ImageDeleteResponseItem {
	data, err := docker_client.GetDockerClient().ImageRemove(context.Background(), id, types.ImageRemoveOptions{})
	if err != nil {
		panic(err)
	}
	return data
}
