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

// countZeroClicksInRotation counts how many times the dial clicks through 0
// during a rotation (not just at the end)
func countZeroClicksInRotation(current int, rotation Rotation) int {
	const dialSize = 100
	distance := rotation.Distance

	if rotation.Direction == 'L' {
		// Rotating left (subtracting)
		// We cross 0 when going from 0 to 99
		if current == 0 {
			// Special case: starting at 0, we cross it every 100 clicks
			return distance / dialSize
		} else if distance >= current {
			// We'll cross 0 at least once
			// First crossing at 'current' clicks, then every 100 clicks after
			return (distance-current)/dialSize + 1
		} else {
			// Not enough distance to reach 0
			return 0
		}
	} else {
		// Rotating right (adding)
		// We cross 0 when going from 99 to 0
		if distance >= dialSize-current {
			// First crossing at (100-current) clicks, then every 100 clicks after
			return (distance+current-dialSize)/dialSize + 1
		} else {
			// Not enough distance to reach 0
			return 0
		}
	}
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

// countAllZeroClicks counts every time the dial clicks through 0
// including during rotations, not just at the end
func countAllZeroClicks(scanner *bufio.Scanner) int {
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

		// Count how many times we click through 0 during this rotation
		count += countZeroClicksInRotation(position, rotation)

		// Update position for next rotation
		position = applyRotation(position, rotation)
	}

	return count
}

func main() {
	// Part One
	file, err := os.Open("../input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	partOneAnswer := countZeroLandings(scanner)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
	file.Close()

	// Part Two
	file, err = os.Open("../input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	partTwoAnswer := countAllZeroClicks(scanner)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part One: %d\n", partOneAnswer)
	fmt.Printf("Part Two: %d\n", partTwoAnswer)
}
