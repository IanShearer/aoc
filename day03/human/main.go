package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const maxNum byte = '9'

func FindHighestJoltage(input string) int64 {
	left := input[0]
	right := byte('0')

	for i, c := range input {
		if i == 0 {
			continue
		}

		if byte(c) > left && i != len(input)-1 {
			left = byte(c)
			right = byte('0')
			continue
		}

		if byte(c) > right {
			right = byte(c)
		}
	}

	s := string([]byte{left, right})
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}

func FindHighestJoltageTwelveBatteries(input string) int64 {
	batteries := make([]byte, 12)
	findMaxForGivenIndex := func(start int, maxEnd int) (int, byte) {
		// we start from the start index for the input, check for the largest number up until the len(input)-maxEnd
		currentMax := byte('0')
		currentMaxPostion := 0
		for i := range len(input) - maxEnd - start {
			if input[start+i] > currentMax {
				currentMax = input[start+i]
				currentMaxPostion = i
			}

			if currentMax == maxNum {
				return start + i + 1, maxNum
			}
		}

		return currentMaxPostion + start + 1, currentMax
	}

	currentPos := 0
	for i := range batteries {
		currentPos, batteries[i] = findMaxForGivenIndex(currentPos, len(batteries)-i-1)
	}

	s := string(batteries)
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}

func main() {
	input, err := os.Open("../input")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var partOne int64
	var partTwo int64
	for scanner.Scan() {
		line := scanner.Text()
		partOne += FindHighestJoltage(line)
		partTwo += FindHighestJoltageTwelveBatteries(line)
	}

	fmt.Printf("Part One: %d\n", partOne)
	fmt.Printf("Part Two: %d\n", partTwo)
}
