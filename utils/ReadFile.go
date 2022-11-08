package utils

import (
	"bufio"
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

func ReadFileLineByLine(path string) []string {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var strings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return strings

}
