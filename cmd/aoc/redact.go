package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/IanShearer/aoc/cmd/aoc/internal"
)

func redactDay() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: aoc redact <day_number>\n")
		fmt.Fprintf(os.Stderr, "Example: aoc redact 1\n")
		os.Exit(1)
	}

	dayNum, err := strconv.Atoi(os.Args[2])
	if err != nil || dayNum < 1 || dayNum > 25 {
		fmt.Fprintf(os.Stderr, "Error: day number must be between 1 and 25\n")
		os.Exit(1)
	}

	// Format day number with leading zero if needed
	dayStr := fmt.Sprintf("%02d", dayNum)

	// Read answers file
	answersPath := filepath.Join(fmt.Sprintf("day%s", dayStr), "answers")
	answers, err := internal.ReadAnswers(answersPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading answers file: %v\n", err)
		os.Exit(1)
	}

	// Read conversation file
	conversationPath := filepath.Join(fmt.Sprintf("day%s", dayStr), "ai", fmt.Sprintf("day%s_conversation.txt", dayStr))
	content, err := os.ReadFile(conversationPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading conversation file: %v\n", err)
		os.Exit(1)
	}

	// Perform redactions
	redacted := string(content)

	// 1. Redact answers
	for _, answer := range answers {
		if answer != "" {
			redacted = strings.ReplaceAll(redacted, answer, "(REDACTED)")
		}
	}

	// 2. Redact puzzle code blocks
	redacted = internal.RedactPuzzleBlocks(redacted, dayNum)

	// Write back to file
	err = os.WriteFile(conversationPath, []byte(redacted), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing conversation file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully redacted day %d conversation\n", dayNum)
}
