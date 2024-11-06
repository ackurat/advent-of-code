package main

import (
	"fmt"
	"strconv"

	"github.com/adamliliemark/advent-of-code/utils"
)

func part2(input []string) (total int) {
	total = 0

	for i := 1; i < len(input)-2; i++ {
		// convert the next two entries to int
		next, err1 := strconv.Atoi(input[i+1])
		nextNext, err2 := strconv.Atoi(input[i+2])
		if err1 != nil || err2 != nil {
			continue
		}
		curr, err2 := strconv.Atoi(input[i])
		currentSum := curr + next + nextNext

		prev, err1 := strconv.Atoi(input[i-1])
		if err1 != nil || err2 != nil {
			continue
		}
		if curr > prev {
			total += 1
		}
	}

	return total
}

func part1(input []string) (total int) {
	total = 0

	for i := 1; i < len(input); i++ {
		prev, err1 := strconv.Atoi(input[i-1])
		curr, err2 := strconv.Atoi(input[i])
		if err1 != nil || err2 != nil {
			continue // Skip this iteration if conversion fails
		}
		if curr > prev {
			total += 1
		}
	}

	return total
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
