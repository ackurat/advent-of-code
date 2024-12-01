package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/ackurat/advent-of-code/utils/go/utils"
)

func calculateElvesTotals(input []string) []int {
	elvesTotals := []int{}
	var elfTotal int

	for _, line := range input {
		switch line {
		case "":
			elvesTotals = append(elvesTotals, elfTotal)
			elfTotal = 0
		default:
			val, _ := strconv.ParseInt(line, 10, 64)
			elfTotal += int(val)
		}
	}

	return elvesTotals
}

func part2(elvesTotals []int) int {
	sort.Slice(elvesTotals, func(i, j int) bool { return elvesTotals[i] > elvesTotals[j] })

	topThree := elvesTotals[0] + elvesTotals[1] + elvesTotals[2]

	return topThree
}

func part1(elvesTotals []int) int {
	return utils.BiggestValue(elvesTotals)
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")
	elvesTotals := calculateElvesTotals(input)

	fmt.Println(part1(elvesTotals))
	fmt.Println(part2(elvesTotals))
}
