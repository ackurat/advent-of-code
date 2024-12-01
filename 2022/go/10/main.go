package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ackurat/advent-of-code/utils"
)

type processor struct {
	register, cycle, part1 int
	part2                  string
}

func (p *processor) tick() {
	p.cycle++
	if p.cycle%40 == 20 {
		p.part1 += p.cycle * p.register
	}
	pixel := (p.cycle - 1) % 40
	if pixel == 0 {
		p.part2 += "\n"
	}
	if utils.NumberInRange(p.register-1, p.register+1, pixel) {
		p.part2 += "\u2588"
	} else {
		p.part2 += " "
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

func main() {
	input := utils.ReadFileLineByLine("input.txt")
	p := processor{register: 1}
	for _, line := range input {
		p.read(line)
	}
	fmt.Println(p.part1)
	fmt.Println(p.part2)
}
