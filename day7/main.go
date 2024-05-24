package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards   string
	rank    int
	counter map[string]int
	bid     int
}

func (h Hand) setup() *Hand {
	h.counter = make(map[string]int)
	for _, card := range h.cards {
		k := string(card)
		i := h.counter[k]
		h.counter[k] = i + 1
	}
	max := 1
	jokers := 0
	var ref string
	for k, v := range h.counter {
		if k == "J" {
			jokers = v
		} else if v > max {
			max = v
			ref = k
		}
	}
	if jokers == 5 {
		max = 5
	} else {
		max += jokers
	}
	switch max {
	case 1:
		h.rank = 1
	case 2:
		h.rank = checkTwoPair(h.counter, &ref)
	case 3:
		h.rank = checkFullHouse(h.counter, &ref)
	default:
		h.rank = max + 2
	}

	return &h
}

func checkTwoPair(counter map[string]int, ref *string) int {
	for k, v := range counter {
		if k == *ref || k == "J" {
			continue
		}
		if v == 2 {
			return 3
		}
	}
	return 2
}

func checkFullHouse(counter map[string]int, ref *string) int {
	for k, v := range counter {
		if k == *ref || k == "J" {
			continue
		}
		if v == 2 {
			return 5
		}
	}
	return 4
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
	print(compute_game(parse_game(strings.Split(string(buf), "\n"))))
}

func compute_game(game []*Hand) int {
	res := 0
	slices.SortFunc(game, compare)
	for i, hand := range game {
		res += hand.bid * (i + 1)
	}
	return res
}

func parse_game(contents []string) []*Hand {
	game := make([]*Hand, 0)
	for _, hand := range contents {
		_detail := strings.Split(hand, " ")
		_bid, err := strconv.Atoi(_detail[1])
		if err != nil {
			log.Fatal(err)
		}
		game = append(game, Hand{cards: _detail[0], bid: _bid}.setup())
	}
	return game
}

func compare(a *Hand, b *Hand) int {
	//cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b
	if a.rank == b.rank {
		return tiebreaker(a.cards, b.cards)
	}
	return a.rank - b.rank
}

func tiebreaker(a string, b string) int {
	card_ranks := map[byte]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}
	for i := range 5 {
		if a[i] == b[i] {
			continue
		}
		return card_ranks[a[i]] - card_ranks[b[i]]
	}
	return 0
}
