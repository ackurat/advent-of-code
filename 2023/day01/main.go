package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/adamliliemark/advent-of-code/utils"
)

func part2(input []string) int {
	numbers := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3hree",
		"four":  "f4ur",
		"five":  "f5ve",
		"six":   "s6x",
		"seven": "s7ven",
		"eight": "e8ght",
		"nine":  "n9ne",
	}
	modifiedInput := make([]string, len(input))
	for idx, line := range input {
		for original, modified := range numbers {
			line = strings.ReplaceAll(line, original, modified)
			modifiedInput[idx] = strings.ReplaceAll(line, original, modified)
		}
	}

	return part1(modifiedInput)
}

func part1(input []string) int {
	digits := []int{}
	r, _ := regexp.Compile("[0-9]")
	for _, line := range input {
		matches := r.FindAllString(line, -1)
		if len(matches) > 0 {
			firstDigit := matches[0]
			lastDigit := matches[len(matches)-1]
			combined := strings.Join([]string{firstDigit, lastDigit}, "")
			digit, _ := strconv.Atoi(combined)
			digits = append(digits, digit)
		}

	}

	return utils.SumOfArray(digits)
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
