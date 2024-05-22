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

	seeds := getSeeds(contents[0])
	maps := parseMaps(contents[1:])
	res := processSeeds(seeds, maps)

	fmt.Println("Final Result:", res)
}

func getSeeds(seedData string) []Interval {
	seedStrings := strings.Split(strings.Split(seedData, ": ")[1], " ")
	seeds := make([]Interval, 0)
	for i := 0; i < len(seedStrings); i += 2 {
		if i+1 >= len(seedStrings) {
			log.Fatalf("Invalid seed data: expected pairs of values but found an odd number.")
		}
		seed, err1 := strconv.Atoi(seedStrings[i])
		seedRange, err2 := strconv.Atoi(seedStrings[i+1])
		if err1 != nil || err2 != nil {
			log.Fatalf("Invalid seed data: expected integer values but found %v and %v.", seedStrings[i], seedStrings[i+1])
		}
		seeds = append(seeds, Interval{seed, seed + seedRange - 1})
	}
	return seeds
}

func parseMaps(mapData []string) []Map {
	var maps []Map
	for _, mapStr := range mapData {
		parsedMap, err := parseMap(mapStr)
		if err != nil {
			log.Fatalf("Error parsing map: %v", err)
		}
		maps = append(maps, parsedMap)
	}
	return maps
}

func parseMap(mapStr string) (Map, error) {
	lines := strings.Split(strings.Trim(mapStr, "\n"), "\n")
	mappings := make([]Mapping, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 3 {
			// Skip lines that do not contain exactly three fields
			continue
		}
		destination, err1 := strconv.Atoi(fields[0])
		source, err2 := strconv.Atoi(fields[1])
		length, err3 := strconv.Atoi(fields[2])
		if err1 != nil || err2 != nil || err3 != nil {
			return Map{}, fmt.Errorf("invalid integer values in mapping: %v", line)
		}
		mappings = append(mappings, Mapping{length, source, destination})
	}
	return Map{mappings}, nil
}

func processSeeds(seeds []Interval, maps []Map) int {
	minValue := math.MaxInt32
	for _, interval := range seeds {
		mappedIntervals := []Interval{interval}
		for _, m := range maps {
			mappedIntervals = applyMappings(m, mappedIntervals)
		}

		for _, inter := range mappedIntervals {
			if inter.left < minValue {
				minValue = inter.left
			}
		}
	}
	return minValue
}

type Map struct {
	mappings []Mapping
}

type Mapping struct {
	length      int
	source      int
	destination int
}

type Interval struct {
	left  int
	right int
}

func (interval *Interval) offset(value int) {
	interval.left += value
	interval.right += value
}

func applyMappings(m Map, intervals []Interval) []Interval {
	var resIntervals []Interval
	for _, mapping := range m.mappings {
		var newIntervals []Interval
		for _, interval := range intervals {
			mapped, unmapped := splitAndMapInterval(interval, mapping)
			if len(mapped) > 0 {
				resIntervals = append(resIntervals, mapped...)
			}
			if len(unmapped) > 0 {
				newIntervals = append(newIntervals, unmapped...)
			}
		}
		intervals = newIntervals
	}
	resIntervals = append(resIntervals, intervals...)
	return resIntervals
}

func splitAndMapInterval(interval Interval, mapping Mapping) ([]Interval, []Interval) {
	// returns computed and non-computed
	// no intersection
	if mapping.source+mapping.length-1 < interval.left || interval.right < mapping.source {
		return []Interval{}, []Interval{interval}
	}

	if mapping.source <= interval.left {
		//ml--l--mr--r
		if mapping.source+mapping.length-1 < interval.right {
			computed := &Interval{interval.left, mapping.source + mapping.length - 1}
			computed.offset(mapping.destination - mapping.source)
			return []Interval{*computed}, []Interval{{mapping.source + mapping.length - 1 + 1, interval.right}}
		}
		//ml--l--r---mr
		computed := &Interval{interval.left, interval.right}
		computed.offset(mapping.destination - mapping.source)
		return []Interval{*computed}, []Interval{}

	}
	//-l--ml--mr--r
	if mapping.source+mapping.length-1 < interval.right {
		computed := &Interval{mapping.source, mapping.source + mapping.length - 1}
		computed.offset(mapping.destination - mapping.source)
		return []Interval{*computed}, []Interval{{interval.left, mapping.source - 1}, {mapping.source + mapping.length - 1 + 1, interval.right}}
	}
	//-l--ml--r--mr
	computed := &Interval{mapping.source, interval.right}
	computed.offset(mapping.destination - mapping.source)
	return []Interval{*computed}, []Interval{{interval.left, mapping.source - 1}}

}
