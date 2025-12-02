# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is an Advent of Code 2025 solutions repository written in Go. The repository contains solutions to daily programming challenges, with separate implementations for human-written and AI-assisted solutions.

Your job is to create a solution desginated in the /ai directory of the given day you will be working in.

## Rules

You may not look at any code in the `/human` directories. Come up with the solutions on your own.

## Project Structure

The repository is organized by day, with each day containing:
- `dayXX/input` - The puzzle input (shared across all solutions)
- `dayXX/human/` - Human-written solution
- `dayXX/ai/` - AI-assisted solution (may be empty initially)

Each solution directory contains:
- `main.go` - The solution implementation
- `main_test.go` - Tests using sample inputs from the problem

## Common Commands

### Running a solution
```bash
cd dayXX/ai
go run main.go
```

### Running tests
```bash
# Run tests for a specific day
cd dayXX/ai
go test

# Run tests with verbose output
go test -v
```

### Running a single test
```bash
cd dayXX/ai
go test -run TestPartOneSample
go test -run TestPartTwoSample
```

## Code Architecture

### Solution Pattern

Each day's solution follows a consistent pattern:

1. **Input Parsing**: Solutions read from `../input` (relative to the solution directory)
2. **Problem Domain Types**: Custom types model the problem domain (e.g., `Lock`, `Direction` for Day 1)
3. **Stateful Processing**: Solutions typically use a struct to maintain state while processing line-by-line input
4. **Dual Output**: `main()` prints both Part One and Part Two answers
5. **Readability**: Code should be written in a way that it is easy for a human to understand and pick up

### Testing Pattern

Tests use the sample inputs provided in each day's problem description:
- `TestPartOneSample` validates Part One logic
- `TestPartTwoSample` validates Part Two logic
- Tests use `strings.NewReader` to simulate file input from inline strings

### Input Handling

All solutions:
- Use `bufio.Scanner` for line-by-line input processing
- Expect input files at `../input` relative to the solution directory
- Handle errors from file operations and parsing

### Verify Solutions

Soltions to the problems will be in  the file `../answers`. The part ones answer will be on the first line, the part twos answer will be on the second line
```
1234
4321
```