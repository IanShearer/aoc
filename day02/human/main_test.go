package main

import "testing"

func TestPartOne(t *testing.T) {
	ranges := []IDRange{
		{Min: 11, Max: 22, PartOneInvalidIDs: make([]uint, 0)},
		{Min: 95, Max: 115, PartOneInvalidIDs: make([]uint, 0)},
		{Min: 998, Max: 1012, PartOneInvalidIDs: make([]uint, 0)},
		{Min: 1188511880, Max: 1188511890, PartOneInvalidIDs: make([]uint, 0)},
		{Min: 222220, Max: 222224, PartOneInvalidIDs: make([]uint, 0)},
		{Min: 1698522, Max: 1698528, PartOneInvalidIDs: make([]uint, 0)},
		{Min: 446443, Max: 446449, PartOneInvalidIDs: make([]uint, 0)},
		{Min: 38593856, Max: 38593862, PartOneInvalidIDs: make([]uint, 0)},
		{Min: 565653, Max: 565659, PartOneInvalidIDs: make([]uint, 0)},
		{Min: 824824821, Max: 824824827, PartOneInvalidIDs: make([]uint, 0)},
		{Min: 2121212118, Max: 2121212124, PartOneInvalidIDs: make([]uint, 0)},
	}

	var sum uint
	for _, r := range ranges {
		r.FindInvalidIDs()
		for _, i := range r.PartOneInvalidIDs {
			sum += i
		}
	}

	if sum != 1227775554 {
		t.Errorf("expected: 1227775554, got: %d", sum)
	}
}

func TestPartTwo(t *testing.T) {
	ranges := []IDRange{
		{Min: 11, Max: 22, PartOneInvalidIDs: make([]uint, 0), PartTwoInvalidIDs: make([]uint, 0)},
		{Min: 95, Max: 115, PartOneInvalidIDs: make([]uint, 0), PartTwoInvalidIDs: make([]uint, 0)},
		{Min: 998, Max: 1012, PartOneInvalidIDs: make([]uint, 0), PartTwoInvalidIDs: make([]uint, 0)},
		{Min: 1188511880, Max: 1188511890, PartOneInvalidIDs: make([]uint, 0), PartTwoInvalidIDs: make([]uint, 0)},
		{Min: 222220, Max: 222224, PartOneInvalidIDs: make([]uint, 0), PartTwoInvalidIDs: make([]uint, 0)},
		{Min: 1698522, Max: 1698528, PartOneInvalidIDs: make([]uint, 0), PartTwoInvalidIDs: make([]uint, 0)},
		{Min: 446443, Max: 446449, PartOneInvalidIDs: make([]uint, 0), PartTwoInvalidIDs: make([]uint, 0)},
		{Min: 38593856, Max: 38593862, PartOneInvalidIDs: make([]uint, 0), PartTwoInvalidIDs: make([]uint, 0)},
		{Min: 565653, Max: 565659, PartOneInvalidIDs: make([]uint, 0), PartTwoInvalidIDs: make([]uint, 0)},
		{Min: 824824821, Max: 824824827, PartOneInvalidIDs: make([]uint, 0), PartTwoInvalidIDs: make([]uint, 0)},
		{Min: 2121212118, Max: 2121212124, PartOneInvalidIDs: make([]uint, 0), PartTwoInvalidIDs: make([]uint, 0)},
	}

	var sum uint
	for _, r := range ranges {
		r.FindInvalidIDs()
		for _, i := range r.PartTwoInvalidIDs {
			sum += i
		}
	}

	if sum != 4174379265 {
		t.Errorf("expected: 4174379265, got: %d", sum)
	}
}
