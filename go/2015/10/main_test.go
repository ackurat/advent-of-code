package main

import "testing"

func TestPart1(t *testing.T) {

	t.Run("example 1", func(t *testing.T) {
		input := "1"
		got := part1And2(input, 1)
		var want string = "11"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		input := "1211"
		got := part1And2(input, 1)
		var want string = "111221"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
