package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Bienvenu dans l'outil d'analyse de fichier de logs. Veuillez indiquer le chemin du fichier que vous souhaitez analyser :")
	var filename string
	fmt.Scan(&filename)

	var errcount int
	errcount = 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "error") || strings.Contains(line, "fail") {
			fmt.Println(line)
			errcount += 1
		}
	}
	fmt.Printf("Il y a %d erreurs \n", errcount)

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur de lecture :", err)
	}
}