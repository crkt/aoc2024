package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		digs := firstAndLastDigit(text)
		sum += digs
	}

	fmt.Println("Sum", sum)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func firstAndLastDigit(row string) int {
	map_of_ints := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	// Add such digit/word to line
	digits := []int{}
	for index, char := range row {
		if unicode.IsDigit(char) {
			digits = append(digits, int(char-'0'))
		} else {
			for k := range map_of_ints {
				// Slice the row with the current index, check prefix
				sub := row[index:]
				if strings.HasPrefix(sub, k) {
					digits = append(digits, map_of_ints[k])
				}
			}

		}
	}

	length := len(digits)
	if length >= 1 {
		first := digits[0]
		last := digits[length-1]
		double_digit := fmt.Sprintf("%d%d", first, last)
		result, err := strconv.Atoi(double_digit)
		if err == nil {
			return result
		} else {
			return 0
		}
	} else {
		return 0
	}
}
