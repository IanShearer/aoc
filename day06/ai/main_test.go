package main

import (
	"strings"
	"testing"
)

func TestPartOneSample(t *testing.T) {
	input := `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `

	lines := strings.Split(input, "\n")
	result := solvePartOne(lines)
	expected := 4277556 // 33210 + 490 + 4243455 + 401

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPartTwoSample(t *testing.T) {
	input := `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `

	lines := strings.Split(input, "\n")
	result := solvePartTwo(lines)
	expected := 3263827 // 1058 + 3253600 + 625 + 8544

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
