package main

import (
	"fmt"

	"github.com/ackurat/advent-of-code/utils"
)

func part1and2(input []string) int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}
	var overlapped int
	for _, line := range input {
		var id, left, top, width, height int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &left, &top, &width, &height)
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				grid[i+top][j+left] += 1
			}
		}
	}

	for _, row := range grid {
		for _, col := range row {
			if col > 1 {
				overlapped += 1
			}
		}
	}

	for _, line := range input {
		var id, left, top, width, height int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &left, &top, &width, &height)
		nonOverlapping := true
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if grid[i+top][j+left] > 1 {
					nonOverlapping = false
				}
			}
		}
		if nonOverlapping {
			fmt.Println("nonoverlapping:", id)
		}
	}

	return overlapped
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1and2(input))

}
