package main

import (
	"encoding/json"
	"fmt"

	"github.com/ackurat/advent-of-code/utils/go/utils"
)

func walk(v interface{}, total *float64) {
	switch v := v.(type) {
	case []interface{}:
		for _, v := range v {
			walk(v, total)
		}
	case map[string]interface{}:
		for _, v := range v {
			walk(v, total)
		}
	case float64:
		*total += v
	}
}

func main() {
	input := utils.ReadFileToString("input.txt")
	var result map[string]interface{}
	json.Unmarshal([]byte(input), &result)
	total := 0.0
	walk(result, &total)
	fmt.Println(total)
}
