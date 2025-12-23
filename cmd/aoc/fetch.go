package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/IanShearer/aoc/cmd/aoc/internal"
)

func fetchDay() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: aoc fetch <day_number>\n")
		fmt.Fprintf(os.Stderr, "Example: aoc fetch 7\n")
		os.Exit(1)
	}

	dayNum, err := strconv.Atoi(os.Args[2])
	if err != nil || dayNum < 1 || dayNum > 25 {
		fmt.Fprintf(os.Stderr, "Error: day number must be between 1 and 25\n")
		os.Exit(1)
	}

	// Load session cookie
	sessionCookie, err := internal.LoadSessionCookie()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading session cookie: %v\n", err)
		os.Exit(1)
	}

	// Fetch puzzle HTML
	htmlContent, err := internal.FetchPuzzleHTML(dayNum, sessionCookie)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching puzzle HTML: %v\n", err)
		os.Exit(1)
	}

	// Extract puzzle content
	puzzleText := internal.ExtractPuzzleContent(htmlContent, dayNum)

	// Write to file
	dayStr := fmt.Sprintf("%02d", dayNum)
	dayDir := fmt.Sprintf("day%s", dayStr)
	outputPath := filepath.Join(dayDir, fmt.Sprintf("day%s_content.txt", dayStr))

	err = os.WriteFile(outputPath, []byte(puzzleText), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing content file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully fetched puzzle content for day %d to %s\n", dayNum, outputPath)
}
