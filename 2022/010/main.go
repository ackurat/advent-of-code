package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adamliliemark/advent-of-code/utils"
)

type processor struct {
	register, cycle, ans int
}

func (p *processor) tick() {
	p.cycle++
	if p.cycle%40 == 20 {
		p.ans += p.cycle * p.register
	}
}

func (p *processor) add(x int) {
	p.register += x
}

func (p *processor) read(input string) {
	if strings.HasPrefix(input, "noop") {
		p.tick()
		return
	} else {
		instr := strings.Fields(input)
		x, _ := strconv.Atoi(instr[1])
		p.tick()
		p.tick()
		p.add(x)
	}
}

func part1(input []string) int {
	p := processor{register: 1}
	for _, line := range input {
		p.read(line)
	}

	return p.ans
}

func main() {
	input := utils.ReadFileLineByLine("input.txt")

	fmt.Println(part1(input))
}
