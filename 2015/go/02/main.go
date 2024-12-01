package main

import (
	"fmt"

	"github.com/adamliliemark/advent-of-code/utils"
)

type Box struct {
	Width  int
	Height int
	Length int
}

func calculateArea(b Box) int {
	return 2*b.Height*b.Length + 2*b.Width*b.Height + 2*b.Length*b.Width
}

func calculateWrappingPaperNeed(b Box) int {
	return int(calculateArea(b)) + utils.SmallestValue([]int{int(b.Height * b.Length), int(b.Length * b.Width), int(b.Width * b.Height)})
}

func part1(input []string) int {
	paperNeeded := 0
	for _, val := range input {
		var length, width, height int
		fmt.Sscanf(val, "%dx%dx%d", &length, &width, &height)
		box := Box{Length: length, Width: width, Height: height}
		paperNeeded += calculateWrappingPaperNeed(box)
	}
	return paperNeeded
}

func part2(input []string) int {
	ribbonNeeded := 0
	for _, val := range input {
		var length, width, height int
		fmt.Sscanf(val, "%dx%dx%d", &length, &width, &height)
		boxFacePerimeters := []int{
			(length + width) * 2,
			(length + height) * 2,
			(width + height) * 2,
		}
		smallestPerimeter := utils.SmallestValue(boxFacePerimeters)
		neededForBow := length * width * height
		ribbonNeeded += smallestPerimeter + neededForBow
	}

	return ribbonNeeded
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))

}
