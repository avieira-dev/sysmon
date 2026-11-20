package procfs

import (
	"bufio"
	"errors"
	"os"
)

func readCPUStats() (string, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}

		return "", errors.New("empty /proc/stat file")
	}

	return scanner.Text(), nil
}

func ReadCPUStats() (CPUStats, error) {
	content, err := readCPUStats()
	if err != nil {
		return CPUStats{}, err
	}

	return parseCPUStats(content)
}
