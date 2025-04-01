package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
