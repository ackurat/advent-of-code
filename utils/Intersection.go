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

func GenericIntersection[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}
