package main

import (
	"testing"
)

func TestPartOneSample(t *testing.T) {
	// Sample banks from the problem description
	testCases := []struct {
		bank     string
		expected int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	for _, tc := range testCases {
		result := findMaxJoltage(tc.bank)
		if result != tc.expected {
			t.Errorf("findMaxJoltage(%s) = %d, expected %d", tc.bank, result, tc.expected)
		}
	}

	// Total should be 357
	total := 98 + 89 + 78 + 92
	expectedTotal := 357
	if total != expectedTotal {
		t.Errorf("Total = %d, expected %d", total, expectedTotal)
	}
}
