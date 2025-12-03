package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Range represents a product ID range
type Range struct {
	Start int64
	End   int64
}

// isInvalidPartOne checks if a product ID is invalid (made of a sequence repeated exactly twice)
func isInvalidPartOne(id int64) bool {
	s := strconv.FormatInt(id, 10)

	// Must have even length to split in half
	if len(s)%2 != 0 {
		return false
	}

	// No leading zeroes allowed
	if s[0] == '0' {
		return false
	}

	// Split in half and check if both halves are equal
	mid := len(s) / 2
	firstHalf := s[:mid]
	secondHalf := s[mid:]

	return firstHalf == secondHalf
}

// isInvalidPartTwo checks if a product ID is invalid (made of a sequence repeated at least twice)
func isInvalidPartTwo(id int64) bool {
	s := strconv.FormatInt(id, 10)

	// No leading zeroes allowed
	if s[0] == '0' {
		return false
	}

	// Try all possible pattern lengths from 1 to len(s)/2
	for patternLen := 1; patternLen <= len(s)/2; patternLen++ {
		// Check if the string length is divisible by the pattern length
		if len(s)%patternLen != 0 {
			continue
		}

		// Extract the pattern
		pattern := s[:patternLen]

		// Check if the entire string consists of this pattern repeated
		valid := true
		for i := 0; i < len(s); i += patternLen {
			if s[i:i+patternLen] != pattern {
				valid = false
				break
			}
		}

		// If we found a valid repeating pattern, it's invalid
		if valid {
			return true
		}
	}

	return false
}

// parseRanges parses the input line into a slice of ranges
func parseRanges(line string) ([]Range, error) {
	line = strings.TrimSpace(line)
	parts := strings.Split(line, ",")

	ranges := make([]Range, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		rangeParts := strings.Split(part, "-")
		if len(rangeParts) != 2 {
			return nil, fmt.Errorf("invalid range format: %s", part)
		}

		start, err := strconv.ParseInt(rangeParts[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid start value: %s", rangeParts[0])
		}

		end, err := strconv.ParseInt(rangeParts[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid end value: %s", rangeParts[1])
		}

		ranges = append(ranges, Range{Start: start, End: end})
	}

	return ranges, nil
}

// sumInvalidIDsPartOne finds all invalid IDs (Part One rules) in the given ranges and returns their sum
func sumInvalidIDsPartOne(ranges []Range) int64 {
	var sum int64

	for _, r := range ranges {
		for id := r.Start; id <= r.End; id++ {
			if isInvalidPartOne(id) {
				sum += id
			}
		}
	}

	return sum
}

// sumInvalidIDsPartTwo finds all invalid IDs (Part Two rules) in the given ranges and returns their sum
func sumInvalidIDsPartTwo(ranges []Range) int64 {
	var sum int64

	for _, r := range ranges {
		for id := r.Start; id <= r.End; id++ {
			if isInvalidPartTwo(id) {
				sum += id
			}
		}
	}

	return sum
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Fprintf(os.Stderr, "Error reading input\n")
		os.Exit(1)
	}

	line := scanner.Text()
	ranges, err := parseRanges(line)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing ranges: %v\n", err)
		os.Exit(1)
	}

	partOneAnswer := sumInvalidIDsPartOne(ranges)
	fmt.Printf("Part One: %d\n", partOneAnswer)

	partTwoAnswer := sumInvalidIDsPartTwo(ranges)
	fmt.Printf("Part Two: %d\n", partTwoAnswer)
}
