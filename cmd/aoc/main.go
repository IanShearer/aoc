package main

import (
	"fmt"
	"os"
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
