package main

import (
	"fmt"
	"log"
	"math"
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
	buf := make([]byte, info.Size())
	file.Read(buf)
	contents := strings.Split(string(buf), "\n\n")
	print(contents)
	seeds := get_seeds(contents[0])
	res := math.MaxInt32
	var maps []Maps
	for _, strmap := range contents[1:] {
		maps = append(maps, get_next_map(strmap).compute())
	}
	for _, seed := range seeds {
		seed, err := strconv.Atoi(seed)
		if err != nil {
			log.Fatal(err)
		}
		for _, _maps := range maps {
			for _, _map := range _maps.maps {
				if _map.source <= seed && seed <= _map.source+_map.length {
					seed += _map.destination - _map.source
					break
				}
			}
		}
		if seed < res {
			res = seed
		}
	}
	fmt.Println("Saída final do programa ~~~uhul")
	fmt.Println(res)
}

func get_seeds(seeds_detail string) []string {
	return strings.Split(seeds_detail, " ")[1:]

}

type Range struct {
	destination string
	source      string
	length      string
}

type Maps struct {
	entries []Range
	maps    []Map
}

type Map struct {
	length      int
	source      int
	destination int
}

func (maps Maps) compute() Maps {
	// [length, source, destination]
	var m Map
	for _, _range := range maps.entries {
		m.length, _ = strconv.Atoi(_range.length)
		m.source, _ = strconv.Atoi(_range.source)
		m.destination, _ = strconv.Atoi(_range.destination)
		maps.maps = append(maps.maps, m)
	}
	return maps
}

func get_next_map(contents string) Maps {
	_map := make([]Range, 0)
	contents = strings.Trim(contents, "\n\n")
	fmt.Println(contents)
	fmt.Println("<<--------->")
	entries := strings.Split(strings.Split(contents, "map:\n")[1], "\n")
	for _, coords := range entries {
		split := strings.Split(coords, " ")
		_map = append(_map, Range{split[0], split[1], split[2]})
	}
	return Maps{_map, make([]Map, 0)}
}
