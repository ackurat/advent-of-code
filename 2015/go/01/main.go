package main

import (
	"fmt"

	"github.com/ackurat/advent-of-code/utils/go"
)

func part1(input string) {
	total := 0

	for _, v := range input {
		if string(v) == "(" {
			total += 1
		} else {
			total -= 1
		}
	}

	fmt.Println(total)
}

func part2(input string) {
	currentFloor := 0

	for i, v := range input {
		if string(v) == "(" {
			currentFloor += 1
		} else {
			currentFloor -= 1
		}
		if currentFloor < 0 {
			fmt.Println(i + 1)
			break
		}
	}
}

func main() {
	input := utils.ReadFileToString("input.txt")
	part1(input)
	part2(input)

}
