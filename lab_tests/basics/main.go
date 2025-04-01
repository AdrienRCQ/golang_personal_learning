package main

import (
	"affichage"
	"fmt"
)

func main() {
	fmt.Println("Je m'appelle", affichage.Nom)
	fmt.Println(affichage.AfficheSexe())
}
