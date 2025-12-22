package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input
	file, err := os.Open("../input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read all lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Parse and solve
	partOne := solvePartOne(lines)
	fmt.Println("Part One:", partOne)

	partTwo := solvePartTwo(lines)
	fmt.Println("Part Two:", partTwo)
}

func solvePartOne(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	// Find the maximum width
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Extract columns - transpose the grid
	columns := make([][]string, maxWidth)
	for col := 0; col < maxWidth; col++ {
		columns[col] = make([]string, len(lines))
		for row := 0; row < len(lines); row++ {
			if col < len(lines[row]) {
				columns[col][row] = string(lines[row][col])
			} else {
				columns[col][row] = " "
			}
		}
	}

	// Group consecutive non-space columns into problems
	var problems [][][]string
	var currentProblem [][]string

	for _, col := range columns {
		// Check if this column is all spaces
		allSpaces := true
		for _, cell := range col {
			if strings.TrimSpace(cell) != "" {
				allSpaces = false
				break
			}
		}

		if allSpaces {
			// End current problem if any
			if len(currentProblem) > 0 {
				problems = append(problems, currentProblem)
				currentProblem = nil
			}
		} else {
			// Add this column to current problem
			currentProblem = append(currentProblem, col)
		}
	}

	// Don't forget the last problem
	if len(currentProblem) > 0 {
		problems = append(problems, currentProblem)
	}

	// Calculate result for each problem
	grandTotal := 0
	for _, problem := range problems {
		result := solveProblem(problem)
		grandTotal += result
	}

	return grandTotal
}

func solveProblem(problem [][]string) int {
	if len(problem) == 0 {
		return 0
	}

	// The problem is a 2D array where each element is problem[col][row]
	// We need to transpose it to get rows
	numRows := len(problem[0])

	// Last row contains the operator
	operatorRow := numRows - 1

	// Extract operator (find first non-space in operator row)
	operator := ""
	for _, col := range problem {
		if operatorRow < len(col) {
			cell := strings.TrimSpace(col[operatorRow])
			if cell != "" {
				operator = cell
				break
			}
		}
	}

	// Extract numbers from rows 0 to operatorRow-1
	var numbers []int
	for row := 0; row < operatorRow; row++ {
		// Build the number from all columns in this row
		var numStr strings.Builder
		for _, col := range problem {
			if row < len(col) {
				numStr.WriteString(col[row])
			}
		}

		// Parse the number
		trimmed := strings.TrimSpace(numStr.String())
		if trimmed != "" {
			num, err := strconv.Atoi(trimmed)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	// Calculate result based on operator
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if operator == "*" {
			result *= numbers[i]
		} else if operator == "+" {
			result += numbers[i]
		}
	}

	return result
}

func solvePartTwo(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	// Find the maximum width
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Extract columns - transpose the grid
	columns := make([][]string, maxWidth)
	for col := 0; col < maxWidth; col++ {
		columns[col] = make([]string, len(lines))
		for row := 0; row < len(lines); row++ {
			if col < len(lines[row]) {
				columns[col][row] = string(lines[row][col])
			} else {
				columns[col][row] = " "
			}
		}
	}

	// Group consecutive non-space columns into problems
	var problems [][][]string
	var currentProblem [][]string

	for _, col := range columns {
		// Check if this column is all spaces
		allSpaces := true
		for _, cell := range col {
			if strings.TrimSpace(cell) != "" {
				allSpaces = false
				break
			}
		}

		if allSpaces {
			// End current problem if any
			if len(currentProblem) > 0 {
				problems = append(problems, currentProblem)
				currentProblem = nil
			}
		} else {
			// Add this column to current problem
			currentProblem = append(currentProblem, col)
		}
	}

	// Don't forget the last problem
	if len(currentProblem) > 0 {
		problems = append(problems, currentProblem)
	}

	// Calculate result for each problem using Part Two logic
	grandTotal := 0
	for _, problem := range problems {
		result := solveProblemPartTwo(problem)
		grandTotal += result
	}

	return grandTotal
}

func solveProblemPartTwo(problem [][]string) int {
	if len(problem) == 0 {
		return 0
	}

	// In Part Two, we read columns right-to-left
	// Each column (from top to bottom, excluding operator row) is a number
	// The operator is in the last row

	numRows := len(problem[0])
	operatorRow := numRows - 1

	// Extract operator (find first non-space in operator row, reading right to left)
	operator := ""
	for i := len(problem) - 1; i >= 0; i-- {
		col := problem[i]
		if operatorRow < len(col) {
			cell := strings.TrimSpace(col[operatorRow])
			if cell != "" {
				operator = cell
				break
			}
		}
	}

	// Extract numbers by reading columns right-to-left
	var numbers []int
	for i := len(problem) - 1; i >= 0; i-- {
		col := problem[i]

		// Build number from this column (top to bottom, excluding operator row)
		var numStr strings.Builder
		for row := 0; row < operatorRow; row++ {
			if row < len(col) {
				numStr.WriteString(col[row])
			}
		}

		// Parse the number
		trimmed := strings.TrimSpace(numStr.String())
		if trimmed != "" {
			num, err := strconv.Atoi(trimmed)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	// Calculate result based on operator
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if operator == "*" {
			result *= numbers[i]
		} else if operator == "+" {
			result += numbers[i]
		}
	}

	return result
}
