package utils

import "math"

func SmallestValue(numbers []int) int {
	smallest := math.MaxInt

	for _, v := range numbers {
		if v < smallest {
			smallest = v
		}
	}

	return smallest
}
