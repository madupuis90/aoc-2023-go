package main

import (
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample1.txt")
	want := 8

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestSamplePart2(t *testing.T) {
	result := part2("sample2.txt")
	want := 2286

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestMaxColorInString(t *testing.T) {
	s := "1 blue, 5 blue, 10 blue, 2 red, 20 red, 11 red, 12 green, 16 green 30 green"
	blueWant := 10
	blueMax := MaxColorInString("blue", s)
	redWant := 20
	redMax := MaxColorInString("red", s)
	greenWant := 30
	greenMax := MaxColorInString("green", s)

	if blueMax != blueWant || redMax != redWant || greenMax != greenWant {
		t.Fatalf(`Wanted blue: %v, red %v, green %v; but got blue: %v, red %v, green %v`, blueWant, redWant, greenWant, blueMax, redMax, greenMax)
	}
}
