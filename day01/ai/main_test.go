package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestPartOneSample(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	scanner := bufio.NewScanner(strings.NewReader(input))
	result := countZeroLandings(scanner)

	expected := 3
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestApplyRotation(t *testing.T) {
	tests := []struct {
		name     string
		position int
		rotation Rotation
		expected int
	}{
		{"Start L68", 50, Rotation{'L', 68}, 82},
		{"82 L30", 82, Rotation{'L', 30}, 52},
		{"52 R48", 52, Rotation{'R', 48}, 0},
		{"0 L5", 0, Rotation{'L', 5}, 95},
		{"95 R60", 95, Rotation{'R', 60}, 55},
		{"55 L55", 55, Rotation{'L', 55}, 0},
		{"0 L1", 0, Rotation{'L', 1}, 99},
		{"99 L99", 99, Rotation{'L', 99}, 0},
		{"0 R14", 0, Rotation{'R', 14}, 14},
		{"14 L82", 14, Rotation{'L', 82}, 32},
		{"Simple wrap right", 99, Rotation{'R', 1}, 0},
		{"Simple wrap left", 0, Rotation{'L', 1}, 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := applyRotation(tt.position, tt.rotation)
			if result != tt.expected {
				t.Errorf("applyRotation(%d, %c%d) = %d, want %d",
					tt.position, tt.rotation.Direction, tt.rotation.Distance, result, tt.expected)
			}
		})
	}
}
