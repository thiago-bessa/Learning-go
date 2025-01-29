package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath string, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}

}

func (fileManager FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fileManager.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("Failed to read lines in file.")
	}

	return lines, nil
}

func (fileManager FileManager) WriteResult(data any) error {
	file, err := os.Create(fileManager.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("Failed to convert data to JSON.")
	}

	return nil
}
