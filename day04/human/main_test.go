package main

import "testing"

func TestPartOne(t *testing.T) {
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

	fp := NewFloorPlan()

	err := fp.ParseInput(input)
	if err != nil {
		t.Fatalf("failed to parse input: %v", err)
	}

	if got := fp.PartOne(); got != 13 {
		t.Fatalf("expected 13: got %d", got)
	}
}

func TestPartTwo(t *testing.T) {
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

	fp := NewFloorPlan()

	err := fp.ParseInput(input)
	if err != nil {
		t.Fatalf("failed to parse input: %v", err)
	}

	if got := fp.PartTwo(0); got != 43 {
		t.Fatalf("expected0 43: got %d", got)
	}
}
