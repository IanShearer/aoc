package main

import (
	"strings"
	"testing"
)

func TestPartOneSample(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	grid := strings.Split(input, "\n")
	result := countAccessibleRolls(grid)
	expected := 13

	if result != expected {
		t.Errorf("Expected %d accessible rolls, got %d", expected, result)
	}
}

func TestPartTwoSample(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	grid := strings.Split(input, "\n")
	result := countRemovableRolls(grid)
	expected := 43

	if result != expected {
		t.Errorf("Expected %d removable rolls, got %d", expected, result)
	}
}
