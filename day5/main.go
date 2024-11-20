package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"examples.go/aoc-2023-go/util"
)

func main() {
	start := time.Now()
	// part1 := part1("input.txt")
	part2 := part2("input.txt")

	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
	fmt.Printf("Program took %s\n", time.Since(start))
}

func part1(f string) int {
	scanner := util.CreateScannerFromFile(f)
	seeds := []int{}
	var mapFilters []*MapFilter

	if scanner.Scan() {
		seeds = parseSeeds(scanner.Text())
	}

	for scanner.Scan() {
		line := scanner.Text()

		// skip empty lines
		if line == "" {
			continue
		}

		// create a new map
		if strings.Contains(line, "map") {
			newMapFilter := createMapFilter()
			if len(mapFilters) > 0 {
				mapFilters[len(mapFilters)-1].nextMap = newMapFilter
			}
			mapFilters = append(mapFilters, newMapFilter)
			continue
		}

		// process numbers
		currentMap := mapFilters[len(mapFilters)-1]
		currentMap.filters = append(currentMap.filters, createFilter(line))
	}
	minValue := math.MaxInt
	var value int
	for _, s := range seeds {
		value = mapFilters[0].compute(s)
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}

func (mf *MapFilter) compute(seed int) int {
	value := seed
	for _, f := range mf.filters {
		if value >= f.lowerBound && value <= f.upperBound {
			value = value + f.boundDiff
			break
		}
	}
	if mf.nextMap != nil {
		value = mf.nextMap.compute(value)
	}
	return value
}

func createMapFilter() *MapFilter {
	return &MapFilter{
		filters: []Filter{},
		nextMap: nil,
	}
}

func createFilter(line string) Filter {
	nums := strings.Fields(line)
	dest, _ := strconv.Atoi(nums[0])
	src, _ := strconv.Atoi(nums[1])
	length, _ := strconv.Atoi(nums[2])

	return Filter{
		upperBound: src + length - 1,
		lowerBound: src,
		boundDiff:  dest - src,
	}
}

type SeedRange struct {
	lowerBound int
	upperBound int
}

type MapFilter struct {
	filters []Filter
	nextMap *MapFilter
}

type Filter struct {
	lowerBound int
	upperBound int
	boundDiff  int
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	seeds := []int{}
	var mapFilters []*MapFilter

	if scanner.Scan() {
		seeds = parseSeeds(scanner.Text())
	}

	for scanner.Scan() {
		line := scanner.Text()

		// skip empty lines
		if line == "" {
			continue
		}

		// create a new map
		if strings.Contains(line, "map") {
			newMapFilter := createMapFilter()
			if len(mapFilters) > 0 {
				mapFilters[len(mapFilters)-1].nextMap = newMapFilter
			}
			mapFilters = append(mapFilters, newMapFilter)
			continue
		}

		// process numbers
		currentMap := mapFilters[len(mapFilters)-1]
		currentMap.filters = append(currentMap.filters, createFilter(line))
	}
	minValue := math.MaxInt
	var value int
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j <= seeds[i]+seeds[i+1]-1; j++ {
			value = mapFilters[0].compute(j)
			if value < minValue {
				minValue = value
			}
		}
	}

	return minValue

}

// Creates map containing all other maps with all possible values... this works for the sample fine because the
// data set is small, but it doesn't work for the real input... it actually uses 60gb of memory. lol
// I will keep it there for history sake
func part1_very_ineficient_way(f string) int {
	scanner := util.CreateScannerFromFile(f)
	seeds := []int{}
	maps := make(map[int]map[int]int)
	mapCounter := -1

	if scanner.Scan() {
		seeds = parseSeeds(scanner.Text())
	}
	for scanner.Scan() {
		line := scanner.Text()

		// skip empty lines
		if line == "" {
			continue
		}

		// increase the map counter; we will use int instead of strings to represents all maps
		if strings.Contains(line, "map") {
			mapCounter++
			maps[mapCounter] = make(map[int]int)
			continue
		}
		populateMap(maps[mapCounter], line)
	}

	var minValue = math.MaxInt
	for _, s := range seeds {
		value := walkMapsToFindLocation(maps, s)
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}

func walkMapsToFindLocation(maps map[int]map[int]int, seed int) int {
	var value, oldValue int
	var exists bool

	value, oldValue = seed, seed
	for i := 0; i < len(maps); i++ {
		value, exists = maps[i][value]
		if exists {
			oldValue = value
		} else {
			value = oldValue
		}
	}
	return value
}

func populateMap(m map[int]int, line string) {
	nums := strings.Fields(line)
	dest, _ := strconv.Atoi(nums[0])
	src, _ := strconv.Atoi(nums[1])
	length, _ := strconv.Atoi(nums[2])

	for i := 0; i < length; i++ {
		m[src+i] = dest + i
	}
}

func parseSeeds(line string) []int {
	var seeds []int
	line = line[strings.IndexRune(line, ':')+1:]
	for _, seed := range strings.Fields(line) {
		s, _ := strconv.Atoi(seed)
		seeds = append(seeds, s)
	}
	return seeds
}
