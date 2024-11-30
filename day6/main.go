package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"examples.go/aoc-2023-go/util"
)

func main() {
	start := time.Now()
	part1 := part1("input.txt")
	part2 := part2("input.txt")

	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
	fmt.Printf("Program took %s\n", time.Since(start))
}

func part1(f string) int {
	total := 1
	scanner := util.CreateScannerFromFile(f)
	time := []int{}
	distance := []int{}

	if scanner.Scan() {
		line := scanner.Text()
		line = line[strings.IndexRune(line, ':')+1:]
		time = ParseStringAsIntSlice(line)
	}
	if scanner.Scan() {
		line := scanner.Text()
		line = line[strings.IndexRune(line, ':')+1:]
		distance = ParseStringAsIntSlice(line)
	}

	for i := 0; i < len(time); i++ {
		total = total * waysToWin(time[i], distance[i])
	}

	return total
}

func ParseStringAsIntSlice(line string) []int {
	s := []int{}
	for _, t := range strings.Fields(line) {
		tt, _ := strconv.Atoi(t)
		s = append(s, tt)
	}
	return s
}

func waysToWin(time int, distance int) int {
	ways := 0
	center := time / 2
	left, right := center, center

	for left >= 0 && right <= time {
		l := (time-left)*left > distance
		if l {
			ways++
		}
		r := (time-right)*right > distance
		if r && left != right {
			ways++
		}
		if l && r {
			left--
			right++
			continue
		}
		break
	}
	return ways
}

func part2(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)
	var time int
	var distance int

	if scanner.Scan() {
		line := scanner.Text()
		line = line[strings.IndexRune(line, ':')+1:]
		line = strings.Join(strings.Fields(line), "")
		time, _ = strconv.Atoi(line)
	}
	if scanner.Scan() {
		line := scanner.Text()
		line = line[strings.IndexRune(line, ':')+1:]
		line = strings.Join(strings.Fields(line), "")
		distance, _ = strconv.Atoi(line)
	}

	total = waysToWin(time, distance)

	return total
}
