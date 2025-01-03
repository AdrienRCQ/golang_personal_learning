package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/shirou/gopsutil/process"
)

func listProcesses() {
	cmd := exec.Command("ps", "-e") // Commande Unix pour lister les processus
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println(string(output))
}

func killprocess() error {
	var name string
	fmt.Scan(&name)

	processes, err := process.Processes()
	if err != nil {
		return err
	}
	for _, p := range processes {
		n, err := p.Name()
		if err != nil {
			return err
		}
		if n == name {
			return p.Kill()
		}
	}
	return fmt.Errorf("process not found")
}

func main() {
	fmt.Println("Gestionnaire de processus")
	fmt.Println("1. Lister les processus")
	fmt.Println("2. Terminer un processus")
	fmt.Println("3. Quitter")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		listProcesses()
	case 2:
		killprocess()
	case 3:
		fmt.Println("Au revoir !")
		os.Exit(0)
	default:
		fmt.Println("Choix invalide.")
	}
}
