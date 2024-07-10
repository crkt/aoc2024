package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	input := string(b)

	segments := strings.Split(input, "\n\n")

	seeds := segments[0]

	fmt.Println(seeds)

	mappingPattern := regexp.MustCompile("(\\d+) (\\d+) (\\d+)")

	seedPattern := regexp.MustCompile("(\\d+) (\\d+)")
	seedMatches := seedPattern.FindAllString(seeds, -1)
	intervals := [][]int{}
	for _, match := range seedMatches {
		res := strings.Split(match, " ")
		x1, _ := strconv.Atoi(res[0])
		delta, _ := strconv.Atoi(res[1])
		x2 := x1 + delta
		intervals = append(intervals, []int{x1, x2, 1})
	}

	location_nums := []int{}

	fmt.Println(intervals)

	for len(intervals) > 0 {
		x1, x2, segment_idx := intervals[0][0], intervals[0][1], intervals[0][2]
		intervals = append(intervals[:0], intervals[1:]...)

		// Last map has been handled
		if segment_idx == 8 {
			location_nums = append(location_nums, x1)
			continue
		}

		// Segment here is the maps of ranges for a given mapping
		segment := segments[segment_idx]
		mappingMatches := mappingPattern.FindAllString(segment, -1)
		for _, mapping := range mappingMatches {
			mapRes := strings.Split(mapping, " ")
			destination, _ := strconv.Atoi(mapRes[0])
			start, _ := strconv.Atoi(mapRes[1])
			delta, _ := strconv.Atoi(mapRes[2])
			end := start + delta
			diff := destination - start
			if x1 >= end || x2 <= start { // no overlap found - move on to the next segment
				continue
			}
			if x1 < start {
				intervals = append(intervals, []int{x1, start, segment_idx})
				x1 = start
			}
			if x2 > end {
				intervals = append(intervals, []int{end, x2, segment_idx})
				x2 = end
			}
			intervals = append(intervals, []int{x1 + diff, x2 + diff, segment_idx + 1})
			break
		}
		intervals = append(intervals, []int{x1, x2, segment_idx + 1})
		// Get all the mappings for the segment index in question

	}

	// Find min value
	fmt.Println(location_nums)
	lowest := location_nums[0]
	for _, l := range location_nums {
		if l < lowest {
			lowest = l
		}
	}
	fmt.Println(lowest)

}
