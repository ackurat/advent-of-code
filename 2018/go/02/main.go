package main

import (
	"fmt"

	"github.com/ackurat/advent-of-code/utils"
)

func part1(input []string) int {
	var doublets int
	var triplets int
	for _, line := range input {
		var doubletFlag, tripletFlag = false, false
		letters := make(map[rune]int)
		for _, letter := range line {
			letters[letter] += 1
		}
		for _, occurrences := range letters {
			switch occurrences {
			case 2:
				if !doubletFlag {
					doublets += 1
					doubletFlag = true
				}
			case 3:
				if !tripletFlag {
					triplets += 1
					tripletFlag = true
				}
			}
		}

	}
	return doublets * triplets
}

func part2(input []string) string {
	for _, outerLine := range input {
		for _, innerLine := range input {
			diff := 0
			for i := 0; i < len(outerLine); i++ {
				if outerLine[i] == innerLine[i] {
					continue
				} else {
					diff += 1
				}
			}
			if diff == 1 {
				return utils.OrderedStringIntersection(outerLine, innerLine)
			}
		}
	}
	return ""
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
