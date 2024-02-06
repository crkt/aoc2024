package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	number int
	left   []int
	right  []int
	wins   []int
}

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

	// [Winning numbers] | [Numbers]
	cards := []Card{}
	for i, row := range rows {
		numbers := strings.Split(row, ":")[1]
		splitted := strings.Split(numbers, "|")
		pipeLeft, pipeRight := splitted[0], splitted[1]
		trimmedLeft, trimmedRight := strings.Trim(pipeLeft, " "), strings.Trim(pipeRight, " ")
		leftNumbers := strings.Split(trimmedLeft, " ")
		rightNumbers := strings.Split(trimmedRight, " ")
		number := i + 1
		left, right := stringsToInts(leftNumbers), stringsToInts(rightNumbers)
		won := cardWinningCards(number, left, right)
		card := Card{number: number, left: left, right: right, wins: won}
		cards = append(cards, card)
	}

	//p1(cards)
	p2(cards)
}

func p1(cards []Card) {
	totalWorth := 0
	for i, card := range cards {
		worth := cardWorth(card.left, card.right)
		fmt.Printf("Card %d - %d\n", i+1, worth)
		totalWorth += worth
	}
	fmt.Println(totalWorth)
}

func p2(cards []Card) {
	copies := map[int]int{}
	for _, c := range cards {
		copies[c.number] += 1
		for _, w := range c.wins {
			// We win the amount of existing copies
			copies[w] += copies[c.number]
		}
	}
	sum := 0
	for _, c := range copies {
		sum += c
	}
	fmt.Println(sum)
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

/**
* Returns the number of copies won for this scratchcard
 */
func cardWinningCards(number int, left []int, right []int) []int {
	leftSeen := map[int]bool{}
	for _, l := range left {
		leftSeen[l] = true
	}

	matching := []int{}
	for _, r := range right {
		if leftSeen[r] {
			// Looks wierd but, you win the "next" card and never the current, so plus one
			matching = append(matching, len(matching)+1+number)
		}
	}

	return matching
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
