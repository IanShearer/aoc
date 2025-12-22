package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Min uint
	Max uint
}

func ParseInput(scanner *bufio.Scanner) ([]Range, []uint) {
	ranges := make([]Range, 0)
	ingredientIDs := make([]uint, 0)
	finishedParsingRanges := false
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "-") {
			s := strings.Split(line, "-")
			newRange := Range{}

			l, err := strconv.ParseUint(s[0], 10, 64)
			if err != nil {
				log.Fatalf("failed to parse left side of range: %v", err)
			}

			r, err := strconv.ParseUint(s[1], 10, 64)
			if err != nil {
				log.Fatalf("failed to parse right side of range: %v", err)
			}

			newRange.Min = uint(l)
			newRange.Max = uint(r)
			ranges = append(ranges, newRange)

			continue
		}

		if len(line) == 0 {
			finishedParsingRanges = true
			continue
		}

		if finishedParsingRanges {
			id, err := strconv.ParseUint(line, 10, 64)
			if err != nil {
				log.Fatalf("failed to parse id: %v", err)
			}

			ingredientIDs = append(ingredientIDs, uint(id))
		}
	}

	return ranges, ingredientIDs
}

func FreshIngredientIDs(ranges []Range, ingredientIDs []uint) uint {
	var freshIngredientCount uint

	for _, id := range ingredientIDs {
		for _, r := range ranges {
			if r.Min <= id && id <= r.Max {
				freshIngredientCount++
				break
			}
		}
	}

	return freshIngredientCount
}

func merge(r1, r2 Range) (bool, Range) {
	if r1.Min <= r2.Min && r1.Max >= r2.Min {
		return true, Range{min(r1.Min, r2.Min), max(r1.Max, r2.Max)}
	}

	return false, Range{}
}

func mergedRanges(ranges []Range, index int) []Range {
	if index >= len(ranges) {
		return ranges
	}

	for i := range len(ranges) {
		if index == i {
			continue
		}

		didMerge, newRange := merge(ranges[index], ranges[i])
		if didMerge {
			if i > index {
				ranges = slices.Delete(ranges, i, i+1)
				ranges = slices.Delete(ranges, index, index+1)
			} else {
				ranges = slices.Delete(ranges, index, index+1)
				ranges = slices.Delete(ranges, i, i+1)
			}

			ranges = append(ranges, newRange)
			sort.Slice(ranges, func(i int, j int) bool {
				return ranges[i].Min < ranges[j].Min
			})
			return mergedRanges(ranges, index)
		}
	}

	return mergedRanges(ranges, index+1)
}

func FreshIngredientIDRangeCount(ranges []Range) uint {
	// compact the ranges
	sort.Slice(ranges, func(i int, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})

	compactedRanges := mergedRanges(ranges, 0)

	// accumulate the size of the new ranges
	var freshIDsCount uint
	for _, r := range compactedRanges {
		// + 1 because we need to be inclusive
		freshIDsCount += r.Max - r.Min + 1
	}

	return freshIDsCount
}

func main() {
	input, err := os.Open("../input")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	ranges, ingredientIDs := ParseInput(scanner)

	fmt.Printf("Part One: %d\n", FreshIngredientIDs(ranges, ingredientIDs))
	fmt.Printf("Part Two: %d\n", FreshIngredientIDRangeCount(ranges))
}
