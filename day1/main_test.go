package main

import (
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample1.txt")
	want := 142

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestSamplePart2(t *testing.T) {
	result := part2("sample2.txt")
	want := 281

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestReverseString(t *testing.T) {
	result := ReverseString("abc123")
	want := "321cba"

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}
