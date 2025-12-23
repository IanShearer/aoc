package internal

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

func ExtractPuzzleContent(htmlContent string, dayNum int) string {
	// Convert HTML to plain text
	text := HtmlToText(htmlContent)

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
	extracted = RemoveAnswerLines(extracted)

	// Clean up the extracted text
	extracted = strings.TrimSpace(extracted)

	return extracted
}

func RemoveAnswerLines(content string) string {
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

func HtmlToText(htmlContent string) string {
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
