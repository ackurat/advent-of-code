package main

import (
	"fmt"
	"strings"

	"github.com/ackurat/advent-of-code/utils"
)

func battle(opponentChoice, ownChoice string) string {
	switch opponentChoice {
	case "A":
		if ownChoice == "Y" {
			return "win"
		} else if ownChoice == "X" {
			return "draw"
		}
	case "B":
		if ownChoice == "Z" {
			return "win"
		} else if ownChoice == "Y" {
			return "draw"
		}
	case "C":
		if ownChoice == "X" {
			return "win"
		} else if ownChoice == "Z" {
			return "draw"
		}
	}
	return "lose"
}

func part1(input []string) int {
	var score int
	for _, line := range input {
		v := strings.Fields(line)
		switch v[1] {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}

		switch battle(v[0], v[1]) {
		case "win":
			score += 6
		case "draw":
			score += 3
		case "lose":
			score += 0
		}
	}
	return score
}

func part2(input []string) int {
	var score int
	for _, line := range input {
		v := strings.Fields(line)
		switch v[1] {
		case "X":
			switch v[0] {
			case "A":
				score += 3
			case "B":
				score += 1
			case "C":
				score += 2
			}
		case "Y":
			score += 3
			switch v[0] {
			case "A":
				score += 1
			case "B":
				score += 2
			case "C":
				score += 3
			}
		case "Z":
			score += 6
			switch v[0] {
			case "A":
				score += 2
			case "B":
				score += 3
			case "C":
				score += 1
			}
		}
	}

	return score
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
