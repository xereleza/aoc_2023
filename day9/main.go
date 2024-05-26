package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	res1, res2 := 0, 0
	for _, nums := range load("day9/input.txt") {
		res1 += extrapolate(diff(nums, [][]int{nums}))
		res2 += extrapolate_backwards(diff(nums, [][]int{nums}))
	}
	fmt.Println("Part 1:", res1)
	fmt.Println("Part 2:", res2)
}

func load(input_path string) [][]int {
	file, err := os.Open(input_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	contents := strings.Split(string(buf), "\n")
	ints := make([][]int, len(contents))
	for i, history := range contents {
		for _, num := range strings.Split(history, " ") {
			n, _ := strconv.Atoi(num)
			ints[i] = append(ints[i], n)
		}
	}
	return ints
}

func extrapolate(cur [][]int) int {
	for i := len(cur) - 1; i > 0; i-- {
		cur[i-1] = append(cur[i-1], cur[i][len(cur[i])-1]+cur[i-1][len(cur[i-1])-1])
	}
	return cur[0][len(cur[0])-1]
}

func extrapolate_backwards(cur [][]int) int {
	for i := len(cur) - 1; i > 0; i-- {
		cur[i-1] = append([]int{cur[i-1][0] - cur[i][0]}, cur[i-1]...)
	}
	return cur[0][0]
}

func diff(ints []int, cur [][]int) [][]int {
	next := make([]int, 0)
	for i := 0; i < len(ints)-1; i++ {
		next = append(next, ints[i+1]-ints[i])
	}
	cur = append(cur, next)
	if !isAllZeroes(&next) {
		return diff(next, cur)
	}
	return cur
}
func isAllZeroes(slice *[]int) bool {
	for _, v := range *slice {
		if v != 0 {
			return false
		}
	}
	return true
}
