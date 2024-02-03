package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	left  []int
	right []int
}

// [Winning numbers] | [Numbers]
// Each match from left side in right side, doubles the amount (1 for the first point, then 2) 4 winning = 8, n = number of winning numbers, result = n * 2 ?
func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rows := []string{}
	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}

	cards := []Card{}
	for _, row := range rows {
		numbers := strings.Split(row, ":")[1]
		splitted := strings.Split(numbers, "|")
		left, right := splitted[0], splitted[1]
		trimmedLeft, trimmedRight := strings.Trim(left, " "), strings.Trim(right, " ")
		leftNumbers := strings.Split(trimmedLeft, " ")
		rightNumbers := strings.Split(trimmedRight, " ")
		cards = append(cards, Card{left: stringsToInts(leftNumbers), right: stringsToInts(rightNumbers)})
	}

	totalWorth := 0
	for i, card := range cards {
		worth := cardWorth(card.left, card.right)
		fmt.Printf("Card %d - %d\n", i+1, worth)
		totalWorth += worth
	}
	fmt.Println(totalWorth)
}

func stringsToInts(strs []string) []int {
	ints := []int{}
	for _, s := range strs {
		// Empty strings are ignored
		if len(s) == 0 {
			continue
		}
		converted, _ := strconv.Atoi(s)
		ints = append(ints, converted)
	}

	return ints
}

func cardWorth(left []int, right []int) int {
	leftSeen := map[int]bool{}
	for _, l := range left {
		leftSeen[l] = true
	}
	matching := 0

	for _, r := range right {
		if leftSeen[r] {
			if matching == 0 {
				matching = 1
			} else {
				matching = matching * 2
			}
		}
	}

	return matching

}
