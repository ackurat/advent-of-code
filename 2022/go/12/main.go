package main

import (
	"fmt"

	"github.com/ackurat/advent-of-code/graphs"
	"github.com/ackurat/advent-of-code/utils"
)

// FIFO is a FIFO queue
type FIFO struct {
	queue []interface{}
}

// New creates new FIFO and returns it
func New() *FIFO {
	return &FIFO{
		queue: make([]interface{}, 0),
	}
}

func (f *FIFO) Len() int {
	return len(f.queue)
}

// Push pushed node to the back of the queue
func (f *FIFO) Push(node interface{}) {
	f.queue = append(f.queue, node)
}

// Front takes a value from the front of the queue and returns it
func (f *FIFO) Front() interface{} {
	if len(f.queue) == 0 {
		return nil
	}

	node := f.queue[0]
	f.queue[0] = nil
	f.queue = f.queue[1:]

	return node
}

var rowBoundary, colBoundary int

func isValid(row, col int, visited map[graphs.Point]bool) bool {
	if row < 0 || col < 0 || row > rowBoundary || col > colBoundary {
		return false
	}
	if visited[graphs.Point{X: row, Y: col}] {
		return false
	}
	return true
}

func part1(input []string) int {

	var heightMap [][]rune

	var start, end graphs.Point
	for i, line := range input {
		lineRunes := []rune(line)
		heightMap = append(heightMap, lineRunes)
		for j, letter := range line {
			if letter == 'S' {
				start = graphs.Point{X: i, Y: j}
				heightMap[i][j] = 'a'
			} else if letter == 'E' {
				end = graphs.Point{X: i, Y: j}
				heightMap[i][j] = 'z'
			}
		}
	}
	rowBoundary = len(heightMap) - 1
	colBoundary = len(heightMap[0]) - 1

	fifo := New()
	visited := make(map[graphs.Point]bool)

	dRow := []int{-1, 0, 1, 0}
	dCol := []int{0, 1, 0, -1}

	fifo.Push(graphs.Point{X: 0, Y: 0})

	for fifo.Len() != 0 {
		point := fifo.Front().(graphs.Point)
		x, y := point.X, point.Y
		fmt.Println(x, y)
		for i := 0; i < 4; i++ {
			adjX := x + dRow[i]
			adjY := y + dCol[i]
			if isValid(adjX, adjY, visited) {
				if heightMap[adjX][adjY]-heightMap[x][y] <= 1 {
					newPoint := graphs.Point{X: adjX, Y: adjY}
					fifo.Push(newPoint)
					visited[newPoint] = true
				}
			}
		}
	}

	fmt.Println(start, end)
	for _, row := range heightMap {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}

	return 0
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
}
