package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`

	scanner := bufio.NewScanner(strings.NewReader(input))
	var partOne int64
	for scanner.Scan() {
		line := scanner.Text()
		joltage := FindHighestJoltage(line)
		partOne += joltage
	}

	if partOne != 357 {
		t.Errorf("expected: 357\ngot: %d", partOne)
	}
}

func TestPartTwo(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`

	scanner := bufio.NewScanner(strings.NewReader(input))
	var partTwo int64
	for scanner.Scan() {
		line := scanner.Text()
		joltage := FindHighestJoltageTwelveBatteries(line)
		partTwo += joltage
	}

	if partTwo != 3121910778619 {
		t.Errorf("expected: 3121910778619\ngot: %d", partTwo)
	}
}
