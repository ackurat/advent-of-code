package main

import (
	"encoding/json"
	"fmt"

	"github.com/ackurat/advent-of-code/utils/go/utils"
)

type packetPair struct {
	left, right []any
}

func splitInputToPairs(packets [][]any) []packetPair {
	var pairs []packetPair
	var left, right []any
	for i, packet := range packets {
		if i%2 == 0 {
			left = packet
		} else {
			right = packet
			pairs = append(pairs, packetPair{left, right})
		}
	}
	return pairs
}

func parseInputToList(input []string) [][]any {
	var packets [][]any
	for _, line := range input {
		var packet []any
		if len(line) == 0 {
			continue
		}
		json.Unmarshal([]byte(line), &packet)
		packets = append(packets, packet)
	}
	return packets
}

func compare(left, right []any) int {
	leftLen, rightLen := len(left), len(right)
	if leftLen == 0 && rightLen == 0 {
		return 0
	}

	for i := 0; ; i++ {
		if i >= leftLen || i >= rightLen {
			return leftLen - rightLen
		}

		leftInt, leftIsInt := left[i].(float64)
		rightInt, rightIsInt := right[i].(float64)

		if leftIsInt && rightIsInt {
			if leftInt != rightInt {
				return int(leftInt - rightInt)
			}
			continue
		}

		var (
			l2, r2 []any
		)
		switch left[i].(type) {
		case float64:
			l2 = []any{left[i]}
		default:
			l2 = left[i].([]any)
		}

		switch right[i].(type) {
		case float64:
			r2 = []any{right[i]}
		default:
			r2 = right[i].([]any)
		}
		cmp := compare(l2, r2)
		if cmp != 0 {
			return cmp
		}
	}
}

func part1(input []string) int {
	packets := parseInputToList(input)
	pairs := splitInputToPairs(packets)

	sum := 0
	for i, pair := range pairs {
		if compare(pair.left, pair.right) < 0 {
			sum += i + 1
		}
	}

	return sum
}

func part2(input []string) int {
	packets := parseInputToList(input)
	i2idx, i6idx := 1, 2
	for _, packet := range packets {
		if compare(packet, []any{float64(2)}) < 0 {
			i2idx += 1
			i6idx += 1
		} else if compare(packet, []any{float64(6)}) < 0 {
			i6idx += 1
		}
	}
	return i2idx * i6idx
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
