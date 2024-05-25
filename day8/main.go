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

func count_steps2(instructions string, _map map[string]*Tree) int {
	steps := 0
	current := make([]string, 0)
	for k := range _map {
		if k[2] == 'A' {
			current = append(current, k)
		}
	}
	finished := false
	for !finished {
		for _, direction := range instructions {
			if finished {
				break
			}
			fmt.Println("before:", current)
			walk(_map, &current, &direction)
			fmt.Println("after:", current)
			finished = check(&current)
			steps += 1
		}
	}
	return steps
}

func check(current *[]string) bool {
	for _, name := range *current {
		if name[2] != 'Z' {
			return false
		}
	}
	return true
}

func walk(_map map[string]*Tree, current *[]string, direction *rune) {
	next := make([]string, 0)
	for _, name := range *current {
		next = append(next, _map[name].walk(*direction).val)
	}
	*current = next
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
