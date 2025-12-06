package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalJoltagePartOne := 0
	totalJoltagePartTwo := big.NewInt(0)

	for scanner.Scan() {
		line := scanner.Text()

		// Part One: select 2 batteries
		maxJoltage := findMaxJoltage(line)
		totalJoltagePartOne += maxJoltage

		// Part Two: select 12 batteries
		maxJoltageStr := findMaxKDigits(line, 12)
		maxJoltageBig := new(big.Int)
		maxJoltageBig.SetString(maxJoltageStr, 10)
		totalJoltagePartTwo.Add(totalJoltagePartTwo, maxJoltageBig)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", totalJoltagePartOne)
	fmt.Printf("Part Two: %s\n", totalJoltagePartTwo.String())
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

// findMaxKDigits finds the largest k-digit subsequence from a bank of batteries
// using a greedy algorithm that maintains order
func findMaxKDigits(bank string, k int) string {
	n := len(bank)
	if k > n {
		return bank
	}

	result := make([]byte, 0, k)
	start := 0

	for len(result) < k {
		remainingToSelect := k - len(result)
		remainingPositions := n - start

		// We can skip at most (remainingPositions - remainingToSelect) positions
		// So we search in the first (remainingPositions - remainingToSelect + 1) positions
		searchRange := remainingPositions - remainingToSelect + 1

		// Find the maximum digit in the allowable search range
		maxDigit := bank[start]
		maxPos := start

		for i := start; i < start+searchRange; i++ {
			if bank[i] > maxDigit {
				maxDigit = bank[i]
				maxPos = i
			}
		}

		result = append(result, maxDigit)
		start = maxPos + 1
	}

	return string(result)
}
