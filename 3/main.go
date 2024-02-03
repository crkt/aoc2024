package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// any number adjacent to a symbol, even diagonally, is a "part number"
// (Periods (.) do not count as a symbol.)

type Node struct {
	next     *Node
	value    rune
	digit    bool
	symbolAt []Pair
}

type List struct {
	head *Node
	tail *Node
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	p1(*scanner)
	// p2(*scanner)
}

func p1(scanner bufio.Scanner) {
	result := [][]rune{}
	for scanner.Scan() {
		row := []rune{}
		text := scanner.Text()
		for _, t := range text {
			row = append(row, t)
		}
		result = append(result, row)
	}
	// Create list of nodes
	sl := List{}
	for i, row := range result {

		for col, t := range row {
			var n Node
			symbolAt := symbolNeigbours(result, i, col)
			digit := runeIsNumber(t)
			n = Node{symbolAt: symbolAt, value: t, digit: digit, next: nil}
			if sl.head == nil {
				sl.head = &n
			} else {
				p := sl.head
				for p.next != nil {
					p = p.next
				}
				p.next = &n
			}
		}
	}

	p := sl.head
	number := ""
	hasSymbol := false
	sum := 0
	for p != nil {
		// when digit, get the chain, and then skip until no digit
		if p.digit {
			number += string(p.value)
			if hasSymbol == false && len(p.symbolAt) > 0 {
				hasSymbol = true
			}
		} else {
			// End of chain. Probably, check if any symbol, if so sum
			if hasSymbol {
				res, _ := strconv.Atoi(number)
				sum += res
			}
			hasSymbol = false
			number = ""
		}
		p = p.next
	}
	fmt.Println("Sum", sum)

}

func p2(scanner bufio.Scanner) {
	result := [][]rune{}
	for scanner.Scan() {
		row := []rune{}
		text := scanner.Text()
		for _, t := range text {
			row = append(row, t)
		}
		result = append(result, row)
	}
	sl := List{}
	for i, row := range result {
		for col, t := range row {
			var n Node
			symbolAt := symbolNeigbours(result, i, col)
			digit := runeIsNumber(t)
			n = Node{symbolAt: symbolAt, value: t, digit: digit, next: nil}
			if sl.head == nil {
				sl.head = &n
			} else {
				p := sl.head
				for p.next != nil {
					p = p.next
				}
				p.next = &n
			}
		}
	}
	p := sl.head
	gearSum := 0

	// Find digit chains that match the same symbol
	number := ""
	symbolAt := []Pair{}
	numbersWithSymbol := make(map[Pair]string)
	for p != nil {
		// While a digit, add to number
		if p.digit {
			number += string(p.value)
		} else {
			number = ""
		}

		// digit, has a symbol
		if len(p.symbolAt) > 0 && p.digit {
			for _, p := range p.symbolAt {
				symbolAt = append(symbolAt, p)
			}
		}

		// Chain will break, remember number if it had any symbol
		if p.next != nil && p.next.digit == false {
			seenSymbols := getUniquePairs(symbolAt)
			for _, pair := range seenSymbols {
				existing, ok := numbersWithSymbol[pair]
				if ok {
					known, _ := strconv.Atoi(existing)
					current, _ := strconv.Atoi(number)
					fmt.Println(known, current)
					gearSum = gearSum + (known * current)
				} else {
					numbersWithSymbol[pair] = number
				}
			}
			symbolAt = []Pair{}
		}

		p = p.next
	}
	// 467835
	fmt.Println("Sum", gearSum)
}

func getUniquePairs(pairs []Pair) []Pair {
	unique := []Pair{}
	seen := map[Pair]bool{}
	for _, p := range pairs {
		if !seen[p] {
			seen[p] = true
			unique = append(unique, p)
		}
	}
	return unique
}

func runeIsNumber(a rune) bool {
	return a >= 48 && a <= 57
}

func isAnySymbol(a []rune) bool {
	any := false
	for _, r := range a {
		if runeIsNumber(r) == false && r != '.' {
			any = true
		}
	}
	return any
}

type Pair struct {
	x, y int
}

func symbolNeigbours(mat [][]rune, row int, col int) []Pair {
	pairs := []Pair{}
	for direction := 0; direction < 9; direction++ {
		if direction == 4 {
			continue // This is ourselves
		}
		n_row := row + ((direction % 3) - 1) // Neighbour row
		n_col := col + ((direction / 3) - 1) // Neighbour col

		if n_row >= 0 && n_row < len(mat) && n_col >= 0 && n_col < len(mat[0]) {
			if isAnySymbol([]rune{mat[n_row][n_col]}) {
				pairs = append(pairs, Pair{x: n_col, y: n_row})
			}
		}
	}

	return pairs
}
