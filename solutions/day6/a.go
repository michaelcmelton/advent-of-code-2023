package day6

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

const SPEED_INCREMENT = 1

type Race struct {
	Time     int
	Distance int
}

func (r Race) CalcLowerBound() int {
	endBound := calcBound(r.Time)

	for i := 0; i < endBound; i++ {
		timeRemaining := r.Time - i
		speed := i * SPEED_INCREMENT
		if speed*timeRemaining > r.Distance {
			return i
		}
	}
	return -1
}

func calcBound(maxTime int) int {
	time := maxTime
	if maxTime%2 != 0 {
		time = maxTime + 1
	}

	endBound := time / 2
	if maxTime%2 != 0 {
		endBound -= 1
	}

	return endBound
}

func (r Race) CalcUpperBound() int {
	endBound := calcBound(r.Time)
	for i := r.Time; i >= endBound; i-- {
		timeRemaining := r.Time - i
		speed := i * SPEED_INCREMENT
		if speed*timeRemaining > r.Distance {
			return i
		}
	}
	return -1
}

func readInput(input []byte) []Race {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var races []Race
	var dataExtract [][]int
	for scanner.Scan() {
		var lineExtract []int
		line := scanner.Text()
		re := regexp.MustCompile(`\S+`)
		matches := re.FindAllString(line, -1)
		for i := range matches {
			if i == 0 {
				continue
			}
			num, err := strconv.ParseInt(matches[i], 10, 64)
			if err != nil {
				fmt.Println(err)
				break
			}

			lineExtract = append(lineExtract, int(num))
		}

		dataExtract = append(dataExtract, lineExtract)
	}

	for i := range dataExtract[0] {
		races = append(races, Race{
			Time:     dataExtract[0][i],
			Distance: dataExtract[1][i],
		})
	}

	return races
}

func SolvePartA(input []byte) int {
	answer := 1
	dataSet := readInput(input)

	for i := range dataSet {
		race := dataSet[i]
		answer *= race.CalcUpperBound() - race.CalcLowerBound() + 1
	}
	return answer
}
