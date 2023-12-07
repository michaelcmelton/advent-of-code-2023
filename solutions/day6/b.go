package day6

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func readInputPartB(input []byte) Race {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var dataExtract []int
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`\S+`)
		matches := re.FindAllString(line, -1)
		joinedNumber := strings.Join(matches[1:], "")
		num, err := strconv.ParseInt(joinedNumber, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		dataExtract = append(dataExtract, int(num))
	}

	return Race{
		Time:     dataExtract[0],
		Distance: dataExtract[1],
	}
}

func SolvePartB(input []byte) int {
	race := readInputPartB(input)

	return race.CalcUpperBound() - race.CalcLowerBound() + 1
}
