package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/IanShearer/aoc/cmd/aoc/internal"
)

func createDay() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: aoc create <day_number>")
		fmt.Println("Example: aoc create 5")
		os.Exit(1)
	}

	var dayNum int
	_, err := fmt.Sscanf(os.Args[2], "%d", &dayNum)
	if err != nil || dayNum < 1 || dayNum > 25 {
		fmt.Printf("Error: day number must be between 1 and 25\n")
		os.Exit(1)
	}

	// Format day with leading zero
	dayDir := fmt.Sprintf("day%02d", dayNum)

	// Create main day directory
	if err := os.Mkdir(dayDir, 0755); err != nil {
		fmt.Printf("Error creating directory %s: %v\n", dayDir, err)
		os.Exit(1)
	}

	// Load session cookie and fetch input
	sessionCookie, err := internal.LoadSessionCookie()
	if err != nil {
		fmt.Printf("Error loading session cookie: %v\n", err)
		os.Exit(1)
	}

	inputContent, err := internal.FetchInput(dayNum, sessionCookie)
	if err != nil {
		fmt.Printf("Error fetching input: %v\n", err)
		os.Exit(1)
	}

	// Create input file with fetched content
	inputFile := filepath.Join(dayDir, "input")
	if err := os.WriteFile(inputFile, []byte(inputContent), 0644); err != nil {
		fmt.Printf("Error creating input file: %v\n", err)
		os.Exit(1)
	}

	answersFile := filepath.Join(dayDir, "answers")
	if err := os.WriteFile(answersFile, []byte(""), 0644); err != nil {
		fmt.Printf("Error creating answers file: %v\n", err)
		os.Exit(1)
	}

	// Create ai and human directories
	aiDir := filepath.Join(dayDir, "ai")
	if err := os.Mkdir(aiDir, 0755); err != nil {
		fmt.Printf("Error creating ai directory: %v\n", err)
		os.Exit(1)
	}

	humanDir := filepath.Join(dayDir, "human")
	if err := os.Mkdir(humanDir, 0755); err != nil {
		fmt.Printf("Error creating human directory: %v\n", err)
		os.Exit(1)
	}

	// Create main.go in human directory
	mainGoContent := "package main\n"
	mainGoFile := filepath.Join(humanDir, "main.go")
	if err := os.WriteFile(mainGoFile, []byte(mainGoContent), 0644); err != nil {
		fmt.Printf("Error creating main.go: %v\n", err)
		os.Exit(1)
	}

	// Create main_test.go in human directory
	mainTestGoFile := filepath.Join(humanDir, "main_test.go")
	if err := os.WriteFile(mainTestGoFile, []byte(mainGoContent), 0644); err != nil {
		fmt.Printf("Error creating main_test.go: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created directory structure for %s\n", dayDir)
}
