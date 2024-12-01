package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ackurat/advent-of-code/utils/go/utils"
)

type Light struct {
	x     int
	y     int
	state bool
}

type DimmableLight struct {
	x     int
	y     int
	state int
}

type XYPair struct {
	x int
	y int
}

func toggle(light *Light) {
	light.state = !light.state
}

func turnOn(light *Light) {
	light.state = true
}

func turnOff(light *Light) {
	light.state = false
}

func incrOne(light *DimmableLight) {
	light.state += 1
}

func decrOne(light *DimmableLight) {
	if light.state > 0 {
		light.state -= 1
	}
}

func incrTwo(light *DimmableLight) {
	light.state += 2
}

func initLights() [1000][1000]Light {
	lights := [1000][1000]Light{}

	for i := 0; i < len(lights)-1; i++ {
		for j := 0; j < len(lights)-1; j++ {
			lights[i][j] = Light{i, j, false}
		}
	}

	return lights
}

func initDimmableLights() [1000][1000]DimmableLight {
	lights := [1000][1000]DimmableLight{}

	for i := 0; i < len(lights)-1; i++ {
		for j := 0; j < len(lights)-1; j++ {
			lights[i][j] = DimmableLight{i, j, 0}
		}
	}

	return lights
}

func parseInputLinePart1(input string) (XYPair, XYPair, func(light *Light)) {
	splitInput := strings.Split(input, " ")
	var fn func(light *Light)

	if splitInput[1] == "on" {
		fn = turnOn
	} else if splitInput[1] == "off" {
		fn = turnOff
	} else {
		fn = toggle
	}

	start := splitInput[len(splitInput)-3]
	end := splitInput[len(splitInput)-1]

	startCoordsStr := strings.Split(start, ",")
	endCoordsStr := strings.Split(end, ",")

	var startCoordsInt [2]int64
	var endCoordsInt [2]int64
	for i := 0; i < 2; i++ {
		startCoordsInt[i], _ = strconv.ParseInt(startCoordsStr[i], 10, 64)
		endCoordsInt[i], _ = strconv.ParseInt(endCoordsStr[i], 10, 64)
	}

	startPair := XYPair{x: int(startCoordsInt[0]), y: int(startCoordsInt[1])}
	endPair := XYPair{x: int(endCoordsInt[0]), y: int(endCoordsInt[1])}
	return startPair, endPair, fn
}

func parseInputLinePart2(input string) (XYPair, XYPair, func(light *DimmableLight)) {
	splitInput := strings.Split(input, " ")
	var fn func(light *DimmableLight)

	if splitInput[1] == "on" {
		fn = incrOne
	} else if splitInput[1] == "off" {
		fn = decrOne
	} else {
		fn = incrTwo
	}

	start := splitInput[len(splitInput)-3]
	end := splitInput[len(splitInput)-1]

	startCoordsStr := strings.Split(start, ",")
	endCoordsStr := strings.Split(end, ",")

	var startCoordsInt [2]int64
	var endCoordsInt [2]int64
	for i := 0; i < 2; i++ {
		startCoordsInt[i], _ = strconv.ParseInt(startCoordsStr[i], 10, 64)
		endCoordsInt[i], _ = strconv.ParseInt(endCoordsStr[i], 10, 64)
	}

	startPair := XYPair{x: int(startCoordsInt[0]), y: int(startCoordsInt[1])}
	endPair := XYPair{x: int(endCoordsInt[0]), y: int(endCoordsInt[1])}
	return startPair, endPair, fn
}

func part1(input []string) int {
	lights := initLights()

	for _, instruction := range input {
		startPair, endPair, fn := parseInputLinePart1(instruction)

		for i := startPair.x; i <= endPair.x; i++ {
			for j := startPair.y; j <= endPair.y; j++ {
				fn(&lights[i][j])
			}
		}
	}

	turnedOnLights := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if lights[i][j].state {
				turnedOnLights += 1
			}
		}
	}

	return turnedOnLights
}

func part2(input []string) int {
	lights := initDimmableLights()

	for _, instruction := range input {
		startPair, endPair, fn := parseInputLinePart2(instruction)

		for i := startPair.x; i <= endPair.x; i++ {
			for j := startPair.y; j <= endPair.y; j++ {
				fn(&lights[i][j])
			}
		}
	}

	totalBrightness := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			totalBrightness += lights[i][j].state
		}
	}

	return totalBrightness
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))

}
