package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReadAnswers(path string) ([]string, error) {
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

func RedactPuzzleBlocks(content string, dayNum int) string {
	// Regular expression to match code blocks
	codeBlockRegex := regexp.MustCompile("(?s)```\\n(.*?)\\n```")

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
		if IsPuzzleBlock(blockContent) {
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

func IsPuzzleBlock(content string) bool {
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
