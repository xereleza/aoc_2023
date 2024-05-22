package main

import (
	"fmt"
	"log"
	"os"
)

type Hand struct {
	cards string
	rank  int
}

func (h *Hand) setup() *Hand {
	h.rank = 2
	return h
}

func main() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	hand := &Hand{cards: "23456"}
	fmt.Printf("%v\n", hand.rank)
	fmt.Println(hand.rank)
	hand.setup()
	fmt.Println(hand.rank)
	//contents := strings.Split(string(buf), "\n")
}

func compare(a string, b string) int {
	//cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b
	return 0
}
