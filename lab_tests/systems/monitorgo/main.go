package main

import (
	"fmt"
	"log"
	"monitorgo/monitoring"
)

func main() {
	ram, err := monitoring.GetRAM()
	if err != nil {
		log.Fatalf("Failed to get RAM info: %v", err)
	}

	cpu, err := monitoring.GetCPU()
	if err != nil {
		log.Fatalf("Failed to get CPU info: %v", err)
	}

	disk, err := monitoring.GetDisk()
	if err != nil {
		log.Fatalf("Failed to get Disk info: %v", err)
	}

	fmt.Println(ram)
	fmt.Println("----------")
	fmt.Println(cpu)
	fmt.Println("----------")
	fmt.Println(disk)
}
