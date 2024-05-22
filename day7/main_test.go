package main

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	//cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b

	// hand1 > hand2
	tests := [][]Hand{
		{Hand{cards: "AAA22"}, Hand{cards: "A22AA"}},
		{Hand{cards: "22K34"}, Hand{cards: "AKQ34"}},
	}

	for _, tt := range tests {
		hand1 := tt[0].setup()
		hand2 := tt[1].setup()

		testname := fmt.Sprintf("%s > %s", hand1.cards, hand2.cards)
		t.Run(testname, func(t *testing.T) {
			if hand1.rank <= hand2.rank {
				t.Errorf("%s <= %s", hand1.cards, hand2.cards)
			}
		})
	}
}
