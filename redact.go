package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <day_number>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s 1\n", os.Args[0])
		os.Exit(1)
	}

	dayNum, err := strconv.Atoi(os.Args[1])
	if err != nil || dayNum < 1 || dayNum > 25 {
		fmt.Fprintf(os.Stderr, "Error: day number must be between 1 and 25\n")
		os.Exit(1)
	}

	// Format day number with leading zero if needed
	dayStr := fmt.Sprintf("%02d", dayNum)

	// Read answers file
	answersPath := filepath.Join(fmt.Sprintf("day%s", dayStr), "answers")
	answers, err := readAnswers(answersPath)
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
	redacted = redactPuzzleBlocks(redacted, dayNum)

	// Write back to file
	err = os.WriteFile(conversationPath, []byte(redacted), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing conversation file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully redacted day %d conversation\n", dayNum)
}

func readAnswers(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var answers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			answers = append(answers, line)
		}
	}

	return answers, scanner.Err()
}

func redactPuzzleBlocks(content string, dayNum int) string {
	// Regular expression to match code blocks
	codeBlockRegex := regexp.MustCompile("(?s)```\n(.*?)\n```")

	partNum := 1
	alreadyRedactedPattern := regexp.MustCompile(`\(REDACTED\) the text in this box is the puzzle`)

	result := codeBlockRegex.ReplaceAllStringFunc(content, func(match string) string {
		// Extract the content between the backticks
		blockContent := match[4 : len(match)-4] // Remove ``` from start and end

		// Check if this block is already redacted
		if alreadyRedactedPattern.MatchString(blockContent) {
			return match
		}

		// Check if this looks like a puzzle description
		if isPuzzleBlock(blockContent) {
			partLabel := "one"
			if partNum == 2 {
				partLabel = "two"
			}
			partNum++

			return fmt.Sprintf("```\n(REDACTED) the text in this box is the puzzle, part %s of advent of code 2025 day %02d\n```", partLabel, dayNum)
		}

		return match
	})

	return result
}

func isPuzzleBlock(content string) bool {
	// Heuristics to identify puzzle blocks:
	// 1. Contains typical puzzle language
	// 2. Relatively long (> 200 characters)
	// 3. Contains narrative/story elements
	// 4. Not code (doesn't have typical code patterns)

	if len(content) < 100 {
		return false
	}

	// Check for common puzzle indicators
	puzzleIndicators := []string{
		"you need to",
		"you must",
		"you find",
		"you arrive",
		"your puzzle",
		"for example",
		"--- Day",
		"--- Part",
		"What is",
		"How many",
		"Find the",
		"Calculate",
	}

	lowerContent := strings.ToLower(content)
	for _, indicator := range puzzleIndicators {
		if strings.Contains(lowerContent, strings.ToLower(indicator)) {
			return true
		}
	}

	// Check if it looks like code (has many special programming characters)
	codeIndicators := []string{
		"func ", "package ", "import ", "def ", "class ", "public ", "private ",
		"const ", "let ", "var ", "function ", "return ", "if (", "for (", "while (",
	}

	for _, indicator := range codeIndicators {
		if strings.Contains(lowerContent, indicator) {
			return false
		}
	}

	// If content has many lines and contains question marks or exclamation points,
	// it's likely narrative text
	lines := strings.Split(content, "\n")
	if len(lines) > 10 && (strings.Contains(content, "?") || strings.Contains(content, "!")) {
		return true
	}

	return false
}
