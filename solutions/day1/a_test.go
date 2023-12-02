package day1_test

import (
	"advent-of-code-2023/solutions/day1"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func Test_SolvePartA_TestCase(t *testing.T) {
	expected := 142
	mockInput := `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
	`

	actual := day1.SolvePartA(mockInput)

	if expected != actual {
		t.Fatalf("exected %d, actual %d", expected, actual)
	}
}

func Test_SolvePartA_ActualCase(t *testing.T) {
	expected := 54953
	file, err := os.Open(filepath.Join("..", "..", "inputs", "day1.txt"))

	if err != nil {
		t.Fatalf("error opening file: %s", err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("error reading file: %s", err)
	}

	actual := day1.SolvePartA(string(bytes))

	if expected != actual {
		t.Fatalf("exected %d, actual %d", expected, actual)
	}
}
