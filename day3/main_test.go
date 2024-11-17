package main

import (
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample1.txt")
	want := 4361

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestSamplePart2(t *testing.T) {
	result := part2("sample2.txt")
	want := 467835

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}