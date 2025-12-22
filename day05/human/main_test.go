package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	ranges, ingredientIDs := ParseInput(scanner)

	if got := FreshIngredientIDs(ranges, ingredientIDs); got != 3 {
		t.Fatalf("expected: 3, got: %d", got)
	}
}

func TestPartTwo(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	ranges, _ := ParseInput(scanner)

	if got := FreshIngredientIDRangeCount(ranges); got != 14 {
		t.Fatalf("expected: 14, got: %d", got)
	}
}

func TestPartTwoDouble(t *testing.T) {
	input := `1-2
1-2
3-4`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	ranges, _ := ParseInput(scanner)

	if got := FreshIngredientIDRangeCount(ranges); got != 4 {
		t.Fatalf("expected: 4, got: %d", got)
	}
}
