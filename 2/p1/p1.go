package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Game struct {
	Id   string
	Sets map[int]map[string]int
}

func main() {
	// Read file

	// Parse the row, for each game, keeping track of game id and number of cubes

	// Figure out the sum of all games that where possible

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// only 12 red cubes, 13 green cubes, and 14 blue cubes

	games := []Game{}
	for scanner.Scan() {
		text := scanner.Text() // Row
		game, _ := gameFromRow(text)
		games = append(games, game)
	}

	played := []Game{}
	for _, game := range games {
		if gameCouldBePlayed(12, 13, 14, game) {
			played = append(played, game)
		}
	}

	sum_of_played := 0
	for _, game := range played {
		value, err := strconv.Atoi(game.Id)
		fmt.Println(value)
		if err == nil {
			sum_of_played += value
		}
	}

	fmt.Println("Sum of all games", sum_of_played)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

func gameCouldBePlayed(red int, green int, blue int, g Game) bool {
	for game := range g.Sets {
		set := g.Sets[game]
		if set["red"] > red {
			return false
		} else if set["blue"] > blue {
			return false
		} else if set["green"] > green {
			return false
		}
	}
	return true
}

func gameFromRow(row string) (Game, error) {
	result := strings.Split(row, ":")

	game := result[0]
	cubes_map := make(map[int]map[string]int)

	cubes_set := strings.Split(result[1], ";")
	for set_number, cube_set := range cubes_set {
		sets := strings.Split(cube_set, ",")
		for _, set := range sets {
			// Bye bye whitespaces
			trimmed := strings.ReplaceAll(set, " ", "")
			// Find all digits, and keep track of index where digits are
			digits := ""
			digits_at := 0
			for index, c := range trimmed {
				if unicode.IsDigit(c) {
					digits = digits + string(c)
					digits_at = index + 1
				}
			}
			// The color is the part with no digits
			color := trimmed[digits_at:]
			_, ok := cubes_map[set_number]
			// ensure inner map exists
			if !ok {
				cubes_map[set_number] = make(map[string]int)
			}
			amount, err := strconv.Atoi(digits)
			if err != nil {
				return Game{}, errors.New("Not a digit")
			}
			previous_value := cubes_map[set_number][color]
			cubes_map[set_number][color] = previous_value + amount
		}
	}

	g := Game{}
	id, _ := digitFromString(game)
	g.Id = fmt.Sprintf("%d", id)
	g.Sets = cubes_map

	return g, nil
}

func digitFromString(a string) (int, error) {
	digits := ""
	for _, x := range a {
		if unicode.IsDigit(x) {
			digits = digits + string(x)
		}
	}
	value, _ := strconv.Atoi(digits)
	return value, errors.New("No digit found")
}
