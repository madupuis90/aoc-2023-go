package main

import (
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample1.txt")
	want := 288

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestSamplePart2(t *testing.T) {
	result := part2("sample2.txt")
	want := 71503

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestWaysToWin(t *testing.T) {
	result := waysToWin(7, 9)
	want := 4

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}
