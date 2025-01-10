package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func main() {
	containers, err := ListDockerContainers()
	if err != nil {
		log.Fatal(err)
	}
	for _, container := range containers {
		fmt.Println(container)
	}
}

func ListDockerContainers() ([][3]string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Unabel to create docker client, please make sure that docker is installed\n%s", err.Error())
		os.Exit(1)
	}

	images, err := cli.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		log.Fatal("Unabel to get images, please make sure that docker daemon is up and running")
		os.Exit(1)
	}

	var containers [][3]string
	for _, image := range images {
		repository := "<none>"
		tag := "<none>"
		if len(image.RepoTags) > 0 {
			splitted := strings.Split(image.RepoTags[0], ":")
			repository = splitted[0]
			tag = splitted[1]
		} else if len(image.RepoDigests) > 0 {
			repository = strings.Split(image.RepoDigests[0], "@")[0]
		}
		containers := append(containers, [3]string{image.ID[7:19], repository, tag})
		fmt.Println(containers)
	}
	return containers, nil

}
