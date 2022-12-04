package main

import (
	"fmt"

	"github.com/adamliliemark/advent-of-code/utils"
)

func part1(input []string) int {
	var sum int
	for _, line := range input {
		var a1, a2, b1, b2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a1, &a2, &b1, &b2)
		if (a1 >= b1 && a2 <= b2) || (a1 <= b1 && a2 >= b2) {
			sum += 1
		}
	}
	return sum
}

func part2(input []string) int {
	var sum int
	for _, line := range input {
		var a1, a2, b1, b2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a1, &a2, &b1, &b2)
		if (a1 >= b1 && a1 <= b2) || (b1 >= a1 && b1 <= a2) {
			sum += 1
		}
	}
	return sum
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))

}
