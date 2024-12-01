package main

import (
	"fmt"

	"github.com/ackurat/advent-of-code/utils"
)

func part1(input []string) int {
	stringLiterals := 0
	characters := 0
	for _, str := range input {
		characters += len(str)
		for i := 1; i < len(str)-1; i++ {
			switch str[i] {
			case '\\':
				n := str[i+1]
				if n == '\\' || n == '"' {
					i += 1
				} else if n == 'x' {
					i += 3
				}
			}
			stringLiterals += 1
		}

	}
	return characters - stringLiterals
}

func part2(input []string) int {
	total := 0
	characters := 0
	for _, str := range input {
		characters += len(str)
		total += 2
		for i := 0; i < len(str); i++ {
			switch str[i] {
			case '"', '\\':
				total += 2
			default:
				total += 1
			}

		}
	}

	return total - characters
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))

}
