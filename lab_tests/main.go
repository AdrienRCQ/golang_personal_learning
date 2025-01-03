package main

import (
	"fmt"
	"runtime"
)

type Person struct {
	name  string
	email string
	age   uint8
}

func main() {

	os := runtime.GOOS
	fmt.Println("Operating system : ", os)
	adduser("Adrien", "adrien.ricque@cpe.fr", 22)
}

func adduser(name string, email string, age uint8) Person {
	newuser := Person{name, email, age}
	fmt.Println("Bienvenue", newuser.name, "! ")
	return newuser
}
