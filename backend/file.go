package main

import (
	"fmt"
	"io"
	"os"
)

func readFromFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func writeToFile(fileName, data string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(data))
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return os.WriteFile(fileName, []byte(data), 0644)
}
