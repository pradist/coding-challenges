package ccwc

import (
	"bufio"
	"os"
)

func CountLines(filePath string) (int, error) {
	file, err := openFile(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return lineCount, nil
}

func openFile(filePath string) (*os.File, error) {
	return os.Open(filePath)
}
