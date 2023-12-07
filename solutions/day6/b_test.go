package day6_test

import (
	"advent-of-code-2023/solutions/day6"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func Test_SolvePartB_SampleInput(t *testing.T) {
	mockInput := []byte(`Time:      7  15   30
Distance:  9  40  200`)

	expected := 71503
	actual := day6.SolvePartB(mockInput)

	if expected != actual {
		t.Fatalf("exp %d, got %d", expected, actual)
	}
}

func Test_SolvePartB_ActualCase(t *testing.T) {
	expected := 30125202
	file, err := os.Open(filepath.Join("..", "..", "inputs", "day6.txt"))

	if err != nil {
		t.Fatalf("error opening file: %s", err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("error reading file: %s", err)
	}

	actual := day6.SolvePartB(bytes)

	if expected != actual {
		t.Fatalf("exp %d, act %d", expected, actual)
	}
}
