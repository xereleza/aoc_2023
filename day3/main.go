package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"unicode"
)

func isDigitValid(x int, y int, matrix []string) bool {
	for i := max(x-1, 0); i <= min(len(matrix)-1, x+1); i++ {
		for j := max(y-1, 0); j <= min(len(matrix[i])-1, y+1); j++ {
			if matrix[i][j] != '.' && !unicode.IsNumber(rune(matrix[i][j])) {
				return true
			}
		}
	}
	return false
}

type Gear struct {
	pos [2]int
}

type validRune struct {
	content []rune
	valid   bool
	star    bool
	gears   []Gear
}

func main() {
	file, err := os.Open("day3/input.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var engineMatrix []string

	for scanner.Scan() {
		engineMatrix = append(engineMatrix, scanner.Text())
	}

	/* for line := range engineMatrix {
		fmt.Println(engineMatrix[line])
	} */

	part1(engineMatrix)
	part2(engineMatrix)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func part1(engineMatrix []string) {
	var res int
	fmt.Printf("Valor atual: %d\n", res)
	for x, str := range engineMatrix {
		var cur_word validRune
		cur_word.valid = false
		for y, b := range str {
			if unicode.IsDigit(rune(b)) {
				cur_word.content = append(cur_word.content, b)
				if isDigitValid(x, y, engineMatrix) {
					cur_word.valid = true
				}
			}
			if !unicode.IsDigit(rune(b)) || y == len(str)-1 {
				if cur_word.valid {
					fmt.Printf("Palavra atual: %c e seu tamanho: %d\n", cur_word.content, len(cur_word.content))
					val, err := strconv.Atoi(string(cur_word.content))
					if err != nil {
						log.Fatal(err)
					}
					res += val
					fmt.Printf("Valor atual: %d\n", res)
				}
				cur_word.content = make([]rune, 0)
				cur_word.valid = false
			}
		}
	}
}

func isDigitStar(x int, y int, matrix []string) (bool, [2]int) {
	for i := max(x-1, 0); i <= min(len(matrix)-1, x+1); i++ {
		for j := max(y-1, 0); j <= min(len(matrix[i])-1, y+1); j++ {
			if matrix[i][j] == '*' {
				return true, [2]int{i, j}
			}
		}
	}
	return false, [2]int{}
}

type StarCandidate struct {
	value    int
	pos      [2]int
	star_pos [2]int
}

func part2(engineMatrix []string) {
	var star_candidates []StarCandidate
	for x, str := range engineMatrix {
		var cur_word validRune
		cur_word.valid = false
		cur_word.star = false
		cur_word.gears = []Gear{}
		for y, b := range str {
			if unicode.IsDigit(rune(b)) {
				cur_word.content = append(cur_word.content, b)
				has_star, pos := isDigitStar(x, y, engineMatrix)
				if has_star {
					cur_word.star = true
					cur_gear := Gear{pos}
					cur_word.gears = append(cur_word.gears, cur_gear)
				}
			}
			if !unicode.IsDigit(rune(b)) || y == len(str)-1 {
				if cur_word.star {
					val, err := strconv.Atoi(string(cur_word.content))
					if err != nil {
						log.Fatal(err)
					}
					for _, star := range cur_word.gears {
						candidate := StarCandidate{val, [2]int{x, y - 1}, star.pos}
						if !slices.Contains(star_candidates, candidate) {
							star_candidates = append(star_candidates, candidate)
						}

					}
				}
				cur_word.content = make([]rune, 0)
				cur_word.gears = make([]Gear, 0)
				cur_word.star = false
			}
		}
	}

	//fmt.Println(star_candidates)

	starmap := make(map[[2]int][]int)

	for _, candidate := range star_candidates {
		starmap[candidate.star_pos] = append(starmap[candidate.star_pos], candidate.value)
	}
	//fmt.Println(starmap)
	res := 0
	for _, value := range starmap {
		if len(value) == 2 {
			res += value[0] * value[1]
		}
	}
	fmt.Print(res)
}
