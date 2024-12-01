package main

import (
	"fmt"
	"math"

	"github.com/ackurat/advent-of-code/utils"
)

type coords struct {
	x int
	y int
}

func (c coords) touching(b coords) bool {
	return math.Abs(float64(c.x-b.x)) <= 1 && math.Abs(float64(c.y-b.y)) <= 1
}

func dYX(a, b int) int {
	if a == b {
		return 0
	}
	return int((a - b) / int(math.Abs(float64(a)-float64(b))))
}

var movements = map[string]coords{
	"U": {x: 0, y: 1},
	"D": {x: 0, y: -1},
	"R": {x: 1, y: 0},
	"L": {x: -1, y: 0},
}

func part1(input []string) int {
	tailPos, headPos := coords{x: 0, y: 0}, coords{x: 0, y: 0}
	tailVisits := make(map[coords]bool)
	tailVisits[tailPos] = true
	for _, line := range input {
		var direction string
		var moves int
		fmt.Sscanf(line, "%s %d", &direction, &moves)
		for i := 0; i < moves; i++ {
			headPos = coords{x: headPos.x + (movements[direction].x), y: headPos.y + (movements[direction].y)}
			if !headPos.touching(tailPos) {
				tailPos = coords{x: tailPos.x + dYX(headPos.x, tailPos.x), y: tailPos.y + dYX(headPos.y, tailPos.y)}
			}
			tailVisits[tailPos] = true
		}
	}
	return len(tailVisits)
}

func part2(input []string) int {
	var knots [10]coords
	for i := 0; i < 10; i++ {
		knots[i] = coords{x: 0, y: 0}
	}
	tailVisits := make(map[coords]bool)
	tailVisits[knots[0]] = true
	for _, line := range input {
		var direction string
		var moves int
		fmt.Sscanf(line, "%s %d", &direction, &moves)
		for i := 0; i < moves; i++ {
			knots[0] = coords{x: knots[0].x + (movements[direction].x), y: knots[0].y + (movements[direction].y)}
			for j := 1; j < len(knots); j++ {
				if !knots[j-1].touching(knots[j]) {
					knots[j] = coords{x: knots[j].x + dYX(knots[j-1].x, knots[j].x), y: knots[j].y + dYX(knots[j-1].y, knots[j].y)}
				}
			}

			tailVisits[knots[len(knots)-1]] = true
		}
	}
	return len(tailVisits)
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
