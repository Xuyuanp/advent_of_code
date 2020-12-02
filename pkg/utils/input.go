package utils

import (
	"bufio"
	"os"
	"strconv"
)

// Input parses input lines from file path
func Input(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// InputInts parses input lines from file path and converts to int
func InputInts(path string) ([]int, error) {
	lines, err := Input(path)
	if err != nil {
		return nil, err
	}
	vals := make([]int, len(lines))

	for i, line := range lines {
		vals[i], err = strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
	}
	return vals, nil
}
