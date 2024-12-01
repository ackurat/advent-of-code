package main

import (
	"fmt"

	"github.com/adamliliemark/advent-of-code/utils"
)

type Reindeer struct {
	Name     string
	Speed    int
	Duration int
	Rest     int
	Distance int
	Flying   int
	Resting  int
	Score    int
}

func (r *Reindeer) tick() {
	if r.Resting > 0 {
		r.Resting -= 1
		if r.Resting == 0 {
			r.Flying = r.Duration
		}
	} else if r.Flying > 0 {
		r.Flying -= 1
		r.Distance += r.Speed
		if r.Flying == 0 {
			r.Resting = r.Rest
		}
	}
}

func part1(reindeers map[*Reindeer]int) int {
	for i := 1; i < 2503; i++ {
		for rein := range reindeers {
			rein.tick()
		}
	}

	var distances []int
	for rein := range reindeers {
		reindeers[rein] = rein.Distance
		distances = append(distances, reindeers[rein])
	}
	return utils.BiggestValue(distances)
}

func part2(reindeers map[*Reindeer]int) int {
	var leadRein string
	for i := 1; i < 1000; i++ {
		for rein := range reindeers {
			rein.tick()
			reindeers[rein] = rein.Distance
		}
		biggest := 0
		for rein := range reindeers {
			// fmt.Println("CHECKING:", rein.Name, rein.Distance, rein.Score)
			if rein.Distance > biggest {
				leadRein = rein.Name
				biggest = rein.Distance
				// fmt.Println("LEADER:", rein.Name, rein.Distance, rein.Score)
			}
		}
		for rein := range reindeers {
			if rein.Name == leadRein {
				rein.Score += 1
			}
		}
	}

	fmt.Println(reindeers)

	return 0
}

func main() {
	input := utils.ReadFileLineByLine("input3.txt")
	var reindeers = make(map[*Reindeer]int)
	for _, line := range input {
		var s, d, r int
		var n string
		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &n, &s, &d, &r)
		reindeer := &Reindeer{Name: n, Speed: s, Duration: d, Rest: r, Distance: 0, Flying: d, Resting: 0, Score: 0}
		reindeers[reindeer] = 0
	}

	// fmt.Println(part1(reindeers))
	fmt.Println(part2(reindeers))

}
