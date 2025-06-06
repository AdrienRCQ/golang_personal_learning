package main

import (
	"encoding/json"
	"fmt"
<<<<<<< HEAD
	"io"
	"net/http"
=======
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
>>>>>>> b88786e (création d'un projet d'interaction go - database)
)

type Container struct {
	ID    string   `json:"Id"`
	Names []string `json:"Names"`
	Image string   `json:"Image"`
	State string   `json:"State"`
}

func main() {
	url := "http://localhost:2375/containers/json"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var containers []Container
	if err := json.Unmarshal(body, &containers); err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("ID: %s, Nom: %s, Image: %s, État: %s\n", container.ID, container.Names, container.Image, container.State)
	}
}
<<<<<<< HEAD
=======

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

func StartContainer(name string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
		return err
	}
	containerConfig := &container.Config{
		Image: name,
	}
	container, err := cli.ContainerCreate(context.Background(), containerConfig, nil, nil, "", nil, nil, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := cli.ContainerStart(context.Background(), container.ID, types.St); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Container started successfully!")
	return nil
}
>>>>>>> b88786e (création d'un projet d'interaction go - database)
