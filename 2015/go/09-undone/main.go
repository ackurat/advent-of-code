package main

import (
	"fmt"

	"github.com/ackurat/advent-of-code/graphs"
	"github.com/ackurat/advent-of-code/utils"
)

func part1(input []string) int {
	var edges []graphs.Edge

	for _, trip := range input {
		var start, end string
		var distance int
		fmt.Sscanf(trip, "%s to %s = %d", &start, &end, &distance)
		startNode := graphs.Node{Name: start}
		endNode := graphs.Node{Name: end}
		edges = append(edges, graphs.Edge{Weight: distance, Start: startNode, End: endNode})
	}
	fmt.Println(edges)

	return 0
}

func traverseEdges(edges []graphs.Edge) {

	for idx, edge := range edges {

	}
}

func dfs()

func main() {
	input := utils.ReadFileLineByLine("input2.txt")

	fmt.Println(part1(input))
}
