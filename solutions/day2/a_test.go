package day2_test

import (
	"advent-of-code-2023/solutions/day2"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func Test_SolvePartA_TestCase(t *testing.T) {
	expected := 8
	mockInput := []byte(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
	`)

	actual := day2.SolvePartA(mockInput, 12, 14, 13)

	if expected != actual {
		t.Fatalf("exp %d, act %d", expected, actual)
	}
}

func Test_SolvePartA_ActualCase(t *testing.T) {
	expected := 2164

	file, err := os.Open(filepath.Join("..", "..", "inputs", "day2.txt"))

	if err != nil {
		t.Fatalf("error opening file: %s", err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("error reading file: %s", err)
	}

	actual := day2.SolvePartA(bytes, 12, 14, 13)

	if expected != actual {
		t.Fatalf("exp %d, act %d", expected, actual)
	}
}
