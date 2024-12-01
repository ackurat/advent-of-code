package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/ackurat/advent-of-code/utils"
)

type item struct {
	worryLevel int
}

type monkey struct {
	name        int
	items       []item
	operation   func(a, b int) int
	oparg       int
	divisor     int
	trueMonkey  int
	falseMonkey int
	inspections int
}

func (m *monkey) inspect() {
	m.inspections++
}

func evalOp(op string) func(a, b int) int {
	var ops = map[string]func(a, b int) int{
		"*":  func(a, b int) int { return a * b },
		"+":  func(a, b int) int { return a + b },
		"-":  func(a, b int) int { return a - b },
		"**": func(a, b int) int { return a * a },
	}
	return ops[op]
}

func parseInput(name int, input string) monkey {
	var divisor, oparg, falseMonkey, trueMonkey int
	var op string
	reg := regexp.MustCompile(`[^0-9]+`)

	var items []item
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "  Starting items:") {
			fields := strings.Fields(line)
			for i := 2; i < len(fields); i++ {
				fields[i] = reg.ReplaceAllString(fields[i], "")
				itemNumber, _ := strconv.Atoi(fields[i])
				item := item{worryLevel: itemNumber}
				items = append(items, item)
			}
		} else if strings.HasPrefix(line, "  Operation") {
			squareOperator := regexp.MustCompile("old . old")
			if squareOperator.Match([]byte(line)) {
				op = "**"
			} else {
				fmt.Sscanf(line, "  Operation: new = old %s %d", &op, &oparg)
			}
		} else if strings.HasPrefix(line, "  Test") {
			fmt.Sscanf(line, "  Test: divisible by %d", &divisor)
		} else if strings.HasPrefix(line, "    If true") {
			fmt.Sscanf(line, "    If true: throw to monkey %d", &trueMonkey)
		} else if strings.HasPrefix(line, "    If false") {
			fmt.Sscanf(line, "    If false: throw to monkey %d", &falseMonkey)
		}
	}

	return monkey{name: name, items: items, operation: evalOp(op), oparg: oparg, divisor: divisor, trueMonkey: trueMonkey, falseMonkey: falseMonkey, inspections: 0}
}

func part1(monkeys []*monkey, rounds int) int {
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				monkey.inspect()
				item.worryLevel = monkey.operation(item.worryLevel, monkey.oparg)
				item.worryLevel = item.worryLevel / 3
				if item.worryLevel%monkey.divisor == 0 {
					monkeys[monkey.trueMonkey].items = append(monkeys[monkey.trueMonkey].items, item)
					monkey.items = monkey.items[1:]
				} else {
					monkeys[monkey.falseMonkey].items = append(monkeys[monkey.falseMonkey].items, item)
					monkey.items = monkey.items[1:]
				}
			}
		}
	}

	var monkeyInspections []int
	for _, monkey := range monkeys {
		monkeyInspections = append(monkeyInspections, monkey.inspections)
	}
	sort.Slice(monkeyInspections, func(i, j int) bool { return monkeyInspections[i] > monkeyInspections[j] })
	return (monkeyInspections[0] * monkeyInspections[1])
}

func part2(monkeys []*monkey, rounds int) int {
	modfactor := 1
	for _, monkey := range monkeys {
		modfactor *= monkey.divisor
	}

	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				monkey.inspect()
				item.worryLevel = monkey.operation(item.worryLevel, monkey.oparg)
				item.worryLevel = item.worryLevel % modfactor
				if item.worryLevel%monkey.divisor == 0 {
					monkeys[monkey.trueMonkey].items = append(monkeys[monkey.trueMonkey].items, item)
					monkey.items = monkey.items[1:]
				} else {
					monkeys[monkey.falseMonkey].items = append(monkeys[monkey.falseMonkey].items, item)
					monkey.items = monkey.items[1:]
				}
			}
		}
	}

	var itemInspections []int
	for _, monkey := range monkeys {
		itemInspections = append(itemInspections, monkey.inspections)
	}
	sort.Slice(itemInspections, func(i, j int) bool { return itemInspections[i] > itemInspections[j] })
	return (itemInspections[0] * itemInspections[1])
}

func main() {
	input := utils.ReadFileToString("monkeys.txt")

	var monkeys []*monkey

	monkeysSplit := strings.Split(input, "\n\n")
	for i, monkey := range monkeysSplit {
		monk := parseInput(i, monkey)
		monkeys = append(monkeys, &monk)
	}

	fmt.Println(part1(monkeys, 20))
	fmt.Println(part2(monkeys, 10000))
}
