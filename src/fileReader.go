package src

import (
	"bufio"
	"os"
	"path/filepath"
)

func ReadFileAsLines(filename string) ([]string, error) {
	path := filepath.Join("src", filename)
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
