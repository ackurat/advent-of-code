package utils

import (
	"fmt"
	"os"
)

func ReadFileToString(path string) string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file %s", path)
	}
	return string(fileContent)
}
