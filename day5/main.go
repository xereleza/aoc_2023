package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	fmt.Println(info.Size())
	buf := make([]byte, info.Size())
	file.Read(buf)
	contents := strings.Split(string(buf), "\n\n")
	//fmt.Println(contents)
	seeds := get_seeds(contents[0])
	fmt.Println("seeds:", seeds)
	for i := 1; i < len(contents); i++ {
		_map := get_next_map(contents[i:])
		fmt.Println("map:", _map)
		fmt.Println("compute:", _map.compute())
	}
}

func get_seeds(seeds_detail string) []string {
	return strings.Split(seeds_detail, " ")[1:]

}

type Range struct {
	destination string
	source      string
	length      string
}

type Map struct {
	entries []Range
}

func (m Map) compute() map[int]int {
	_map := make(map[int]int)
	for _, _range := range m.entries {
		length, _ := strconv.Atoi(_range.length)
		source, _ := strconv.Atoi(_range.source)
		destination, _ := strconv.Atoi(_range.destination)
		_map[source] = destination
		for i := 1; i < length; i++ {
			_map[source+i] = destination + i
		}
	}
	return _map
}

func get_next_map(contents []string) Map {
	_map := make([]Range, 0)
	entries := strings.Split(strings.Split(contents[0], "map:\n")[1], "\n")
	for _, coords := range entries {
		split := strings.Split(coords, " ")
		_map = append(_map, Range{split[0], split[1], split[2]})
	}
	return Map{_map}
}
