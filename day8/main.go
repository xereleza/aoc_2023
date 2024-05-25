package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	instructions, _map := load_map("day8/input.txt")
	res := count_steps1(instructions, _map)
	fmt.Println("Part 1:", res)
	res = count_steps2(instructions, _map)
	fmt.Println("Part 2:", res)
}

func count_steps1(instructions string, _map map[string]*Tree) int {
	steps := 0
	for current := "AAA"; current != "ZZZ"; {
		for _, direction := range instructions {
			if current == "ZZZ" {
				break
			}
			current = _map[current].walk(direction).val
			steps += 1
		}
	}
	return steps
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func count_steps2(instructions string, _map map[string]*Tree) int {
	periods := make([]int, 0)
	for k := range _map {
		if k[2] != 'A' {
			continue
		}
		_steps := 0
		for current := k; current[2] != 'Z'; {
			for _, direction := range instructions {
				if current[2] == 'Z' {
					break
				}
				current = _map[current].walk(direction).val
				_steps += 1
			}
		}
		periods = append(periods, _steps)
	}

	return LCM(periods[0], periods[1], periods...)
}

func load_map(input_path string) (string, map[string]*Tree) {
	file, err := os.Open(input_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	contents := strings.Split(string(buf), "\n\n")
	instructions := contents[0]
	nodes := make(map[string]*Tree)
	values := make([]string, 0)
	leaves := make([][]string, 0)
	for _, node := range strings.Split(contents[1], "\n") {
		detail := strings.Split(node, " = ")
		v := detail[0]
		values = append(values, v)
		nodes[v] = &Tree{val: v}
		leaves = append(leaves, strings.Split(strings.TrimSuffix(strings.TrimPrefix(detail[1], "("), ")"), ", "))
	}
	for i, val := range values {
		nodes[val].left = nodes[leaves[i][0]]
		nodes[val].right = nodes[leaves[i][1]]
	}

	return instructions, nodes

}

type Tree struct {
	val   string
	left  *Tree
	right *Tree
}

func (t *Tree) walk(direction rune) *Tree {
	switch direction {
	case 'R':
		return t.right
	case 'L':
		return t.left
	default:
		return &Tree{}
	}
}
