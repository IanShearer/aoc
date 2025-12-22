package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operator uint

const (
	UnknownOperator Operator = iota
	Plus
	Multiply
)

type Position uint

const (
	UnknownPosition Position = iota
	Left
	Right
)

type Column struct {
	Numbers []Number

	Operator      Operator
	OperatorIndex int

	Position Position
}

type Number struct {
	Value  int
	String string
	Index  int
}

func performMath(numbers []int, operator Operator) int {
	sum := 0
	for _, n := range numbers {
		if operator == Multiply {
			if sum == 0 {
				sum = n
			} else {
				sum = sum * n
			}
		} else {
			sum = sum + n
		}
	}

	return sum
}

func (c Column) PartOne() int {
	numbers := make([]int, len(c.Numbers))
	for i := range c.Numbers {
		numbers[i] = c.Numbers[i].Value
	}
	return performMath(numbers, c.Operator)
}

// overcomplicated but it works :woozy:
func (c Column) PartTwo() int {
	// create the new numbers
	maxDistance := 0
	for _, n := range c.Numbers {
		div := 1
		val := 1
		dis := 0
		for val != 0 {
			val = n.Value / div
			div *= 10
			dis++
		}

		dis -= 1
		if dis > maxDistance {
			maxDistance = dis
		}
	}
	numbers := make([]int, maxDistance)
	numberStrings := make([]string, maxDistance)

	// each number here is basically just the partial to a row
	// with their given start index, we know when the character will start
	for _, n := range c.Numbers {
		distanceFromOperator := n.Index - c.OperatorIndex
		maxDistance := len(n.String)
		for j := range maxDistance {
			numberStrings[distanceFromOperator+j] += string(n.String[j])
		}
	}

	for i, n := range numberStrings {
		number, _ := strconv.ParseInt(n, 10, 64)
		numbers[i] = int(number)
	}

	return performMath(numbers, c.Operator)
}

var numberRegex = regexp.MustCompile(`\d+`)

func isNumberRow(row string) bool {
	return numberRegex.MatchString(row)
}

func ParseInput(scanner *bufio.Scanner) []Column {
	columns := make([]Column, 0)

	firstRow := true
	for scanner.Scan() {
		row := scanner.Text() + "\n" // add new line to parse full number at the end of the row

		if isNumberRow(row) {
			var currentNumber strings.Builder
			colIndex := 0
			currentIndex := 0
			for index, character := range row {
				if character >= '0' && character <= '9' {
					currentNumber.WriteRune(character)
					currentIndex = index
				} else {
					if currentNumber.Len() > 0 {
						num, err := strconv.ParseInt(currentNumber.String(), 10, 64)
						if err != nil {
							panic(err)
						}

						n := Number{
							Value:  int(num),
							String: currentNumber.String(),
							Index:  currentIndex - currentNumber.Len() + 1,
						}

						if firstRow {
							columns = append(columns, Column{})
						}
						columns[colIndex].Numbers = append(columns[colIndex].Numbers, n)

						colIndex++
						currentNumber.Reset()
					}
				}
			}
		} else {
			colIndex := 0
			for index, symbol := range row {
				switch symbol {
				case '+':
					columns[colIndex].Operator = Plus
					columns[colIndex].OperatorIndex = index
					colIndex++
				case '*':
					columns[colIndex].Operator = Multiply
					columns[colIndex].OperatorIndex = index
					colIndex++
				}
			}
		}

		firstRow = false
	}

	return columns
}

func main() {
	input, err := os.Open("../input")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	columns := ParseInput(scanner)
	partOne := 0
	partTwo := 0
	for _, c := range columns {
		p1 := c.PartOne()
		partOne += p1
		p2 := c.PartTwo()
		partTwo += p2
	}

	fmt.Printf("Part One: %d\n", partOne)
	fmt.Printf("Part Two: %d\n", partTwo)
}
