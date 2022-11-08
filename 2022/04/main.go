package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func part1(input string) int {

	i := 1
	for {
		in := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(in))
		hashSlice := hash[:3]
		if hex.EncodeToString(hashSlice) < "000009" {
			fmt.Println(hashSlice)
			fmt.Println(hex.EncodeToString(hashSlice))
			break
		}
		i += 1
	}

	return i
}

func part2(input string) int {
	i := 1
	for {
		in := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(in))
		hashSlice := hash[:3]
		if bytes.Equal([]byte{0, 0, 0}, hashSlice) {
			break
		}
		i += 1
	}

	return i
}

func main() {
	input := "bgvyzdsv"

	// fmt.Println(part1(input))
	fmt.Println(part2(input))

}
