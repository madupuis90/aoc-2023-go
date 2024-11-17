package main

import (
	"fmt"
	"strconv"
	"unicode"

	"examples.go/aoc-2023-go/util"
)

type Graph2D = [][]rune

func main() {
	part1 := part1("input.txt")
	part2 := part2("input.txt")

	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func part1(f string) int {
	graph2D := createGraph2DFromFile(f)
	total := 0
	accumulatorString := ""
	isAdjacent := false

	// Need to do this logic on new row and when at the end of a digit string
	addAndReset := func() {
		if isAdjacent {
			num, _ := strconv.Atoi(accumulatorString) // get the numerical value of the string
			total = total + num                       // add the value to the total
		}
		accumulatorString = "" // reset the accumulator
		isAdjacent = false     // reset the flag
	}

	for i, row := range graph2D {
		addAndReset()
		for j, r := range row {
			if unicode.IsDigit(r) {
				accumulatorString += string(r) // concat all numbers
				if !isAdjacent {               // if not already adjacent
					isAdjacent, _ = findRuneAdjacentTo(graph2D, Vertex{i, j}, isSymbol) // set flag if any number is adjacent
				}
			} else { // end of digit string
				addAndReset()
			}
		}
	}
	return total
}

func part2(f string) int {
	graph2D := createGraph2DFromFile(f)
	total := 0

	numString1 := ""
	numString2 := ""

	for i := range graph2D {
		for j := range graph2D[i] {
			if graph2D[i][j] == '*' {
				// check surroundings and replace numbers with .
				found, p := findRuneAdjacentTo(graph2D, Vertex{i, j}, unicode.IsDigit)
				if found {
					numString1 = findSuroundingNumberString(graph2D, p)
				} else {
					numString1 = ""
				}

				// check srroundings a second time for the second number
				found, p = findRuneAdjacentTo(graph2D, Vertex{i, j}, unicode.IsDigit)
				if found {
					numString2 = findSuroundingNumberString(graph2D, p)
				} else {
					numString2 = ""
				}

				// if 2 numbers were found add their product
				if numString1 != "" && numString2 != "" {
					num1, _ := strconv.Atoi(numString1)
					num2, _ := strconv.Atoi(numString2)
					total = total + num1*num2
				}
			}
		}
	}
	return total

}

func findSuroundingNumberString(graph2D Graph2D, p Vertex) string {
	numString := ""
	numString = string(graph2D[p.X][p.Y])
	graph2D[p.X][p.Y] = '.'
	// check left for numbers
	c := 0
	for {
		c++
		if p.Y-c >= 0 { // check bounds
			r := graph2D[p.X][p.Y-c]
			if unicode.IsDigit(r) {
				numString = string(r) + numString
				graph2D[p.X][p.Y-c] = '.'
			} else {
				break
			}
		} else {
			break
		}
	}
	// check right for numbers
	c = 0
	for {
		c++
		if p.Y+c <= int(len(graph2D[p.X])-1) { // check bounds
			r := graph2D[p.X][p.Y+c]
			if unicode.IsDigit(r) {
				numString = numString + string(r)
				graph2D[p.X][p.Y+c] = '.'
			} else {
				break
			}
		} else {
			break
		}
	}
	return numString
}

func part2_sampleWork_notInput(f string) int {

	graph2D := createGraph2DFromFile(f)
	total := 0
	lastAccumulatorString := ""
	accumulatorString := ""
	isDollarAdjacent := false
	isStarAdjacent := false

	// The idea is to remove all symbols that do not represent a gear, leaving only '*'
	// (1) we will check for numbers adjacent to '*' and replace the '*' with '$' when adjacent first time
	// (2) we will check for numbers adjacent to '$' --> this means a gear
	// edit: this ended up being more complicated than I thought when I realized that the 2 accumulator won't
	// be sequential... this solution do not really work. I will leave it up, for history sake.

	// Need to do this logic on new row and when at the end of a digit string
	addAndReset := func() {

		if accumulatorString != "" {
			fmt.Printf("%-10v %-10v %-5v %-5v \n", accumulatorString, lastAccumulatorString, isDollarAdjacent, isStarAdjacent)
		}

		if isDollarAdjacent {
			num, _ := strconv.Atoi(lastAccumulatorString) // get the numerical value of the string
			num2, _ := strconv.Atoi(accumulatorString)
			total = total + num*num2 // add the value to the total
			lastAccumulatorString = ""
		}
		if isStarAdjacent {
			lastAccumulatorString = accumulatorString
		}
		accumulatorString = ""   // reset the accumulator
		isDollarAdjacent = false // reset the flag
		isStarAdjacent = false
	}

	normalizeGraph2D(graph2D)

	for i, row := range graph2D {
		addAndReset()
		for j, r := range row {
			if unicode.IsDigit(r) {
				accumulatorString += string(r)
				if !isDollarAdjacent {
					isDollarAdjacent, _ = findRuneAdjacentTo(graph2D, Vertex{i, j}, isDollar)
				}
				found, starPos := findRuneAdjacentTo(graph2D, Vertex{i, j}, isStar)
				if found {
					graph2D[starPos.X][starPos.Y] = '$'
					if !isStarAdjacent {
						isStarAdjacent = found
					}
				}
			} else {
				addAndReset()
			}
		}
	}
	// printGraph2D(graph2D)

	return total
}

func createGraph2DFromFile(f string) Graph2D {
	scanner := util.CreateScannerFromFile(f)
	var graph2D [][]rune
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		graph2D = append(graph2D, runes)
	}
	return graph2D
}

type runePredicate func(r rune) bool

func isStar(r rune) bool {
	return r == '*'
}
func isDollar(r rune) bool {
	return r == '$'
}
func isSymbol(r rune) bool {
	// Sample did not cover every symbol, I got annoyed at always missing a symbol so negation check it is!
	return !unicode.IsDigit(r) && r != '.'
}

type Vertex struct {
	X int
	Y int
}

func findRuneAdjacentTo(g Graph2D, r Vertex, p runePredicate) (bool, Vertex) {
	for i := r.X - 1; i <= r.X+1; i++ {
		if i < 0 || i >= len(g) { // check out of bound
			continue
		}
		for j := r.Y - 1; j <= r.Y+1; j++ {
			if j < 0 || j >= len(g[r.X]) { // check out of bound
				continue
			}
			if p(g[i][j]) {
				return true, Vertex{i, j}
			}
		}
	}
	return false, Vertex{}
}

// (part2) Remove all symbols that are not a digit, star or dot
// since we don't care anymore about other symbol --> allows using symbols as marker
func normalizeGraph2D(graph2D Graph2D) {
	for i, row := range graph2D {
		for j, r := range row {
			if !unicode.IsDigit(r) && r != '*' {
				graph2D[i][j] = '.'
			}
		}
	}
}

func printGraph2D(graph2D Graph2D) {
	for _, row := range graph2D {
		fmt.Printf("\n")
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Printf("\n")
	}
}
