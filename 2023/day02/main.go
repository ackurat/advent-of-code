package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/adamliliemark/advent-of-code/utils"
)

type Round struct {
	Blue  int `json:"blue"`
	Red   int `json:"red"`
	Green int `json:"green"`
}

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func parse(input []string) map[string][]Round {
	gameList := make(map[string][]Round)

	for _, line := range input {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		gameId := strings.Split(strings.TrimSpace(parts[0]), " ")[1]
		gameData := strings.TrimSpace(parts[1])

		rounds := strings.Split(gameData, ";")
		roundResults := make([]Round, len(rounds))

		for i, game := range rounds {
			game = strings.TrimSpace(game)
			colors := strings.Split(game, ",")

			var blue, red, green int
			for _, color := range colors {
				color = strings.TrimSpace(color)
				count, err := strconv.Atoi(strings.Split(color, " ")[0])
				if err != nil {
					continue
				}

				switch {
				case strings.Contains(color, "blue"):
					blue += count
				case strings.Contains(color, "red"):
					red += count
				case strings.Contains(color, "green"):
					green += count
				}
			}

			roundResults[i] = Round{blue, red, green}
		}

		gameList[gameId] = roundResults
	}

	return gameList
}

func part2(input []string) int {
	return 0
}

func checkOneRound(round Round) bool {
	return round.Blue < MAX_BLUE && round.Red < MAX_RED && round.Green < MAX_GREEN
}

func checkOneGame(rounds []Round) bool {
	for _, round := range rounds {
		if !checkOneRound(round) {
			return false
		}
	}
	return true
}

func part1(input []string) int {
	allGames := parse(input)
	total := 0
	for id, game := range allGames {
		if checkOneGame(game) {
			toAdd, _ := strconv.Atoi(id)
			total += toAdd
		}
	}

	prettyPrint(allGames)
	return total
}

func prettyPrint(games map[string][]Round) {
	jsonData, err := json.MarshalIndent(games, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func main() {
	input := utils.ReadFileLineByLine("example.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
