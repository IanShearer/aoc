package main

import (
	"bufio"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "create":
		createDay()
	case "redact":
		redactDay()
	case "fetch":
		fetchDay()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: aoc <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  create <day_number>  Create directory structure for a day")
	fmt.Println("  redact <day_number>  Redact answers and puzzle text from conversation")
	fmt.Println("  fetch <day_number>   Fetch puzzle content from adventofcode.com")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  aoc create 5")
	fmt.Println("  aoc redact 4")
	fmt.Println("  aoc fetch 7")
}

// CREATE COMMAND HELPERS
func loadSessionCookie() (string, error) {
	content, err := os.ReadFile(".env")
	if err != nil {
		return "", fmt.Errorf("failed to read .env file: %w", err)
	}

	// Parse the .env file to find session=VALUE
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "session=") {
			return strings.TrimPrefix(line, "session="), nil
		}
	}

	return "", fmt.Errorf("session cookie not found in .env file")
}

func fetchInput(dayNum int, sessionCookie string) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/2025/day/%d/input", dayNum)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Add session cookie
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch input: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch input: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

// CREATE COMMAND
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
	sessionCookie, err := loadSessionCookie()
	if err != nil {
		fmt.Printf("Error loading session cookie: %v\n", err)
		os.Exit(1)
	}

	inputContent, err := fetchInput(dayNum, sessionCookie)
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

// REDACT COMMAND
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

// FETCH COMMAND
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
	sessionCookie, err := loadSessionCookie()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading session cookie: %v\n", err)
		os.Exit(1)
	}

	// Fetch puzzle HTML
	htmlContent, err := fetchPuzzleHTML(dayNum, sessionCookie)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching puzzle HTML: %v\n", err)
		os.Exit(1)
	}

	// Extract puzzle content
	puzzleText := extractPuzzleContent(htmlContent, dayNum)

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

func fetchPuzzleHTML(dayNum int, sessionCookie string) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/2025/day/%d", dayNum)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Add session cookie
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch HTML: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch HTML: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

func extractPuzzleContent(htmlContent string, dayNum int) string {
	// Convert HTML to plain text
	text := htmlToText(htmlContent)

	// Find the start marker: "--- Day N:"
	startMarker := fmt.Sprintf("--- Day %d:", dayNum)
	startIdx := strings.Index(text, startMarker)
	if startIdx == -1 {
		return ""
	}

	// Find end markers to determine where to stop extraction
	endMarkers := []string{
		"To begin, get your puzzle input",
		"Although it hasn't changed",
		"Both parts of this puzzle are complete",
		"At this point, you should return to your Advent calendar",
	}

	endIdx := -1
	for _, marker := range endMarkers {
		idx := strings.Index(text[startIdx:], marker)
		if idx != -1 {
			endIdx = idx
			break
		}
	}

	var extracted string
	if endIdx != -1 {
		// Found an end marker, extract up to but not including it
		extracted = text[startIdx : startIdx+endIdx]
	} else {
		// No end marker found, take everything from start to end
		extracted = text[startIdx:]
	}

	// Remove all lines containing "Your puzzle answer was"
	extracted = removeAnswerLines(extracted)

	// Clean up the extracted text
	extracted = strings.TrimSpace(extracted)

	return extracted
}

func removeAnswerLines(content string) string {
	lines := strings.Split(content, "\n")
	var filtered []string

	for _, line := range lines {
		// Skip lines that contain "Your puzzle answer was"
		if !strings.Contains(line, "Your puzzle answer was") {
			filtered = append(filtered, line)
		}
	}

	return strings.Join(filtered, "\n")
}

func htmlToText(htmlContent string) string {
	// Remove script and style tags completely
	scriptRegex := regexp.MustCompile(`(?is)<script[^>]*>.*?</script>`)
	htmlContent = scriptRegex.ReplaceAllString(htmlContent, "")

	styleRegex := regexp.MustCompile(`(?is)<style[^>]*>.*?</style>`)
	htmlContent = styleRegex.ReplaceAllString(htmlContent, "")

	// Replace block-level tags with newlines
	blockTags := []string{"p", "div", "article", "section", "h1", "h2", "h3", "h4", "h5", "h6", "li", "br"}
	for _, tag := range blockTags {
		openRegex := regexp.MustCompile(fmt.Sprintf(`(?i)<%s[^>]*>`, tag))
		htmlContent = openRegex.ReplaceAllString(htmlContent, "\n")

		closeRegex := regexp.MustCompile(fmt.Sprintf(`(?i)</%s>`, tag))
		htmlContent = closeRegex.ReplaceAllString(htmlContent, "\n")
	}

	// Handle <pre> and <code> tags specially - just remove them without adding newlines
	codeRegex := regexp.MustCompile(`(?i)</?(?:code|pre|em|strong|span|a)[^>]*>`)
	htmlContent = codeRegex.ReplaceAllString(htmlContent, "")

	// Remove all remaining HTML tags
	tagRegex := regexp.MustCompile(`<[^>]+>`)
	htmlContent = tagRegex.ReplaceAllString(htmlContent, "")

	// Decode HTML entities
	htmlContent = html.UnescapeString(htmlContent)

	// Clean up excessive newlines (more than 2 in a row)
	multiNewlineRegex := regexp.MustCompile(`\n{3,}`)
	htmlContent = multiNewlineRegex.ReplaceAllString(htmlContent, "\n\n")

	return htmlContent
}
