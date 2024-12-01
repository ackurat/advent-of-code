package main

import (
	"fmt"

	"github.com/ackurat/advent-of-code/utils"
)

func calcCharSum(r rune) int {
	if r > 90 {
		return int(r - 96)
	} else if r > 64 && r < 91 {
		return int(r - 38)
	}
	return 0

}

func part1(input []string) int {
	var sum int
	for _, line := range input {
		firstHalf := line[0:(len(line) / 2)]
		secondHalf := line[(len(line) / 2):]
		inters := utils.UnOrderedStringIntersection(firstHalf, secondHalf)
		for _, r := range inters {
			sum += calcCharSum(r)
		}
	}
	return sum
}

func part2(input []string) int {
	var sum int
	chunks := utils.SplitToChunks(input, 3)
	for _, chunk := range chunks {
		inters := utils.UnOrderedStringIntersection(chunk[0], chunk[1], chunk[2])
		for _, r := range inters {
			sum += calcCharSum(r)
		}
	}
	return sum
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))

}
