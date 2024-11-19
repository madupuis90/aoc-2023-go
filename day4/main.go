package main

import (
	"fmt"
	"strings"

	"examples.go/aoc-2023-go/util"
)

func main() {
	part1 := part1("input.txt")
	part2 := part2("input.txt")

	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func part1(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)

	for scanner.Scan() {
		winningNumbers := make(map[string]string)
		subtotal := -1
		line := scanner.Text()
		line = line[strings.IndexRune(line, ':')+1:]
		parts := strings.Split(line, "|")

		for _, n := range strings.Fields(parts[0]) {
			winningNumbers[n] = n
		}

		for _, n := range strings.Fields(parts[1]) {
			_, exists := winningNumbers[n]
			if exists {
				subtotal++
			}
		}

		if subtotal >= 0 { // because we start at -1; 0,1,2,4,8 is what we want instead of 1,2,4,8,16
			total = total + 1<<subtotal // left shift to do 2^subtotal
		}
	}

	return total
}

func part2(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)

	winnings := make(map[int]int) // added cache to speed up things a bit, more optimisation could be done
	scratchCards := make(map[int]string)
	index := 0
	for scanner.Scan() {
		index++
		line := scanner.Text()
		line = line[strings.IndexRune(line, ':')+1:]
		scratchCards[index] = line
	}
	for i := range scratchCards {
		winnings[i] = processScratchcard(scratchCards, winnings, i)
		total += winnings[i]
	}

	return total

}

func processScratchcard(scratchBoards map[int]string, winnings map[int]int, index int) int {

	if value, exists := winnings[index]; exists {
		return value
	}

	w := make(map[string]string)
	invocation := 1
	parts := strings.Split(scratchBoards[index], "|")

	for _, n := range strings.Fields(parts[0]) {
		w[n] = n
	}
	counter := 1
	for _, n := range strings.Fields(parts[1]) {
		_, exists := w[n]
		if exists {
			nextIndex := index + counter
			if nextIndex <= len(scratchBoards) {
				invocation += processScratchcard(scratchBoards, winnings, nextIndex)
				counter++
			}
		}
	}
	return invocation
}
