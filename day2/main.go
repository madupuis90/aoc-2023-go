package main

import (
	"fmt"
	"regexp"
	"strconv"

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
	blueLimit := 14
	redLimit := 12
	greenLimit := 13

	for scanner.Scan() {
		line := scanner.Text()

		gameRegex := regexp.MustCompile(`^Game (\d+):`)
		gameMatch := gameRegex.FindStringSubmatch(line)
		gameNum, _ := strconv.Atoi(gameMatch[1])

		blueMax := MaxColorInString("blue", line)
		redMax := MaxColorInString("red", line)
		greenMax := MaxColorInString("green", line)

		if blueMax <= blueLimit && redMax <= redLimit && greenMax <= greenLimit {
			total = total + gameNum
		}

	}

	return total
}

func part2(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)

	for scanner.Scan() {
		line := scanner.Text()

		blueMax := MaxColorInString("blue", line)
		redMax := MaxColorInString("red", line)
		greenMax := MaxColorInString("green", line)

		total = total + blueMax*redMax*greenMax

	}

	return total
}

func MaxColorInString(color string, s string) int {
	var max int
	r := regexp.MustCompile(`(\d+) ` + color)
	matches := r.FindAllStringSubmatch(s, -1)

	for _, match := range matches {
		n, _ := strconv.Atoi(match[1])
		if n > max {
			max = n
		}
	}
	return max
}
