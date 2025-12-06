package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalJoltage := 0

	for scanner.Scan() {
		line := scanner.Text()
		maxJoltage := findMaxJoltage(line)
		totalJoltage += maxJoltage
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", totalJoltage)
}

// findMaxJoltage finds the maximum joltage possible from a bank of batteries
// by selecting exactly two batteries and forming a two-digit number
func findMaxJoltage(bank string) int {
	maxJoltage := 0

	// Try all pairs of positions (i, j) where i < j
	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			// Get the digit values
			digit1 := int(bank[i] - '0')
			digit2 := int(bank[j] - '0')

			// Form the two-digit number
			joltage := digit1*10 + digit2

			// Track the maximum
			if joltage > maxJoltage {
				maxJoltage = joltage
			}
		}
	}

	return maxJoltage
}
