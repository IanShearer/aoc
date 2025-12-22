package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	columns := ParseInput(scanner)
	partOne := 0
	for _, c := range columns {
		p1 := c.PartOne()
		partOne += p1
	}

	if partOne != 4277556 {
		t.Fatalf("expected 4277556, got %d", partOne)
	}
}

func TestPartTwo(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	columns := ParseInput(scanner)
	partTwo := 0
	for _, c := range columns {
		p2 := c.PartTwo()
		partTwo += p2
	}

	if partTwo != 3263827 {
		t.Fatalf("expected 3263827, got %d", partTwo)
	}
}
