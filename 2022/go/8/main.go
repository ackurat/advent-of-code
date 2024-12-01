package main

import (
	"fmt"
	"strconv"

	"github.com/adamliliemark/advent-of-code/graphs"
	"github.com/adamliliemark/advent-of-code/utils"
)

func parseInput(input []string) (trees [][]int) {
	for _, line := range input {
		var row []int
		for _, c := range line {
			val, _ := strconv.Atoi(string(c))
			row = append(row, val)
		}
		trees = append(trees, row)
	}
	return
}

func part1(input []string) int {
	trees := parseInput(input)

	set := utils.NewSet[graphs.Point]()
	for row := 0; row < len(trees); row++ {
		leftHighest := -1
		for col := 0; col < len(trees[row]); col++ {
			if trees[row][col] > leftHighest {
				leftHighest = trees[row][col]
				set.Add(graphs.Point{X: row, Y: col})
			}
		}
		rightHighest := -1
		for col := len(trees[row]) - 1; col >= 0; col-- {
			if trees[row][col] > rightHighest {
				rightHighest = trees[row][col]
				set.Add(graphs.Point{X: row, Y: col})
			}
		}
	}
	for col := 0; col < len(trees[col])-1; col++ {
		topHighest := -1
		for row := 0; row < len(trees)-1; row++ {
			if trees[row][col] > topHighest {
				topHighest = trees[row][col]
				set.Add(graphs.Point{X: row, Y: col})

			}
		}
		bottomHighest := -1
		for row := len(trees) - 1; row >= 0; row-- {
			if trees[row][col] > bottomHighest {
				bottomHighest = trees[row][col]
				set.Add(graphs.Point{X: row, Y: col})
			}
		}
	}

	return set.Len()
}

func part2(input []string) int {
	trees := parseInput(input)

	var scores []int
	for row := 1; row < len(trees)-1; row++ {
		for col := 1; col < len(trees[row])-1; col++ {
			currentTree := trees[row][col]
			var (
				visL, visR, visT, visB int
			)

			for colN := col + 1; colN < len(trees[row]); colN++ {
				if trees[row][colN] < currentTree {
					visR++
				} else {
					visR++
					break
				}
			}

			for colN := col - 1; colN >= 0; colN-- {
				if trees[row][colN] < currentTree {
					visL++
				} else {
					visL++
					break
				}
			}

			for rowN := row + 1; rowN < len(trees[col]); rowN++ {
				if trees[rowN][col] < currentTree {
					visT++
				} else {
					visT++
					break
				}
			}

			for rowN := row - 1; rowN >= 0; rowN-- {
				if trees[rowN][col] < currentTree {
					visB++
				} else {
					visB++
					break
				}
			}

			scenicScore := visL * visR * visT * visB
			scores = append(scores, scenicScore)
		}
	}

	return utils.BiggestValue(scores)
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
