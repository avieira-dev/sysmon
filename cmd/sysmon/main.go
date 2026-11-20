package main

import (
	"fmt"

	"github.com/avieira-dev/sysmon/internal/procfs"
)

func main() {
	cpuStats, err := procfs.ReadCPUStats()
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("User:", cpuStats.User)
	fmt.Println("Nice:", cpuStats.Nice)
	fmt.Println("System:", cpuStats.System)
	fmt.Println("Idle:", cpuStats.Idle)
	fmt.Println("IOWait:", cpuStats.IOWait)
	fmt.Println("IRQ:", cpuStats.IRQ)
	fmt.Println("SoftIRQ:", cpuStats.SoftIRQ)
	fmt.Println("Steal:", cpuStats.Steal)
	fmt.Println("Guest:", cpuStats.Guest)
	fmt.Println("GuestNice:", cpuStats.GuestNice)
}
