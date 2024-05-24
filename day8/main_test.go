package main

import (
	"fmt"
	"testing"
)

func TestExamples(t *testing.T) {
	for i, expected := range []int{2, 6} {
		t.Run(fmt.Sprintf("example %d", i+1), func(t *testing.T) {
			instructions, _map := load_map(fmt.Sprintf("test%d.txt", i+1))
			result := count_steps(instructions, _map)
			if result != expected {
				t.Errorf("count result: %d, expected: %d", result, expected)
			}
		})
	}
}
