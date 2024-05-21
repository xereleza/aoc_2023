package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func main() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	contents := strings.Split(string(buf), "\n")
	part_1(contents)
	part_2(contents)

}

func part_2(contents []string) {
	time := parse_2(contents[0])
	distance := parse_2(contents[1])

	fmt.Println("Part 2:", measure(time, distance))
}

func part_1(contents []string) {
	times := parse_1(contents[0])
	distances := parse_1(contents[1])
	res := int64(1)
	for i, time := range times {
		res *= measure(time, distances[i])
	}
	fmt.Println("Part 1:", res)
}

func measure(time int, distance int) int64 {
	lower, upper := quadratic_roots(-1, time, -distance)
	if lower.IsInteger() {
		lower = lower.Add(decimal.NewFromInt(1))
	} else {
		lower = lower.Ceil()
	}
	if upper.IsInteger() {
		upper = upper.Sub(decimal.NewFromInt(1))
	} else {
		upper = upper.Truncate(0)
	}
	return upper.IntPart() - lower.IntPart() + 1
}

func quadratic_roots(a int, b int, c int) (decimal.Decimal, decimal.Decimal) {
	d := float64(b*b - 4*a*c)
	return decimal.NewFromFloat((-float64(b) + math.Sqrt(d)) / float64(2*a)), decimal.NewFromFloat((-float64(b) - math.Sqrt(d)) / float64(2*a))
}

func parse_1(contents string) []int {
	parsed := make([]int, 0)
	for _, _str := range strings.FieldsFunc(contents, func(r rune) bool { return r == ' ' })[1:] {
		_int, _ := strconv.Atoi(_str)
		parsed = append(parsed, _int)
	}
	return parsed
}

func parse_2(contents string) int {
	parsed, _ := strconv.Atoi(strings.Join(strings.FieldsFunc(contents, func(r rune) bool { return r == ' ' })[1:], ""))
	return parsed
}
