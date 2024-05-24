package main

import (
	"fmt"
	"testing"
)

/* func TestJokerLogic(t *testing.T) {
	game := []*Hand{
		Hand{cards: "32T3K", bid: 765}.setup(),
		Hand{cards: "T55J5", bid: 684}.setup(),
		Hand{cards: "KK677", bid: 28}.setup(),
		Hand{cards: "KTJJT", bid: 220}.setup(),
		Hand{cards: "QQQJA", bid: 483}.setup(),
	}
} */

func TestRank(t *testing.T) {
	var hand *Hand

	// pair
	hand = Hand{cards: "96Q8J"}.setup()
	if hand.rank != 2 {
		t.Errorf("%d != 2", hand.rank)
	}

	// two-pair
	hand = Hand{cards: "2A2KA"}.setup()
	if hand.rank != 3 {
		t.Errorf("%d != 3", hand.rank)
	}

	// fullhouse
	hand = Hand{cards: "2A2JA"}.setup()
	if hand.rank != 5 {
		t.Errorf("%d != 5", hand.rank)
	}

	// three of a kind
	hands := []*Hand{
		Hand{cards: "A8TAJ"}.setup(),
		Hand{cards: "AJKAT"}.setup(),
		Hand{cards: "A266J"}.setup(),
		Hand{cards: "A26JJ"}.setup(),
	}

	t.Run("three of a kind", func(t *testing.T) {
		for _, hand := range hands {
			if hand.rank != 4 {
				t.Errorf("%d != 4 (%s)", hand.rank, hand.cards)
			}
		}
	})

	// four of a kind
	hands = []*Hand{
		Hand{cards: "KTJJT"}.setup(),
		Hand{cards: "QQQJA"}.setup(),
	}

	t.Run("four of a kind", func(t *testing.T) {
		for _, hand := range hands {
			if hand.rank != 6 {
				t.Errorf("%d != 6 (%s)", hand.rank, hand.cards)
			}
		}
	})

	// cinco malandro
	hand = Hand{cards: "JJJJJ"}.setup()
	if hand.rank != 7 {
		t.Errorf("%d != 7", hand.rank)
	}
}

func TestCompare(t *testing.T) {
	//cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b

	// hand1 > hand2
	tests := [][]Hand{
		{Hand{cards: "AAA22"}, Hand{cards: "A22AA"}},
		{Hand{cards: "22K34"}, Hand{cards: "AKQ34"}},
		{Hand{cards: "AAAAA"}, Hand{cards: "AAAA4"}},

		{Hand{cards: "KTJJT"}, Hand{cards: "QQQJA"}},
		{Hand{cards: "KTJJT"}, Hand{cards: "T55J5"}},
		{Hand{cards: "QQQJA"}, Hand{cards: "T55J5"}},
	}

	for _, tt := range tests {
		hand1 := tt[0].setup()
		hand2 := tt[1].setup()

		testname := fmt.Sprintf("%s > %s", hand1.cards, hand2.cards)
		t.Run(testname, func(t *testing.T) {
			if compare(hand1, hand2) <= 0 {
				t.Errorf("%s <= %s", hand1.cards, hand2.cards)
			}
		})
	}

	// hand1 == hand2
	tests = [][]Hand{
		{Hand{cards: "AAA22"}, Hand{cards: "AAA22"}},
		{Hand{cards: "22K34"}, Hand{cards: "22K34"}},
		{Hand{cards: "AAAAA"}, Hand{cards: "AAAAA"}},
	}

	for _, tt := range tests {
		hand1 := tt[0].setup()
		hand2 := tt[1].setup()

		testname := fmt.Sprintf("%s = %s", hand1.cards, hand2.cards)
		t.Run(testname, func(t *testing.T) {
			if compare(hand1, hand2) != 0 {
				t.Errorf("%s != %s", hand1.cards, hand2.cards)
			}
		})
	}
}
