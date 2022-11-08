package main

import "testing"

func TestPart1(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		got := part1matcher("haegwjzuvuyypxyu")
		var want bool = false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("2", func(t *testing.T) {
		got := part1matcher("afghjkpq")
		var want bool = false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("3", func(t *testing.T) {
		got := part1matcher("jchzalrnumimnmhp")
		var want bool = false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		got := part1matcher("ugknbfddgicrmopn")
		var want bool = true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
