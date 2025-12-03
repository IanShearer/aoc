package main

import (
	"testing"
)

func TestIsInvalidPartOne(t *testing.T) {
	tests := []struct {
		id      int64
		invalid bool
	}{
		{11, true},      // 1 repeated twice
		{22, true},      // 2 repeated twice
		{99, true},      // 9 repeated twice
		{1010, true},    // 10 repeated twice
		{6464, true},    // 64 repeated twice
		{123123, true},  // 123 repeated twice
		{222222, true},  // 222 repeated twice
		{446446, true},  // 446 repeated twice
		{38593859, true},    // 3859 repeated twice (from example)
		{1188511885, true},  // 118851 repeated twice (from example)
		{101, false},    // Not a repeated pattern
		{12, false},     // Different digits
		{123, false},    // Odd length
		{1234, false},   // Not repeated
		{111, false},    // Repeated 3 times (not exactly twice)
		{0101, false},   // Leading zeroes (though this would be parsed as 101)
	}

	for _, tt := range tests {
		got := isInvalidPartOne(tt.id)
		if got != tt.invalid {
			t.Errorf("isInvalidPartOne(%d) = %v, want %v", tt.id, got, tt.invalid)
		}
	}
}

func TestIsInvalidPartTwo(t *testing.T) {
	tests := []struct {
		id      int64
		invalid bool
	}{
		{11, true},          // 1 repeated twice
		{22, true},          // 2 repeated twice
		{99, true},          // 9 repeated twice
		{111, true},         // 1 repeated 3 times
		{999, true},         // 9 repeated 3 times
		{1010, true},        // 10 repeated twice
		{6464, true},        // 64 repeated twice
		{123123, true},      // 123 repeated twice
		{123123123, true},   // 123 repeated 3 times
		{222222, true},      // 222 repeated twice
		{446446, true},      // 446 repeated twice
		{38593859, true},    // 3859 repeated twice
		{1188511885, true},  // 118851 repeated twice
		{565656, true},      // 56 repeated 3 times
		{824824824, true},   // 824 repeated 3 times
		{2121212121, true},  // 21 repeated 5 times
		{1212121212, true},  // 12 repeated 5 times
		{12341234, true},    // 1234 repeated twice
		{1111111, true},     // 1 repeated 7 times
		{101, false},        // Not a repeated pattern
		{12, false},         // Different digits
		{123, false},        // Not repeated
		{1234, false},       // Not repeated
		{0101, false},       // Leading zeroes (though this would be parsed as 101)
	}

	for _, tt := range tests {
		got := isInvalidPartTwo(tt.id)
		if got != tt.invalid {
			t.Errorf("isInvalidPartTwo(%d) = %v, want %v", tt.id, got, tt.invalid)
		}
	}
}

func TestPartOneSample(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	ranges, err := parseRanges(input)
	if err != nil {
		t.Fatalf("Error parsing ranges: %v", err)
	}

	got := sumInvalidIDsPartOne(ranges)
	want := int64(1227775554)

	if got != want {
		t.Errorf("sumInvalidIDsPartOne() = %d, want %d", got, want)
	}
}

func TestPartTwoSample(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	ranges, err := parseRanges(input)
	if err != nil {
		t.Fatalf("Error parsing ranges: %v", err)
	}

	got := sumInvalidIDsPartTwo(ranges)
	want := int64(4174379265)

	if got != want {
		t.Errorf("sumInvalidIDsPartTwo() = %d, want %d", got, want)
	}
}

func TestParseRanges(t *testing.T) {
	input := "11-22,95-115"
	ranges, err := parseRanges(input)
	if err != nil {
		t.Fatalf("Error parsing ranges: %v", err)
	}

	if len(ranges) != 2 {
		t.Errorf("Expected 2 ranges, got %d", len(ranges))
	}

	if ranges[0].Start != 11 || ranges[0].End != 22 {
		t.Errorf("First range incorrect: got {%d, %d}, want {11, 22}", ranges[0].Start, ranges[0].End)
	}

	if ranges[1].Start != 95 || ranges[1].End != 115 {
		t.Errorf("Second range incorrect: got {%d, %d}, want {95, 115}", ranges[1].Start, ranges[1].End)
	}
}
