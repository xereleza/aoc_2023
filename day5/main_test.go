package main

import (
	"testing"
)

func TestApplyMappings(t *testing.T) {
	// Test case with a single mapping
	interval := Interval{10, 20}
	mapping := Mapping{5, 10, 20}
	expected := []Interval{{15, 25}}
	if res := applyMappings(Map{[]Mapping{mapping}}, []Interval{interval}); !intervalsEqual(res, expected) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected, res)
	}

	// Test case with multiple mappings
	interval = Interval{30, 40}
	mappings := []Mapping{
		{5, 30, 40},
		{10, 40, 50},
	}
	expected = []Interval{{35, 45}, {40, 50}}
	if res := applyMappings(Map{mappings}, []Interval{interval}); !intervalsEqual(res, expected) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected, res)
	}
}

// Helper function to check if two slices of intervals are equal
func intervalsEqual(intervals1, intervals2 []Interval) bool {
	if len(intervals1) != len(intervals2) {
		return false
	}
	for i := range intervals1 {
		if intervals1[i] != intervals2[i] {
			return false
		}
	}
	return true
}
func TestProcessSeeds(t *testing.T) {
	// Test case with individual seed numbers
	seeds1 := []Interval{{79, 79}, {14, 14}, {55, 55}, {13, 13}}
	maps1 := []Map{
		{
			mappings: []Mapping{{2, 98, 50}, {48, 50, 52}},
		},
		{
			mappings: []Mapping{{3, 0, 49}, {2, 15, 53}},
		},
		{
			mappings: []Mapping{{3, 49, 88}, {2, 53, 18}},
		},
		{
			mappings: []Mapping{{3, 88, 45}, {3, 18, 77}},
		},
		{
			mappings: []Mapping{{2, 45, 0}, {3, 77, 69}},
		},
		{
			mappings: []Mapping{{3, 0, 60}, {3, 69, 56}},
		},
		{
			mappings: []Mapping{{3, 60, 37}, {3, 56, 93}},
		},
	}
	expected1 := 82
	if res := processSeeds(seeds1, maps1); res != expected1 {
		t.Errorf("Test case 1 failed. Expected %d, got %d", expected1, res)
	}

	// Test case with seed number ranges
	seeds2 := []Interval{{79, 14}, {55, 13}}
	expected2 := 46
	if res := processSeeds(seeds2, maps1); res != expected2 {
		t.Errorf("Test case 2 failed. Expected %d, got %d", expected2, res)
	}
}

func TestSplitAndMapInterval(t *testing.T) {
	// Test case where interval is completely outside mapping range
	interval := Interval{5, 10}
	mapping := Mapping{5, 15, 25}
	expectedMapped := []Interval{{5, 10}}
	expectedUnmapped := []Interval{}
	mapped, unmapped := splitAndMapInterval(interval, mapping)
	if !intervalsEqual(mapped, expectedMapped) || !intervalsEqual(unmapped, expectedUnmapped) {
		t.Errorf("Test case 1 failed. Expected mapped: %v, unmapped: %v, got mapped: %v, unmapped: %v", expectedMapped, expectedUnmapped, mapped, unmapped)
	}

	// Test case where interval is completely inside mapping range
	interval = Interval{20, 25}
	expectedMapped = []Interval{{35, 40}}
	expectedUnmapped = []Interval{}
	mapped, unmapped = splitAndMapInterval(interval, mapping)
	if !intervalsEqual(mapped, expectedMapped) || !intervalsEqual(unmapped, expectedUnmapped) {
		t.Errorf("Test case 2 failed. Expected mapped: %v, unmapped: %v, got mapped: %v, unmapped: %v", expectedMapped, expectedUnmapped, mapped, unmapped)
	}

	// Test case where interval partially overlaps with mapping range
	interval = Interval{10, 20}
	expectedMapped = []Interval{{20, 30}}
	expectedUnmapped = []Interval{{10, 15}}
	mapped, unmapped = splitAndMapInterval(interval, mapping)
	if !intervalsEqual(mapped, expectedMapped) || !intervalsEqual(unmapped, expectedUnmapped) {
		t.Errorf("Test case 3 failed. Expected mapped: %v, unmapped: %v, got mapped: %v, unmapped: %v", expectedMapped, expectedUnmapped, mapped, unmapped)
	}
}
