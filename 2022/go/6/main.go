package main

import (
	"fmt"

	"github.com/ackurat/advent-of-code/utils/go/utils"
)

func part1(input string) (marker int) {
	for i := 3; i < len(input); i++ {
		set := utils.NewSet[byte]()
		for j := 0; j < 4; j++ {
			set.Add(input[i-j])
		}
		if len(set) == 4 {
			marker = i + 1
			break
		}
	}
	return
}

func part2(input string) (marker int) {
	for i := 14; i < len(input); i++ {
		set := utils.NewSet[byte]()
		for j := 0; j < 15; j++ {
			set.Add(input[i-j])
		}
		if len(set) == 14 {
			marker = i + 1
			break
		}
	}
	return
}

func main() {
	input := utils.ReadFileToString("input2.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
