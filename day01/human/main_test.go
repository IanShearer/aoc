package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestPartOneSample(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	lock := NewLock()
	for scanner.Scan() {
		err := lock.Turn(scanner.Text())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if lock.TimesAtZero() != 3 {
		t.Fatalf("unexepected result.\n\nExpecting: 3\nGot: %d", lock.TimesAtZero())
	}
}

func TestPartTwoSample(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	lock := NewLock()
	for scanner.Scan() {
		err := lock.Turn(scanner.Text())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if lock.NumberOfClicks() != 6 {
		t.Fatalf("unexepected result.\n\nExpecting: 6\nGot: %d", lock.NumberOfClicks())
	}
}
