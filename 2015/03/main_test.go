package main

import "testing"

func TestPart1(t *testing.T) {

	t.Run("example 1", func(t *testing.T) {
		input := "^>v<"
		got := part1(input)
		var want int = 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		input := "^v^v^v^v^v"
		got := part1(input)
		var want int = 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		input := "^v"
		got := part2(input)
		var want int = 3

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		input := "^v^v^v^v^v"
		got := part2(input)
		var want int = 11

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
