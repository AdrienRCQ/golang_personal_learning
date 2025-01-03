package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getos() {
	filePath := "/etc/os-release"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier %s : %v/n", filePath, err)
		return
	}
	defer file.Close()

	osInfo := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			key := parts[0]
			value := strings.Trim(parts[1], "\"")
			osInfo[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier %s : %v\n", filePath, err)
		return
	}
	// Afficher les informations principales
	fmt.Println("Informations sur le système d'exploitation :")
	if name, exists := osInfo["NAME"]; exists {
		fmt.Printf("Nom : %s\n", name)
	}
	if version, exists := osInfo["VERSION"]; exists {
		fmt.Printf("Version : %s\n", version)
	}
	if prettyName, exists := osInfo["PRETTY_NAME"]; exists {
		fmt.Printf("Nom complet : %s\n", prettyName)
	}
}
