package main

import (
	"testing"
)

func TestIsInvalid(t *testing.T) {
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
		{0101, false},   // Leading zeroes (though this would be parsed as 101)
	}

	for _, tt := range tests {
		got := isInvalid(tt.id)
		if got != tt.invalid {
			t.Errorf("isInvalid(%d) = %v, want %v", tt.id, got, tt.invalid)
		}
	}
}

func TestPartOneSample(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	ranges, err := parseRanges(input)
	if err != nil {
		t.Fatalf("Error parsing ranges: %v", err)
	}

	got := sumInvalidIDs(ranges)
	want := int64(1227775554)

	if got != want {
		t.Errorf("sumInvalidIDs() = %d, want %d", got, want)
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
