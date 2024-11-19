package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode"

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

	// Iterate over lines, build array of digits to get first and last
	for scanner.Scan() {
		var digits []string
		line := scanner.Text()
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, string(char))
			}
		}
		if len(digits) == 0 {
			log.Fatalf("Expected line to contain at least one digit")
		}

		first := digits[0]
		last := digits[len(digits)-1]
		digit, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatalf("Expected digits to be only numerical values")
		}
		total = total + digit
	}

	return total
}

func part2(f string) int {
	total := 0
	spelledDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"eno":   "1",
		"owt":   "2",
		"eerht": "3",
		"ruof":  "4",
		"evif":  "5",
		"xis":   "6",
		"neves": "7",
		"thgie": "8",
		"enin":  "9",
	}
	scanner := util.CreateScannerFromFile(f)

	/*
		I was initially using a single regex and building a slice of all matches just line in part one, but I couldn't
		get the right answer with the input altough the sample was fine... I ended up googling and found out that there are
		a few strings with overlaps 'eighthree'... This would require for me to modify my regex to include a look ahead, but
		GO does not support it in STD lib. I switched to this janky solution where I look for the first match using the normal
		regex, then look for the first match in the reverse string using the reverse regex which effectively runs the regex
		from the end.
	*/
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	rre := regexp.MustCompile(`(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|\d)`) // reverse regex
	for scanner.Scan() {
		line := scanner.Text()
		rline := ReverseString(line)

		first := re.FindString(line)
		last := rre.FindString(rline) // reverse regex + reverse line

		value, exists := spelledDigits[first]
		if exists {
			first = value
		}
		value, exists = spelledDigits[last]
		if exists {
			last = value
		}

		digit, err := strconv.Atoi(first + last)

		if err != nil {
			log.Fatalf("Expected digits to be only numerical values")
		}
		total = total + digit
	}
	return total
}

func ReverseString(s string) string {
	l := len(s)
	s2 := make([]rune, l)
	for i, char := range s {
		s2[l-i-1] = char
	}
	return string(s2)
}
