package utils

import (
	"math"
	"sort"
)

func SmallestValue(numbers []int) int {
	smallest := math.MaxInt

	for _, v := range numbers {
		if v < smallest {
			smallest = v
		}
	}

	return smallest
}

func BiggestValue(numbers []int) int {
	biggest := 0

	for _, v := range numbers {
		if v > biggest {
			biggest = v
		}
	}
	return biggest
}

// Determines if number is in supplied range (inclusive)
func NumberInRange(lower, upper, number int) bool {
	return number >= lower && number <= upper
}

type Pair struct {
	Key   *interface{}
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func BiggestMap(m map[*interface{}]int) (interface{}, int) {
	p := make(PairList, len(m))

	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	biggest, val := p[len(p)].Key, p[len(p)].Value
	return biggest, val
}
