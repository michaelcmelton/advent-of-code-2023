package day4

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func readAllInput(scanner *bufio.Scanner) []string {
	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

func SolvePartA(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	fileContents := readAllInput(scanner)
	var answer int

	for i := len(fileContents) - 1; i >= 0; i-- {
		var num int
		line := fileContents[i]
		data := strings.Split(line, ": ")[1]

		winnersStr := strings.Split(strings.Split(data, " | ")[0], " ")
		var winnersInt []int

		for i := range winnersStr {
			if winnersStr[i] == "" {
				continue
			}
			winnerNum, err := strconv.ParseInt(strings.Trim(winnersStr[i], " "), 10, 64)
			if err != nil {
				fmt.Println(err)
				return 0
			}

			winnersInt = append(winnersInt, int(winnerNum))
		}
		cardStr := strings.Split(strings.Split(data, " | ")[1], " ")

		var cardInt []int
		for i := range cardStr {
			if cardStr[i] == "" {
				continue
			}
			cardNum, err := strconv.ParseInt(strings.Trim(cardStr[i], " "), 10, 64)
			if err != nil {
				fmt.Println(err)
				return 0
			}

			cardInt = append(cardInt, int(cardNum))
		}

		for _, winner := range winnersInt {
			if slices.Contains(cardInt, winner) {
				num++
			}
		}
		calc := math.Pow(2, float64(num-1))
		answer += int(calc)
	}
	return answer
}
