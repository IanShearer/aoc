package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	ErrEmptyInput = errors.New("empty input")
)

const (
	STARTING_LOCATION int = 50
	MAX_POSITION      int = 100
	MIN_POSITION      int = 0
)

// Direction is an enum that symbolizes either a left or right turn
type Direction int

const (
	NoDirection Direction = iota
	Left
	Right
)

func ParseDirection(input rune) (Direction, error) {
	switch input {
	case 'L':
		return Left, nil
	case 'R':
		return Right, nil
	default:
		return NoDirection, fmt.Errorf("invalid input, expecting L or R: %v", input)
	}
}

type Lock struct {
	Position int

	positionZeroCounter uint
	numberOfClicks      uint
}

func NewLock() Lock {
	return Lock{Position: STARTING_LOCATION}
}

func (l *Lock) TimesAtZero() uint {
	return l.positionZeroCounter
}

func (l *Lock) NumberOfClicks() uint {
	return l.numberOfClicks
}

// Turn turns the lock by a given input
func (l *Lock) Turn(input string) error {
	direction, turns, err := l.parse(input)
	if err != nil {
		return err
	}

	if direction == Right {
		l.turnRight(turns)
	} else {
		l.turnLeft(turns)
	}

	// For part two: we count the number of clicks. A quick way to check the minimum number of times we
	// made a click is to just divide by the number of full rotations we made.
	//
	// we have to check if we loop pass zero in the turn functions up above. if we do, we need to increment the number
	// of clicks by one more.
	l.numberOfClicks += uint(int(turns) / MAX_POSITION)

	// For part one: if at the end of turning the dial we end on zero
	// we increment that as a stop at zero which is accumulated for part ones answer
	if l.Position == 0 {
		l.positionZeroCounter++
	}
	return nil
}

// parse parses a single line into a direction and number of turns it has/
//
// the format for a line is either a 'R' or 'L' for the first character, which
// corespond with "Right" and "Left"
// and a number for the remaining characters. Which symbolizes the number of turns
func (l *Lock) parse(input string) (Direction, int, error) {
	if len(input) == 0 {
		return NoDirection, 0, ErrEmptyInput
	}

	d, err := ParseDirection(rune(input[0]))
	if err != nil {
		return NoDirection, 0, err
	}

	turns, err := strconv.ParseInt(input[1:], 10, 64)
	if err != nil {
		return NoDirection, 0, err
	}

	return d, int(turns), nil
}

func (l *Lock) turnLeft(turns int) {
	// check if we made a click while not being on the starting zero position
	if (l.Position-int(turns%MAX_POSITION)) <= MIN_POSITION && l.Position != 0 {
		l.numberOfClicks++
	}

	// update position
	l.Position = (l.Position - int(turns%MAX_POSITION)) % MAX_POSITION
	if l.Position < MIN_POSITION {
		l.Position += MAX_POSITION
	}
}

func (l *Lock) turnRight(turns int) {
	// check if we made a click while not being on the starting zero position
	if (l.Position+int(turns%MAX_POSITION)) >= MAX_POSITION && l.Position != 0 {
		l.numberOfClicks++
	}

	// update position
	l.Position = (l.Position + int(turns%MAX_POSITION)) % MAX_POSITION
}

func main() {
	input, err := os.Open("../input")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	lock := NewLock()
	for scanner.Scan() {
		err := lock.Turn(scanner.Text())
		if err != nil {
			log.Fatalf("failed to turn lock: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while scanning input: %v", err)
	}

	fmt.Printf("Part One: %d\n", lock.TimesAtZero())
	fmt.Printf("Part Two: %d\n", lock.NumberOfClicks())
}
