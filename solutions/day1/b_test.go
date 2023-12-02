package day1_test

import (
	"advent-of-code-2023/solutions/day1"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func Test_SolvePartB_TestCase(t *testing.T) {
	expected := 281
	mockInput := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
	`

	actual := day1.SolvePartB(mockInput)

	if expected != actual {
		t.Fatalf("exected %d, actual %d", expected, actual)
	}
}

func Test_SolvePartB_ActualCase(t *testing.T) {
	expected := 53868
	file, err := os.Open(filepath.Join("..", "..", "inputs", "day1.txt"))

	if err != nil {
		t.Fatalf("error opening file: %s", err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("error reading file: %s", err)
	}

	actual := day1.SolvePartB(string(bytes))

	if expected != actual {
		t.Fatalf("exected %d, actual %d", expected, actual)
	}
}
