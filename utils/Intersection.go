package utils

import (
	"strings"
)

func OrderedStringIntersection(s1, s2 string) string {
	var intersection strings.Builder
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			intersection.WriteByte(s1[i])
		}
	}
	return intersection.String()
}

func UnOrderedStringIntersection(s ...string) string {
	var intersection strings.Builder
	var sets []Set[string]
	inters := make(map[string]int)

	for _, str := range s {
		strSet := NewSet[string]()
		strSet.AddListOfItems(strings.Split(str, ""))
		sets = append(sets, strSet)
		for char := range strSet {
			inters[char] += 1
		}
	}

	for char, occurrences := range inters {
		if occurrences == len(sets) {
			intersection.WriteString(char)
		}
	}

	return intersection.String()
}
