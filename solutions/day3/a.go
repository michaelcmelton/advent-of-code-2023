package day3

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

/*

You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?

*/

func locateNeighbors(indexes []int, grid [][]string) [][]int {
	var neighborIndexes [][]int
	offsets := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	re := regexp.MustCompile(`\d`)

	for i := range offsets {
		offset := offsets[i]
		offsetI := indexes[0] + offset[0]
		offsetJ := indexes[1] + offset[1]

		if offsetI >= len(grid) || offsetI < 0 {
			continue
		}

		if offsetJ >= len(grid[indexes[0]]) || offsetJ < 0 {
			continue
		}

		if re.MatchString(grid[offsetI][offsetJ]) {
			neighborIndexes = append(neighborIndexes, []int{offsetI, offsetJ})
		}
	}

	return neighborIndexes
}

func getNumber(idx []int, grid [][]string) int {
	row := grid[idx[0]]
	leftOffset := idx[1] - 1
	rightOffset := idx[1] + 1
	var number int64

	re := regexp.MustCompile(`\D`)

	//look forward for next .
	if leftOffset >= 0 && re.MatchString(row[leftOffset]) {
		return -1
	}

	if row[leftOffset] == "." {
		var ints []string
		for i := idx[1]; i < len(row); i++ {
			if row[i] == "." {
				break
			}

			ints = append(ints, row[i])
		}

		var err error
		number, err = strconv.ParseInt(strings.Join(ints, ""), 10, 64)
		if err != nil {
			fmt.Println(err)
			return -1
		}
	}

	//look backward for next .
	if row[rightOffset] == "." && number == 0 {
		var ints []string
		for i := idx[1]; i >= 0; i-- {
			if row[i] == "." {
				break
			}

			ints = append(ints, row[i])
		}

		slices.Reverse(ints)
		var err error
		number, err = strconv.ParseInt(strings.Join(ints, ""), 10, 64)
		if err != nil {
			fmt.Println(err)
			return -1
		}

	}

	return int(number)
}

func SolvePartA(input []byte) int {
	rows := strings.Split(string(input), "\n")
	var grid [][]string

	numbers := map[int]string{}
	var answer int

	for i := range rows {
		row := rows[i]
		split := strings.Split(row, "")
		grid = append(grid, split)
	}

	symbolRegex := regexp.MustCompile(`[^(\d|\.|\n)]`)

	for i := range grid {
		// iterate the grid
		row := grid[i]

		for j := range row {
			//iterate the row
			cell := row[j]
			// if cell value is a not a match to the regex, skip it
			if !symbolRegex.MatchString(cell) {
				continue
			}

			// gather meaningful neighbors
			cellNeighbors := locateNeighbors([]int{i, j}, grid)

			for k := range cellNeighbors {
				idxSet := cellNeighbors[k]
				number := getNumber(idxSet, grid)
				idxStringSet := []string{strconv.Itoa(idxSet[0]), strconv.Itoa(idxSet[0])}
				if val, ok := numbers[number]; ok {
					if val == idxStringSet[0]+strconv.Itoa(idxSet[1]-2) {
						continue
					}
					numbers[number] = strings.Join(idxStringSet, "")
				} else {
					numbers[number] = strings.Join(idxStringSet, "")
				}
			}
		}
	}

	for gear, _ := range numbers {
		if gear != -1 {
			answer += gear
		}
	}

	return answer
}
