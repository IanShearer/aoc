package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestPartOneSample(t *testing.T) {
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

	scanner := bufio.NewScanner(strings.NewReader(input))
	partOne, _ := solve(scanner)

	expected := 3
	if partOne != expected {
		t.Errorf("Expected %d, got %d", expected, partOne)
	}
}

func TestPartTwoSample(t *testing.T) {
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

	scanner := bufio.NewScanner(strings.NewReader(input))
	_, partTwo := solve(scanner)

	expected := 14
	if partTwo != expected {
		t.Errorf("Expected %d, got %d", expected, partTwo)
	}
}
