package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/adamliliemark/advent-of-code/utils"
)

func part1matcher(input string) bool {
	guard, _ := regexp.Match(`(ab)|(cd)|(pq)|(xy)`, []byte(input))
	if guard {
		return false
	}

	firstMatch, _ := regexp.Match(`(([^aeiou]*[aeiou]){3,})`, []byte(input))
	if firstMatch {
		for idx, char := range input {
			if idx == 0 {
				continue
			}
			v := strings.Compare(string(char), string(input[idx-1]))
			if v == 0 {
				return true
			}
		}

	}
	return false
}

func part1(input []string) int {
	nice := 0
	for _, str := range input {
		if part1matcher(str) {
			nice += 1
		}
	}

	return nice
}

func pairs(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if strings.Count(input, input[i:i+2]) >= 2 {
			return true
		}
	}
	return false
}

func repeat(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			return true
		}
	}
	return false
}

func part2(input []string) int {
	nice := 0
	for _, str := range input {
		if pairs(str) && repeat(str) {
			nice += 1
		}
	}

	return nice
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	// fmt.Println(part1(input))

	fmt.Println(part2(input))

}
