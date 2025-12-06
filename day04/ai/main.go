package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, err := readGrid("../input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	partOne := countAccessibleRolls(grid)
	fmt.Printf("Part One: %d\n", partOne)
}

func readGrid(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func countAccessibleRolls(grid []string) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Directions for 8 neighbors: up, down, left, right, and 4 diagonals
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] != '@' {
				continue
			}

			// Count adjacent rolls
			adjacentRolls := 0
			for _, dir := range directions {
				newRow := row + dir[0]
				newCol := col + dir[1]

				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
					if grid[newRow][newCol] == '@' {
						adjacentRolls++
					}
				}
			}

			if adjacentRolls < 4 {
				count++
			}
		}
	}

	return count
}
