package day4_test

import (
	"advent-of-code-2023/solutions/day4"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func Test_SolvePartA_TestInput(t *testing.T) {
	expected := 13
	mockInput := []byte(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`)

	actual := day4.SolvePartA(mockInput)

	if expected != actual {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}

func Test_SolvePartA_ActualCase(t *testing.T) {
	expected := 21959
	file, err := os.Open(filepath.Join("..", "..", "inputs", "day4.txt"))

	if err != nil {
		t.Fatalf("error opening file: %s", err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("error reading file: %s", err)
	}

	actual := day4.SolvePartA(bytes)

	if expected != actual {
		t.Fatalf("exp %d, act %d", expected, actual)
	}
}
