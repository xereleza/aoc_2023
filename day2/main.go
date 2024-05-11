package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Contents struct {
	red   int
	green int
	blue  int
}

func compute_game(game string, result *int) {
	//CONTENTS := Contents{red: 12, green: 13, blue: 14}
	fmt.Println(game)
	arr := strings.Split(game, ":")
	game_id, _ := strconv.Atoi(arr[0][5:])
	game_details := strings.Split(arr[1], ";")
	if compute_set(game_details) {
		*result += game_id
	}
}

func compute_game2(game string, result *int) {
	found := &Contents{red: 0, green: 0, blue: 0}
	arr := strings.Split(game, ":")
	//game_id, _ := strconv.Atoi(arr[0][5:])
	game_details := strings.Split(arr[1], ";")
	compute_minimal(game_details, found)
	*result += compute_power(found)
}

func compute_power(game_set *Contents) int {
	return game_set.red * game_set.green * game_set.blue
}

func compute_minimal(details []string, found *Contents) {
	for _, item := range details {
		split := strings.Split(strings.TrimPrefix(item, " "), " ")
		var qtde int
		for index, texto := range split {
			if index%2 == 0 {
				qtde, _ = strconv.Atoi(texto)
			}
			if index%2 != 0 {
				switch string(texto[0]) {
				case "r":
					if qtde > found.red {
						found.red = qtde
					}
				case "g":
					if qtde > found.green {
						found.green = qtde
					}
				case "b":
					if qtde > found.blue {
						found.blue = qtde
					}
				}
			}
		}
	}
}

func compute_set(set []string) bool {
	for _, item := range set {
		split := strings.Split(strings.TrimPrefix(item, " "), " ")
		var qtde int
		for index, texto := range split {
			if index%2 == 0 {
				qtde, _ = strconv.Atoi(texto)
			}

			if index%2 != 0 {
				switch string(texto[0]) {
				case "r":
					if qtde > 12 {
						return false
					}
				case "g":
					if qtde > 13 {
						return false
					}
				case "b":
					if qtde > 14 {
						return false
					}
				}
			}
		}
		fmt.Println(split)
	}
	return true
}

func main() {
	final := 0
	file, err := os.Open("input") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	/* for scanner.Scan() {
		compute_game(scanner.Text(), &final)
	}

	fmt.Printf("Part 1: " + strconv.Itoa(final))

	final = 0 */

	for scanner.Scan() {
		compute_game2(scanner.Text(), &final)
	}

	fmt.Printf("Part 2: " + strconv.Itoa(final))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
