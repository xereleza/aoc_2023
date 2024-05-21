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
	//fmt.Println(contents)
	seeds := get_seeds(contents[0])
	res := math.MaxInt32
	var maps []Maps
	for _, strmap := range contents[1:] {
		_map := get_next_map(&strmap)
		//fmt.Println(_map)
		maps = append(maps, _map.compute())
	}
	fmt.Println(maps)

	for i := 0; i < len(seeds); i += 2 {

		_seed, _ := strconv.Atoi(seeds[i])
		intRange, _ := strconv.Atoi(seeds[i+1])

		fmt.Println(_seed)
		fmt.Println(intRange)

		seed := &Interval{_seed, _seed + intRange - 1}
		mapped, not_mapped := make([]Interval, 0), make([]Interval, 0)

		for _, _maps := range maps {
			for _, _map := range _maps.maps {

				_mapped, _not_mapped := intervalSplit(seed, _map.source, _map.source+_map.length-1, _map.destination-_map.source)

				mapped = append(mapped, _mapped...)
				not_mapped = append(not_mapped, _not_mapped...)
			}
			fmt.Println(mapped)
			fmt.Println(not_mapped)
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

func get_next_map(contents *string) Maps {
	_map := make([]Range, 0)
	*contents = strings.Trim(*contents, "\n\n")
	entries := strings.Split(strings.Split(*contents, "map:\n")[1], "\n")
	for _, coords := range entries {
		split := strings.Split(coords, " ")
		_map = append(_map, Range{split[0], split[1], split[2]})
	}
	return Maps{_map, make([]Map, 0)}
}

type Interval struct {
	left  int
	right int
}

func (interval *Interval) offset(value int) {
	interval.left += value
	interval.right += value
}

func intervalSplit(interval *Interval, mapLeft int, mapRight int, offset int) ([]Interval, []Interval) {
	// returns computed and non-computed
	// no intersection
	if mapRight < interval.left || interval.right < mapLeft {
		return []Interval{}, []Interval{*interval}
	}

	if mapLeft < interval.left {
		//ml--l--mr--r
		if mapRight < interval.right {
			computed := &Interval{interval.left, mapRight}
			computed.offset(offset)
			return []Interval{*computed}, []Interval{{mapRight + 1, interval.right}}
			//ml--l--r---mr
		}
		computed := &Interval{interval.left, interval.right}
		computed.offset(offset)
		return []Interval{*computed}, []Interval{}

	}
	//-l--ml--mr--r
	if mapRight < interval.right {
		computed := &Interval{mapLeft, mapRight}
		computed.offset(offset)
		return []Interval{*computed}, []Interval{{interval.left, mapLeft - 1}, {mapRight + 1, interval.right}}
	}
	//-l--ml--r--mr
	computed := &Interval{mapLeft, interval.right}
	computed.offset(offset)
	return []Interval{*computed}, []Interval{{interval.left, mapLeft - 1}, {mapRight + 1, interval.right}}

}
