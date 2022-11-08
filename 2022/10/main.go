package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lookAndSay(input string) (r string) {
	digit := string(input[0])
	count := 1
	var strBuilder strings.Builder

	for i := 1; i < len(input); i++ {
		current := input[i]
		if string(current) != string(digit) {
			strBuilder.WriteString(strconv.Itoa(count) + digit)
			count = 1
			digit = string(current)
		} else {
			count += 1
		}
	}
	strBuilder.WriteString(strconv.Itoa(count) + digit)
	r = strBuilder.String()
	return r
}

func part1And2(input string, iterations int) string {
	start := input
	for i := 0; i < iterations; i++ {
		start = lookAndSay(start)
	}
	return start
}

func main() {
	input := "1113222113"

	fmt.Println(len(part1And2(input, 40)))
	fmt.Println(len(part1And2(input, 50)))

}
