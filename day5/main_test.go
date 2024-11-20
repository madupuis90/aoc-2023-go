package main

import (
	"reflect"
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample1.txt")
	want := 35

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestSamplePart2(t *testing.T) {
	result := part2("sample2.txt")
	want := 46

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestParseSeeds(t *testing.T) {
	results := parseSeeds("seeds: 1 2 10 20 100 555")
	want := []int{1, 2, 10, 20, 100, 555}

	if !reflect.DeepEqual(results, want) {
		t.Fatalf(`Wanted %v, but got %v`, want, results)
	}
}

func TestPopulateMap(t *testing.T) {
	results := map[int]int{}
	populateMap(results, "50 98 2")
	want := map[int]int{98: 50, 99: 51}

	if !reflect.DeepEqual(results, want) {
		t.Fatalf(`Wanted %v, but got %v`, want, results)
	}
}

func TestPopulateMaps(t *testing.T) {
	results := map[int]int{}
	populateMap(results, "50 98 2")
	populateMap(results, "52 50 48")
	want := map[int]int{50: 52, 51: 53, 52: 54, 53: 55, 54: 56, 55: 57, 56: 58, 57: 59, 58: 60, 59: 61, 60: 62, 61: 63, 62: 64, 63: 65, 64: 66, 65: 67, 66: 68, 67: 69, 68: 70, 69: 71, 70: 72, 71: 73, 72: 74, 73: 75, 74: 76, 75: 77, 76: 78, 77: 79, 78: 80, 79: 81, 80: 82, 81: 83, 82: 84, 83: 85, 84: 86, 85: 87, 86: 88, 87: 89, 88: 90, 89: 91, 90: 92, 91: 93, 92: 94, 93: 95, 94: 96, 95: 97, 96: 98, 97: 99, 98: 50, 99: 51}

	if !reflect.DeepEqual(results, want) {
		t.Fatalf(`Wanted %v, but got %v`, want, results)
	}
}
