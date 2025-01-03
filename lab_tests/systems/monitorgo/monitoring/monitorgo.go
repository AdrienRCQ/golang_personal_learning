package monitoring

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

// GetRAM retrieves information about the system's RAM.
func GetRAM() (string, error) {
	// Récupérer les informations de mémoire
	v, err := mem.VirtualMemory()
	if err != nil {
		return "", fmt.Errorf("erreur lors de la récupération de la mémoire : %v", err)
	}

	info := fmt.Sprintf(
		"Total Memory : %v MB\nUsed Memory : %v\nFree Memory : %vMB",
		v.Total/1024/1024, //format pour avoir une donnée compréhensible
		v.Used/1024/1024,
		v.Free/1024/1024,
	)
	return info, nil //info => string et nil car pas d'erreur
}

// GetCPU retrieves information about CPU usage.
func GetCPU() (string, error) {
	cpuPercentages, err := cpu.Percent(0, true)
	if err != nil {
		return "", fmt.Errorf("error retrieving CPU usage: %v", err)
	}

	info := "CPU Usage:\n"
	for i, perc := range cpuPercentages {
		info += fmt.Sprintf("Core %d: %.2f%%\n", i, perc)
	}
	return info, nil
}

// GetDisk retrieves information about disk usage.
func GetDisk() (string, error) {
	// Retrieve disk usage information
	usage, err := disk.Usage("/")
	if err != nil {
		return "", fmt.Errorf("error retrieving disk usage: %v", err)
	}

	info := fmt.Sprintf(
		"Total Disk Space: %v GB\nUsed Disk Space: %v GB\nFree Disk Space: %v GB",
		usage.Total/1024/1024/1024,
		usage.Used/1024/1024/1024,
		usage.Free/1024/1024/1024,
	)
	return info, nil
}
