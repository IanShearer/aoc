package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func (r Range) contains(id int) bool {
	return id >= r.start && id <= r.end
}

func (r Range) size() int {
	return r.end - r.start + 1
}

// mergeRanges takes a list of ranges and merges overlapping/adjacent ones
func mergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	// Sort ranges by start position
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		current := ranges[i]

		// If current range overlaps or is adjacent to last merged range
		if current.start <= last.end+1 {
			// Merge by extending the end if necessary
			if current.end > last.end {
				last.end = current.end
			}
		} else {
			// No overlap, add as new range
			merged = append(merged, current)
		}
	}

	return merged
}

func solve(scanner *bufio.Scanner) (int, int) {
	var ranges []Range

	// Parse ranges
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, Range{start: start, end: end})
	}

	// Part Two: Count all unique IDs in merged ranges
	merged := mergeRanges(ranges)
	totalFreshIDs := 0
	for _, r := range merged {
		totalFreshIDs += r.size()
	}

	// Part One: Check ingredient IDs
	freshCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		id, _ := strconv.Atoi(line)

		// Check if ID is in any range
		for _, r := range ranges {
			if r.contains(id) {
				freshCount++
				break
			}
		}
	}

	return freshCount, totalFreshIDs
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	partOne, partTwo := solve(scanner)

	fmt.Printf("Part One: %d\n", partOne)
	fmt.Printf("Part Two: %d\n", partTwo)
}
