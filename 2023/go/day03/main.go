package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/adamliliemark/advent-of-code/utils"
)

type Part struct {
	row        int
	start, end int
	number     int
}

type Symbol struct {
	row, col  int
	sym       byte
	adjacents *[]Part
}

func isSymbol(b byte) bool {
	return b != 46 && (b < 48 || b > 57)
}

func isStar(b byte) bool {
	return b == 42
}

func checkAdjacent(part Part, symbols []Symbol) bool {
	for _, symbol := range symbols {
		if (symbol.row == part.row || symbol.row+1 == part.row || symbol.row-1 == part.row) && (symbol.col+1 >= part.start && symbol.col-1 <= part.end) {
			if isStar(symbol.sym) {

				*symbol.adjacents = append(*symbol.adjacents, part)
			}
			return true
		}
	}
	return false
}

func prepare(input []string, symbols *[]Symbol, parts *[]Part) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if unicode.IsDigit(rune(input[i][j])) {
				var partsNumber []rune
				currentPart := Part{row: i, start: j}
				for {
					if !unicode.IsDigit(rune(input[i][j])) {
						num, _ := strconv.Atoi(string(partsNumber))
						currentPart.number = num
						j--
						break
					}
					currentPart.end = j
					partsNumber = append(partsNumber, rune(input[i][j]))
					j++
					if j == len(input) {
						num, _ := strconv.Atoi(string(partsNumber))
						currentPart.number = num
						break
					}
				}
				*parts = append(*parts, currentPart)
			} else if isSymbol(input[i][j]) {
				*symbols = append(*symbols, Symbol{row: i, col: j, sym: input[i][j], adjacents: &[]Part{}})
			}
		}
	}
}

func part1(symbols *[]Symbol, parts *[]Part) (total int) {
	for _, part := range *parts {
		if checkAdjacent(part, *symbols) {
			total += part.number
		}
	}
	return
}

func part2(symbols *[]Symbol) (total int) {
	for _, symbol := range *symbols {
		if len(*symbol.adjacents) == 2 {
			total += (*symbol.adjacents)[0].number * (*symbol.adjacents)[1].number
		}
	}
	return
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")
	symbols := []Symbol{}
	parts := []Part{}
	prepare(input, &symbols, &parts)
	fmt.Println(part1(&symbols, &parts))
	fmt.Println(part2(&symbols))
}
