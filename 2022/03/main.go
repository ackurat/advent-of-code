package main

import (
	"fmt"

	"github.com/adamliliemark/advent-of-code/utils"
)

type House struct {
	x int
	y int
}

func part1(input string) int {
	movements := map[string][]int{
		"^": {0, 1},
		"v": {0, -1},
		"<": {-1, 0},
		">": {1, 0},
	}

	currentHouse := House{0, 0}
	houseSet := make(map[House]bool)
	houseSet[currentHouse] = true

	for _, movement := range input {
		movstr := string(movement)
		newHouse := House{currentHouse.x + movements[movstr][0], currentHouse.y + movements[movstr][1]}
		houseSet[newHouse] = true
		currentHouse = newHouse
	}

	return len(houseSet)
}

func part2(input string) int {
	movements := map[string][]int{
		"^": {0, 1},
		"v": {0, -1},
		"<": {-1, 0},
		">": {1, 0},
	}

	santaCurrentHouse := House{0, 0}
	roboCurrentHouse := House{0, 0}
	houseSet := make(map[House]bool)
	houseSet[santaCurrentHouse] = true

	for i, movement := range input {
		movstr := string(movement)
		var newHouse House
		if i%2 == 0 {
			newHouse = House{santaCurrentHouse.x + movements[movstr][0], santaCurrentHouse.y + movements[movstr][1]}
			santaCurrentHouse = newHouse
		} else {
			newHouse = House{roboCurrentHouse.x + movements[movstr][0], roboCurrentHouse.y + movements[movstr][1]}
			roboCurrentHouse = newHouse
		}
		houseSet[newHouse] = true
	}

	return len(houseSet)
}

func main() {
	input := utils.ReadFileToString("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))

}
