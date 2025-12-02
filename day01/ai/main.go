package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Rotation represents a single dial rotation instruction
type Rotation struct {
	Direction byte // 'L' or 'R'
	Distance  int
}

// parseRotation parses a line like "L68" or "R48" into a Rotation
func parseRotation(line string) (Rotation, error) {
	if len(line) < 2 {
		return Rotation{}, fmt.Errorf("invalid rotation: %s", line)
	}

	direction := line[0]
	distance, err := strconv.Atoi(line[1:])
	if err != nil {
		return Rotation{}, fmt.Errorf("invalid distance in %s: %v", line, err)
	}

	return Rotation{Direction: direction, Distance: distance}, nil
}

// applyRotation applies a rotation to the current dial position
// The dial has positions 0-99 and wraps around
func applyRotation(current int, rotation Rotation) int {
	const dialSize = 100

	var newPosition int
	if rotation.Direction == 'L' {
		// Rotate left (toward lower numbers)
		newPosition = current - rotation.Distance
	} else {
		// Rotate right (toward higher numbers)
		newPosition = current + rotation.Distance
	}

	// Handle wraparound using modulo
	// Go's modulo can return negative values, so we add dialSize before taking modulo
	newPosition = ((newPosition % dialSize) + dialSize) % dialSize

	return newPosition
}

// countZeroLandings counts how many times the dial lands on 0
func countZeroLandings(scanner *bufio.Scanner) int {
	const startPosition = 50

	position := startPosition
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		rotation, err := parseRotation(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing rotation: %v\n", err)
			continue
		}

		position = applyRotation(position, rotation)

		if position == 0 {
			count++
		}
	}

	return count
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	partOneAnswer := countZeroLandings(scanner)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part One: %d\n", partOneAnswer)
	fmt.Printf("Part Two: %d\n", 0) // Placeholder for Part Two
}
