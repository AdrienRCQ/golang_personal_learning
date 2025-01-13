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
	menu()
	fmt.Println("Fin du test !!! ")
	
}

func menu () {
	var number int
	fmt.Scanln(&number) // Récupération de l'input

	fmt.Println(number)
	switch number {
		case 1 : 
			dockerimages, err := ListDockerImages()
			if err != nil {
				log.Fatal(err)
			}
			for _, container := range dockerimages {
				fmt.Println(container)
			}
		case 2 :
			fmt.Println("Nothing here")

	}
}

func ListDockerImages() ([][3]string, error) { // Cette fonction va identifier les container présents et les stocker dans un tableau
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

	var dockerimages [][3]string // création d'un tableau de 3 éléments pour stocker nos dockerimages (enfin les ID, nom et tags)
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
		dockerimages := append(dockerimages, [3]string{image.ID[7:19], repository, tag}) //ajout des ontainers trouvé dans notre tableau
		fmt.Println(dockerimages)
	}
	return dockerimages, nil

}
