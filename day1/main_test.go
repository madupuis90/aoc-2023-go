package day1

import (
	"fmt"
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample1.txt")
	want := 142
	fmt.Println(result)

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestSamplePart2(t *testing.T) {
	result := part2("sample2.txt")
	want := 281
	fmt.Println(result)

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestInputPart1(t *testing.T) {
	result := part1("input.txt")
	fmt.Println(result)
}

func TestInputPart2(t *testing.T) {
	result := part2("input.txt")
	fmt.Println(result)
}

func TestReverseString(t *testing.T) {
	result := ReverseString("abc123")
	want := "321cba"

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}
