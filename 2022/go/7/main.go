package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ackurat/advent-of-code/utils/go/graphs"
	"github.com/ackurat/advent-of-code/utils/go/utils"
)

func makeFilesystem(node *graphs.WeightedNode, input []string) {
	if len(input) == 0 {
		return
	}

	cmd := input[0]
	input = input[1:]

	if strings.HasPrefix(cmd, "$") {
		var command, operand string
		fmt.Sscanf(cmd, "$ %s %s", &command, &operand)
		if command == "cd" {
			if operand == ".." {
				node = node.Parent
			} else {
				child := &graphs.WeightedNode{Parent: node}
				node.Children = append(node.Children, child)
				node = child
			}
		} else if command == "ls" {
			for len(input) > 0 && input[0][0] != '$' {
				weight, _ := strconv.Atoi(strings.Split(input[0], " ")[0])
				node.Weight += weight
				input = input[1:]
			}
		}
	}
	makeFilesystem(node, input)
}

func traverseFilesystem(node *graphs.WeightedNode, weights []int) []int {
	if node != nil {
		weight := node.Weight
		for _, child := range node.Children {
			weights = traverseFilesystem(child, weights)
			weight += weights[len(weights)-1]
		}
		return append(weights, weight)
	} else {
		return weights
	}
}

func part1(input []string) (answer int) {
	fs := &graphs.WeightedNode{}
	makeFilesystem(fs, input)
	for _, size := range traverseFilesystem(fs, []int{}) {
		if size < 100000 {
			answer += size
		}
	}

	return
}

func part2(input []string) int {
	fs := &graphs.WeightedNode{}
	makeFilesystem(fs, input)
	directories := traverseFilesystem(fs, []int{})
	sort.Ints(directories)
	sum := directories[len(directories)-1]
	freeSpace := 70000000 - sum
	needToDelete := 30000000 - freeSpace
	for _, directory := range directories {
		if directory > needToDelete {
			return directory
		}
	}

	return 0
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))

}
