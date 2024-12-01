package main

import (
	"fmt"
	"strconv"

	"github.com/ackurat/advent-of-code/utils/go/utils"
)

func part1(input []string) int {
	var total int
	for _, line := range input {
		i, _ := strconv.ParseInt(line, 10, 32)
		total += int(i)
	}
	return total
}

func part2(input []string) int {
	var total int
	freqs := utils.NewSet(0)
	for {
		for _, line := range input {
			i, _ := strconv.ParseInt(line, 10, 32)
			total += int(i)
			if freqs.Contains(total) {
				return total
			}
			freqs.Add(total)
		}
	}
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
