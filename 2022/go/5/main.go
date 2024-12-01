package main

import (
	"fmt"
	"strings"

	"github.com/ackurat/advent-of-code/utils"
)

func initStacks() []*utils.Stack[string] {
	// too lazy to read in from the file due to structure of input...
	stack1 := utils.NewStack("R", "N", "F", "V", "L", "J", "S", "M")
	stack2 := utils.NewStack("P", "N", "D", "Z", "F", "J", "W", "H")
	stack3 := utils.NewStack("W", "R", "C", "D", "G")
	stack4 := utils.NewStack("N", "B", "S")
	stack5 := utils.NewStack("M", "Z", "W", "P", "C", "B", "F", "N")
	stack6 := utils.NewStack("P", "R", "M", "W")
	stack7 := utils.NewStack("R", "T", "N", "G", "L", "S", "W")
	stack8 := utils.NewStack("Q", "T", "H", "F", "N", "B", "V")
	stack9 := utils.NewStack("L", "M", "H", "Z", "N", "F")

	stacks := []*utils.Stack[string]{stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9}
	return stacks
}

func part1(input []string) string {
	stacks := initStacks()
	for _, line := range input {
		var boxes, fromStack, toStack int
		fmt.Sscanf(line, "move %d from %d to %d", &boxes, &fromStack, &toStack)
		for i := 0; i < boxes; i++ {
			stacks[toStack-1].Push(stacks[fromStack-1].Pop())
		}
	}

	var topBoxes strings.Builder
	for _, stack := range stacks {
		topBoxes.WriteString(stack.Peek())
	}

	return topBoxes.String()
}

func part2(input []string) string {
	stacks := initStacks()
	for _, line := range input {
		var boxes, fromStack, toStack int
		fmt.Sscanf(line, "move %d from %d to %d", &boxes, &fromStack, &toStack)
		tempstack := utils.NewStack[string]()
		for i := 0; i < boxes; i++ {
			tempstack.Push(stacks[fromStack-1].Pop())
		}
		for i := 0; i < boxes; i++ {
			stacks[toStack-1].Push(tempstack.Pop())
		}
	}

	var topBoxes strings.Builder
	for _, stack := range stacks {
		topBoxes.WriteString(stack.Peek())
	}

	return topBoxes.String()

}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
