package day3_test

import (
	"advent-of-code-2023/solutions/day3"
	"testing"
)

func Test_PartA_ExampleInput(t *testing.T) {
	expected := 4361
	mockInput := []byte(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

	actual := day3.SolvePartA(mockInput)

	if expected != actual {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}
