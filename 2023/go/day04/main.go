package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/adamliliemark/advent-of-code/utils"
)

type Card struct {
	id      int
	winning []int
	guesses []int
	wins    int
}

func parse(input []string, cards *[]Card) {
	for idx, line := range input {
		card := Card{id: idx - 1}
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winningStr := strings.Split(strings.TrimSpace(numbers[0]), " ")
		guessesStr := strings.Split(strings.TrimSpace(numbers[1]), " ")
		for _, str := range winningStr {
			num, err := strconv.Atoi(str)
			if err == nil {
				card.winning = append(card.winning, num)
			}
		}
		for _, str := range guessesStr {
			num, err := strconv.Atoi(str)
			if err == nil {
				card.guesses = append(card.guesses, num)
			}
		}
		*cards = append(*cards, card)
	}
}

func part1(cards []Card) (total int) {
	for _, card := range cards {
		intersection := utils.GenericIntersection[int](card.guesses, card.winning)
		if len(intersection) > 0 {
			total += int(math.Pow(2, float64(len(intersection)-1)))
		}
	}
	return
}

func part2(cards []Card) int {
	scratchCardOccurrences := utils.DefaultValuedSlice[int](len(cards), 1)

	for i, card := range cards {
		intersection := utils.GenericIntersection[int](card.guesses, card.winning)
		for index := 1; index <= len(intersection); index++ {
			scratchCardOccurrences[i+index] += scratchCardOccurrences[i]
		}
	}

	return utils.SumOfArray(scratchCardOccurrences)
}

func main() {
	input := utils.ReadFileLineByLine("example.txt")
	cards := []Card{}
	parse(input, &cards)

	fmt.Println(part1(cards))
	fmt.Println(part2(cards))
}
