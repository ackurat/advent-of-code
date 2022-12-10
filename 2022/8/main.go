package main

import (
	"fmt"
	"strconv"

	"github.com/adamliliemark/advent-of-code/utils"
)

func part1(input []string) int {
	trees := make([][]int64, 5)
	for i := range trees {
		trees[i] = make([]int64, 5)
	}

	for i, line := range input {
		for j, c := range line {
			trees[i][j], _ = strconv.ParseInt(string(c), 10, 32)
		}
	}

	for _, r := range trees {
		for _, c := range r {
			fmt.Print(c)
		}
		fmt.Println()
	}
	return 0
}

func main() {
	input := utils.ReadFileLineByLine("input2.txt")

	fmt.Println(part1(input))
}
