package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	part1()
	part2()

}

func part1() {
	file, err := os.Open("day4/input.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := 0
	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ":")
		has_scored := false
		score := 0
		game_details := strings.Split(game[1], "|")

		winning_num := clean_numbers(game_details[0])

		dealt_hand := clean_numbers(game_details[1])

		for _, num := range dealt_hand {
			if slices.Contains(winning_num, num) {
				if !has_scored {
					score = 1
					has_scored = true
				} else {
					score *= 2
				}
			}
		}
		res += score
	}
	fmt.Println("Part 1:", res)
}

func part2() {
	file, err := os.Open("day4/input.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	record := make(map[int]int)

	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ":")

		game_id, _ := strconv.Atoi(strings.Trim(game[0][5:], " "))
		record[game_id] += 1

		game_details := strings.Split(game[1], "|")

		winning_num := clean_numbers(game_details[0])

		dealt_hand := clean_numbers(game_details[1])

		cur_matches := count_matches(winning_num, dealt_hand)

		for i := 1; i <= cur_matches; i++ {
			for range record[game_id] {
				record[game_id+i] += 1
			}
		}
	}
	res := 0
	for _, v := range record {
		res += v
	}
	fmt.Println("Part 2:", res)
}

func count_matches(winning_num []string, dealt_hand []string) int {
	count := 0
	for _, num := range dealt_hand {
		if slices.Contains(winning_num, num) {
			count += 1
		}
	}
	return count
}

func clean_numbers(numbers_str string) []string {
	return strings.FieldsFunc(numbers_str, func(c rune) bool {
		return c == ' '
	})
}
