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
	value    rune
	digit    bool
	symbolAt []Pair
}

type Pair struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// p1(*scanner)
	p2(*scanner)
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
	nodes := getNodesFromInput(result)

	number := ""
	hasSymbol := false
	sum := 0
	for _, p := range nodes {
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
	// Find neigbours for symbols and numbers
	nodes := getNodesFromInput(result)

	// Find digit chains that match the same symbol
	gearSum := 0
	number := ""
	symbolAt := []Pair{}
	numbersWithSymbol := make(map[Pair]string)
	for _, p := range nodes {

		// While a digit, add to number
		if p.digit {
			number += string(p.value)
			if len(p.symbolAt) > 0 {
				for _, p := range p.symbolAt {
					symbolAt = append(symbolAt, p)
				}
			}
		} else {
			// End of number chain, check if any symbols
			seenSymbols := getUniquePairs(symbolAt)
			for _, pair := range seenSymbols {
				existing, ok := numbersWithSymbol[pair]
				// Other node has set this pair already, we know we have a gear
				if ok {
					known, _ := strconv.Atoi(existing)
					current, _ := strconv.Atoi(number)
					fmt.Println(known, current)
					gearSum = gearSum + (known * current)
				} else {
					numbersWithSymbol[pair] = number
				}
			}
			// Restore 'state' after checking result
			symbolAt = []Pair{}
			number = ""
		}
	}
	fmt.Println("Sum", gearSum)
}

func getNodesFromInput(input [][]rune) []Node {
	nodes := []Node{}
	for i, row := range input {
		for col, t := range row {
			var n Node
			symbolAt := symbolNeigbours(input, i, col)
			digit := runeIsNumber(t)
			n = Node{symbolAt: symbolAt, value: t, digit: digit}
			nodes = append(nodes, n)
		}
	}
	return nodes
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
