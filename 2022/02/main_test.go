package main

import "testing"

func TestCalculateArea(t *testing.T) {
	box := Box{2, 3, 4}
	got := calculateArea(box)
	var want int = 52

	if got != want {
		t.Errorf("%d %d", got, want)
	}
}

func TestCalculateWrappingPaperNeed(t *testing.T) {
	box := Box{2, 3, 4}
	got := calculateWrappingPaperNeed(box)
	want := 58

	if got != want {
		t.Errorf("%d %d", got, want)
	}
}
