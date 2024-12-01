package main

import (
	"fmt"
	"slices"
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

type Hand struct {
	cards    string
	bid      int
	handType HandType
}

func part1(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)

	hands := []*Hand{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		c := fields[0]
		b, _ := strconv.Atoi(fields[1])
		hands = append(hands, CreateHand(c, b))
	}

	slices.SortFunc(hands, func(h1 *Hand, h2 *Hand) int {
		rankMap := map[rune]int{
			'2': 2,
			'3': 3,
			'4': 4,
			'5': 5,
			'6': 6,
			'7': 7,
			'8': 8,
			'9': 9,
			'T': 10,
			'J': 11,
			'Q': 12,
			'K': 13,
			'A': 14,
		}

		if h1.handType > h2.handType {
			return 1
		} else if h2.handType > h1.handType {
			return -1
		} else { // equal
			for i := 0; i < len(h1.cards); i++ {
				if h1.cards[i] == h2.cards[i] {
					continue
				}
				return rankMap[rune(h1.cards[i])] - rankMap[rune(h2.cards[i])]
			}
		}
		return 0
	})

	for i, h := range hands {
		total = total + h.bid*(i+1)
		// fmt.Printf("%-5v %-15v %-15v %-15v %-15v %-15v\n", h.cards, h.handType.ToString(), h.bid, i+1, h.bid*i+1, total)
	}

	return total
}

type HandType int

func (h HandType) ToString() string {

	switch h {
	case 1:
		return "HIGH_CARD"
	case 2:
		return "PAIR"
	case 3:
		return "TWO_PAIR"
	case 4:
		return "THREE"
	case 5:
		return "FULL_HOUSE"
	case 6:
		return "FOUR"
	case 7:
		return "FIVE"
	}
	return "Unknown"
}

const (
	HIGH_CARD  HandType = 1
	PAIR       HandType = 2
	TWO_PAIR   HandType = 3
	THREE      HandType = 4
	FULL_HOUSE HandType = 5
	FOUR       HandType = 6
	FIVE       HandType = 7
)

func findHandType(cards string) HandType {
	hmap := map[rune]int{
		'2': 0,
		'3': 0,
		'4': 0,
		'5': 0,
		'6': 0,
		'7': 0,
		'8': 0,
		'9': 0,
		'T': 0,
		'J': 0,
		'Q': 0,
		'K': 0,
		'A': 0,
	}
	for _, r := range cards {
		hmap[r]++
	}
	currentType := HIGH_CARD // default
	for _, v := range hmap {
		switch v {
		case 0:
			continue
		case 1:
			continue
		case 2:
			if currentType == PAIR {
				currentType = TWO_PAIR
			} else if currentType == THREE {
				currentType = FULL_HOUSE
			} else {
				currentType = PAIR
			}
		case 3:
			if currentType == PAIR {
				currentType = FULL_HOUSE
			} else {
				currentType = THREE
			}
		case 4:
			currentType = FOUR
		case 5:
			currentType = FIVE
		}
	}
	return currentType
}

func CreateHand(cards string, bid int) *Hand {
	h := findHandType(cards)
	return &Hand{cards: cards, bid: bid, handType: h}
}

func part2(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)

	hands := []*Hand{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		c := fields[0]
		b, _ := strconv.Atoi(fields[1])
		hands = append(hands, CreateHand2(c, b))
	}

	slices.SortFunc(hands, func(h1 *Hand, h2 *Hand) int {
		rankMap := map[rune]int{
			'J': 1,
			'2': 2,
			'3': 3,
			'4': 4,
			'5': 5,
			'6': 6,
			'7': 7,
			'8': 8,
			'9': 9,
			'T': 10,
			'Q': 12,
			'K': 13,
			'A': 14,
		}

		if h1.handType > h2.handType {
			return 1
		} else if h2.handType > h1.handType {
			return -1
		} else { // equal
			for i := 0; i < len(h1.cards); i++ {
				if h1.cards[i] == h2.cards[i] {
					continue
				}
				return rankMap[rune(h1.cards[i])] - rankMap[rune(h2.cards[i])]
			}
		}
		return 0
	})

	for i, h := range hands {
		total = total + h.bid*(i+1)
		fmt.Printf("%-5v %-15v %-15v %-15v %-15v %-15v\n", h.cards, h.handType.ToString(), h.bid, i+1, h.bid*i+1, total)
	}

	return total
}

func CreateHand2(cards string, bid int) *Hand {
	h := findHandType2(cards)
	return &Hand{cards: cards, bid: bid, handType: h}
}

func findHandType2(cards string) HandType {
	hmap := map[rune]int{
		'2': 0,
		'3': 0,
		'4': 0,
		'5': 0,
		'6': 0,
		'7': 0,
		'8': 0,
		'9': 0,
		'T': 0,
		'J': 0,
		'Q': 0,
		'K': 0,
		'A': 0,
	}
	for _, r := range cards {
		hmap[r]++
	}

	// add J to the highest number
	if hmap['J'] != 0 {
		var highest rune
		for k := range hmap {
			if k == 'J' {
				continue
			}
			if hmap[k] > hmap[highest] {
				highest = k
			}
		}
		hmap[highest] = hmap[highest] + hmap['J']
	}

	currentType := HIGH_CARD // default
	for k, v := range hmap {
		if k == 'J' {
			continue
		}

		switch v {
		case 0:
			continue
		case 1:
			continue
		case 2:
			if currentType == PAIR {
				currentType = TWO_PAIR
			} else if currentType == THREE {
				currentType = FULL_HOUSE
			} else {
				currentType = PAIR
			}
		case 3:
			if currentType == PAIR {
				currentType = FULL_HOUSE
			} else {
				currentType = THREE
			}
		case 4:
			currentType = FOUR
		case 5:
			currentType = FIVE
		}
	}
	return currentType
}
