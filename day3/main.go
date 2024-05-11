package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type validRune struct {
	content []rune
	valid bool
}


func main() {
	file, err := os.Open("day3/input") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var engineMatrix []string

	for scanner.Scan() {
		engineMatrix = append(engineMatrix, scanner.Text())
	}

	for line := range engineMatrix {
		fmt.Println(engineMatrix[line])
	}



	var res int
	fmt.Printf("Valor atual: %d\n", res)
	for x, str := range engineMatrix {
		var cur_word validRune
		cur_word.valid = false
		for y, b := range str {
			if unicode.IsDigit(rune(b)) {
				cur_word.content = append(cur_word.content, b)
				if isDigitValid(x, y, engineMatrix){
					cur_word.valid = true
				}
			} else if b == '.' && len(cur_word.content) != 0 {
				if cur_word.valid {
					fmt.Printf("Palavra atual: %c e seu tamanho: %d\n", cur_word.content, len(cur_word.content))
					val, _ := strconv.Atoi(string(cur_word.content))
					res += val
					fmt.Printf("Valor atual: %d\n", res)
				}
				cur_word.content = make([]rune, 0)
				cur_word.valid = false
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
