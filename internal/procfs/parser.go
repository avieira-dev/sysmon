package procfs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseCPUStats(content string) (CPUStats, error) {
	fields := strings.Fields(content)
	if len(fields) < 11 {
		return CPUStats{}, errors.New("invalid /proc/stat format: expected at least 11 fields")
	}

	if fields[0] != "cpu" {
		return CPUStats{}, fmt.Errorf("invalid format in /proc/stat: expected an aggregated CPU line (cpu), but found %s", fields[0])
	}

	fieldNames := []string{
		"User",
		"Nice",
		"System",
		"Idle",
		"IOWait",
		"IRQ",
		"SoftIRQ",
		"Steal",
		"Guest",
		"GuestNice",
	}

	var convertedValues [10]uint64
	for i := 1; i <= 10; i++ {
		value, err := strconv.ParseUint(fields[i], 10, 64)
		if err != nil {
			return CPUStats{}, fmt.Errorf("failed to parse CPU field %s: %w", fieldNames[i-1], err)
		}

		convertedValues[i-1] = value
	}

	return CPUStats{
		User:      convertedValues[0],
		Nice:      convertedValues[1],
		System:    convertedValues[2],
		Idle:      convertedValues[3],
		IOWait:    convertedValues[4],
		IRQ:       convertedValues[5],
		SoftIRQ:   convertedValues[6],
		Steal:     convertedValues[7],
		Guest:     convertedValues[8],
		GuestNice: convertedValues[9],
	}, nil
}
