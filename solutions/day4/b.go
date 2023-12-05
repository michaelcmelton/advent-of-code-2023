package day4

import (
	"bufio"
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePartB(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	fileContents := readAllInput(scanner)
	var answer int
	pile := make([]int, len(fileContents))

	for l := range pile {
		pile[l] = 1
	}

	for i := 0; i < len(fileContents); i++ {
		line := fileContents[i]
		data := strings.Split(line, ": ")[1]

		winnersStr := strings.Split(strings.Split(data, " | ")[0], " ")
		var winnersInt []int

		for m := range winnersStr {
			if winnersStr[m] == "" {
				continue
			}
			winnerNum, err := strconv.ParseInt(strings.Trim(winnersStr[m], " "), 10, 64)
			if err != nil {
				fmt.Println(err)
				return 0
			}

			winnersInt = append(winnersInt, int(winnerNum))
		}
		cardStr := strings.Split(strings.Split(data, " | ")[1], " ")

		var cardInt []int
		for k := range cardStr {
			if cardStr[k] == "" {
				continue
			}
			cardNum, err := strconv.ParseInt(strings.Trim(cardStr[k], " "), 10, 64)
			if err != nil {
				fmt.Println(err)
				return 0
			}

			cardInt = append(cardInt, int(cardNum))
		}

		var numOfWinners int
		for _, winner := range winnersInt {
			if slices.Contains(cardInt, winner) {
				numOfWinners++
			}
		}

		for j := 1; j <= numOfWinners; j++ {
			cardCpyIdx := i + j
			if cardCpyIdx >= len(fileContents) {
				break
			}

			pile[cardCpyIdx] += 1 * pile[i]

		}
	}

	for n := range pile {
		answer += pile[n]
	}

	return answer
}
