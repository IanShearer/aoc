package main

import (
	"math/big"
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

func TestPartTwoSample(t *testing.T) {
	// Sample banks from the problem description for Part Two
	testCases := []struct {
		bank     string
		expected string
	}{
		{"987654321111111", "987654321111"},
		{"811111111111119", "811111111119"},
		{"234234234234278", "434234234278"},
		{"818181911112111", "888911112111"},
	}

	for _, tc := range testCases {
		result := findMaxKDigits(tc.bank, 12)
		if result != tc.expected {
			t.Errorf("findMaxKDigits(%s, 12) = %s, expected %s", tc.bank, result, tc.expected)
		}
	}

	// Total should be 3121910778619
	total := big.NewInt(0)
	nums := []string{"987654321111", "811111111119", "434234234278", "888911112111"}
	for _, numStr := range nums {
		num := new(big.Int)
		num.SetString(numStr, 10)
		total.Add(total, num)
	}

	expectedTotal := "3121910778619"
	if total.String() != expectedTotal {
		t.Errorf("Total = %s, expected %s", total.String(), expectedTotal)
	}
}
