package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Content uint

const (
	Unknown Content = iota
	Empty
	Paper
	PaperMarkedRemoved
)

func ParseContent(b rune) (Content, error) {
	switch b {
	case '.':
		return Empty, nil
	case '@':
		return Paper, nil
	}

	return Unknown, errors.New("failed to parse content")
}

type FloorPlan struct {
	Width  uint
	Height uint

	Contents []Content
}

func NewFloorPlan() FloorPlan {
	return FloorPlan{Contents: make([]Content, 0)}
}

func (f *FloorPlan) ParseInput(input string) error {
	for line := range strings.SplitSeq(input, "\n") {
		for _, b := range line {
			c, err := ParseContent(b)
			if err != nil {
				return err
			}

			f.Contents = append(f.Contents, c)
		}

		f.Width = uint(len(line))
		f.Height++
	}

	return nil
}

func (f FloorPlan) NorthWest(x, y uint) bool {
	if x == 0 || y == 0 {
		return false
	}

	c := f.Contents[((y-1)*f.Width)+x-1]
	return c == Paper || c == PaperMarkedRemoved
}

func (f FloorPlan) North(x, y uint) bool {
	if y == 0 {
		return false
	}

	c := f.Contents[((y-1)*f.Width)+x]
	return c == Paper || c == PaperMarkedRemoved
}

func (f FloorPlan) NorthEast(x, y uint) bool {
	if y == 0 || x == f.Width-1 {
		return false
	}

	c := f.Contents[((y-1)*f.Width)+x+1]
	return c == Paper || c == PaperMarkedRemoved
}

func (f FloorPlan) West(x, y uint) bool {
	if x == 0 {
		return false
	}

	c := f.Contents[(y*f.Width)+x-1]
	return c == Paper || c == PaperMarkedRemoved
}

func (f FloorPlan) East(x, y uint) bool {
	if x == f.Width-1 {
		return false
	}

	c := f.Contents[(y*f.Width)+x+1]
	return c == Paper || c == PaperMarkedRemoved
}

func (f FloorPlan) SouthWest(x, y uint) bool {
	if x == 0 || y == f.Height-1 {
		return false
	}

	c := f.Contents[((y+1)*f.Width)+x-1]
	return c == Paper || c == PaperMarkedRemoved
}

func (f FloorPlan) South(x, y uint) bool {
	if y == f.Height-1 {
		return false
	}

	c := f.Contents[((y+1)*f.Width)+x]
	return c == Paper || c == PaperMarkedRemoved
}

func (f FloorPlan) SouthEast(x, y uint) bool {
	if y == f.Height-1 || x == f.Width-1 {
		return false
	}

	c := f.Contents[((y+1)*f.Width)+x+1]
	return c == Paper || c == PaperMarkedRemoved
}

func (f FloorPlan) CanAccessRoll(x, y uint) bool {
	if f.Contents[y*f.Width+x] != Paper {
		return false
	}

	totalRollsAround := 0

	if f.NorthWest(x, y) {
		totalRollsAround++
	}

	if f.North(x, y) {
		totalRollsAround++
	}

	if f.NorthEast(x, y) {
		totalRollsAround++
	}

	if f.West(x, y) {
		totalRollsAround++
	}

	if f.East(x, y) {
		totalRollsAround++
	}

	if f.SouthWest(x, y) {
		totalRollsAround++
	}

	if f.South(x, y) {
		totalRollsAround++
	}

	if f.SouthEast(x, y) {
		totalRollsAround++
	}

	return totalRollsAround < 4
}

func (f FloorPlan) PartOne() uint {
	var accesibleRolls uint
	for y := range f.Height {
		for x := range f.Width {
			accesible := f.CanAccessRoll(x, y)
			if accesible {
				accesibleRolls++
			}
		}
	}

	return accesibleRolls
}

func (f *FloorPlan) PartTwo(previousRolls uint) uint {
	var accesibleRolls uint
	for y := range f.Height {
		for x := range f.Width {
			accesible := f.CanAccessRoll(x, y)
			if accesible {
				accesibleRolls++
				f.Contents[(y*f.Height)+x] = PaperMarkedRemoved
			}
		}
	}

	if accesibleRolls == 0 {
		return previousRolls
	}

	f.removePaper()
	return f.PartTwo(accesibleRolls + previousRolls)
}

func (f *FloorPlan) removePaper() {
	for i := range f.Contents {
		if f.Contents[i] == PaperMarkedRemoved {
			f.Contents[i] = Empty
		}
	}
}

func main() {
	f, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	fp := NewFloorPlan()
	err = fp.ParseInput(string(f))
	if err != nil {
		log.Fatalf("failed to parse input: %v", err)
	}

	fmt.Printf("Part One: %d\n", fp.PartOne())
	fmt.Printf("Part Two: %d\n", fp.PartTwo(0))
}
