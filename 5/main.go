package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
	origin      string
	destination string
}

type ParsedRange struct {
	destination int
	origin      int
	delta       int
}

type SeedRange struct {
	x1      int
	x2      int
	segment int
}

// Create a structure where the 'maps' found in the data contains the ranges/deltas with a destination and origin
func parse(scanner *bufio.Scanner) map[Mapping][]ParsedRange {
	maps := make(map[Mapping][]ParsedRange)
	var current Mapping

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		var text = scanner.Text()
		var fields = strings.Fields(text)

		if strings.Contains(text, "map") {
			str := strings.Split(fields[0], "-to-")
			m := Mapping{origin: str[0], destination: str[1]}
			current = m
			maps[current] = []ParsedRange{}
		} else {
			p := ParsedRange{}
			p.destination, _ = strconv.Atoi(fields[0])
			p.origin, _ = strconv.Atoi(fields[1])
			p.delta, _ = strconv.Atoi(fields[2])
			maps[current] = append(maps[current], p)
		}
	}

	return maps

}

func calculate(maps map[Mapping][]ParsedRange, seeds []string) int {
	var current int
	var locations []int

	for _, seed := range seeds {
		state := "seed"
		current, _ = strconv.Atoi(seed)
		for state != "location" { // Location would be the last map to check
			for m := range maps {
				if m.origin == state {
					for _, p := range maps[m] {
						// If we have a mapped value that covers the current number - look it up
						if current >= p.origin && current <= p.origin+p.delta {
							current = p.destination + (current - p.origin)
							break
						}
					}
					state = m.destination
					break
				}
			}
		}
		locations = append(locations, current)
	}

	// Find the lower location for any given seed used - check with first location against other locations
	lowest := locations[0]
	for _, loc := range locations {
		if loc < lowest {
			lowest = loc
		}
	}

	return lowest
}

func part1(s *bufio.Scanner) int {
	s.Scan()
	seeds := strings.Fields(strings.Split(s.Text(), ":")[1])
	maps := parse(s)
	lowest := calculate(maps, seeds)

	return lowest
}

func part2(s *bufio.Scanner) int {
	s.Scan()
	seeding := strings.Fields(strings.Split(s.Text(), ":")[1])
	var intervals []SeedRange

	// Create intervals with a start segment value
	for i := 0; i < len(seeding); i += 2 {
		start, _ := strconv.Atoi(seeding[i])
		delta, _ := strconv.Atoi(seeding[i+1])
		x2 := delta + start
		seedRange := SeedRange{x1: start, x2: x2, segment: 1}
		intervals = append(intervals, seedRange)
	}

	ranges := parse(s)

	location_nums := []int{}
	for {
		seedRange := intervals[0]

		// We have checked 8 depths for this seedRange, so we are done with it
		if seedRange.segment == 8 {
			continue
		}

	}

	lowest := location_nums[0]
	for _, loc := range location_nums {
		if loc < lowest {
			lowest = loc
		}
	}

	return lowest
}

func main() {
	file, _ := os.Open("test.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//result := part1(scanner)
	result := part2(scanner)
	fmt.Println("Lowest location", result)

}
