package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type IDRange struct {
	Min uint
	Max uint

	PartOneInvalidIDs []uint
	PartTwoInvalidIDs []uint
}

func (r *IDRange) FindInvalidIDs() {
	for i := r.Min; i <= r.Max; i++ {
		s := strconv.FormatUint(uint64(i), 10)
		half := len(s) / 2

		// part one: has to be exactly half
		if len(s)%2 == 0 {
			if strings.EqualFold(s[half:], s[:half]) {
				r.PartOneInvalidIDs = append(r.PartOneInvalidIDs, i)
			}
		}

		// part two: all repeating numbers, any length
		for j := 0; j <= half; j++ {
			arr := splitIntoChunks(s, j)
			uniq := slices.Compact(arr)
			if len(arr) > 1 && len(uniq) == 1 {
				r.PartTwoInvalidIDs = append(r.PartTwoInvalidIDs, i)
				break
			}
		}
	}
}

func splitIntoChunks(s string, n int) []string {
	if n <= 0 {
		return nil
	}

	chunks := make([]string, 0)
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

func NewIDRange(input string) (IDRange, error) {
	idRange := IDRange{}

	ids := strings.Split(input, "-")
	if len(ids) != 2 {
		return idRange, errors.New("invalid input")
	}

	min, err := strconv.ParseUint(ids[0], 10, 64)
	if err != nil {
		return idRange, err
	}

	max, err := strconv.ParseUint(ids[1], 10, 64)
	if err != nil {
		return idRange, err
	}

	idRange.Min = uint(min)
	idRange.Max = uint(max)
	idRange.PartOneInvalidIDs = make([]uint, 0)
	idRange.PartTwoInvalidIDs = make([]uint, 0)

	return idRange, nil
}

func ParseInput() ([]IDRange, error) {
	input, err := os.Open("../input")
	if err != nil {
		return nil, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	ranges := make([]IDRange, 0)
	for scanner.Scan() {
		line := scanner.Text()

		for r := range strings.SplitSeq(line, ",") {
			newRange, err := NewIDRange(r)
			if err != nil {
				return nil, err
			}
			ranges = append(ranges, newRange)
		}
	}

	return ranges, nil
}

func main() {
	ranges, err := ParseInput()
	if err != nil {
		log.Fatalf("failed to parse input: %v", err)
	}

	var sumPartOneInvalidIDs uint
	var sumPartTwoInvalidIDs uint
	for _, r := range ranges {
		r.FindInvalidIDs()

		for _, invalid := range r.PartOneInvalidIDs {
			sumPartOneInvalidIDs += invalid
		}

		for _, invalid := range r.PartTwoInvalidIDs {
			sumPartTwoInvalidIDs += invalid
		}
	}

	fmt.Printf("Part One: %d\n", sumPartOneInvalidIDs)
	fmt.Printf("Part Two: %d\n", sumPartTwoInvalidIDs)
}
