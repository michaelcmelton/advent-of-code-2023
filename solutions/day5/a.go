package day5

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
)

func readAllInput(scanner *bufio.Scanner) []string {
	var textInput []string
	for scanner.Scan() {
		textInput = append(textInput, scanner.Text())
	}

	return textInput
}

func convertToInt(number string) int {
	re := regexp.MustCompile(`\d+`)
	if re.MatchString(number) {
		ret, _ := strconv.ParseInt(number, 10, 64)
		return int(ret)
	}

	return -1
}

func addToMap(line string, dataMap [][]int) [][]int {
	re := regexp.MustCompile(`\d+`)
	intMatches := re.FindAllString(line, -1)
	if intMatches == nil {
		return dataMap
	}

	var lineSlice []int

	for _, intMatch := range intMatches {
		lineSlice = append(lineSlice, convertToInt(intMatch))
	}

	dataMap = append(dataMap, lineSlice)

	return dataMap
}

func getNextValue(input int, dataMap [][]int) int {
	var nextVal int
	for i := range dataMap {
		set := dataMap[i]
		maxVal := set[1] + set[2]
		minVal := set[1]

		if input < minVal || input > maxVal {
			continue
		}

		diff := input - set[1]
		nextVal = set[0] + diff
	}

	if nextVal == 0 {
		return input
	}

	return nextVal
}

func SolvePartA(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	fileContentsByLine := readAllInput(scanner)

	var answer int
	var seeds [][]int
	var seedSoilMap [][]int
	var soilFertMap [][]int
	var fertWaterMap [][]int
	var waterLightMap [][]int
	var lightTempMap [][]int
	var tempHumidityMap [][]int
	var humidityLocationMap [][]int
	var breakCount int
	for i := range fileContentsByLine {
		line := fileContentsByLine[i]

		if line == "" {
			breakCount++
			continue
		}

		if i == 0 {
			re := regexp.MustCompile(`\d+`)
			seedsMatches := re.FindAllString(line, -1)
			for _, seed := range seedsMatches {
				seeds = append(seeds, []int{convertToInt(seed)})
			}
		}

		switch breakCount {
		case 1:
			seedSoilMap = addToMap(line, seedSoilMap)
		case 2:
			soilFertMap = addToMap(line, soilFertMap)
		case 3:
			fertWaterMap = addToMap(line, fertWaterMap)
		case 4:
			waterLightMap = addToMap(line, waterLightMap)
		case 5:
			lightTempMap = addToMap(line, lightTempMap)
		case 6:
			tempHumidityMap = addToMap(line, tempHumidityMap)
		case 7:
			humidityLocationMap = addToMap(line, humidityLocationMap)
		}
	}

	for i := range seeds {
		seedNumber := seeds[i][0]
		soilNumber := getNextValue(seedNumber, seedSoilMap)
		fertNumber := getNextValue(soilNumber, soilFertMap)
		waterNumber := getNextValue(fertNumber, fertWaterMap)
		lightNumber := getNextValue(waterNumber, waterLightMap)
		tempNumber := getNextValue(lightNumber, lightTempMap)
		humidityNumber := getNextValue(tempNumber, tempHumidityMap)
		locationNumber := getNextValue(humidityNumber, humidityLocationMap)

		if answer == 0 || locationNumber < answer {
			answer = locationNumber
		}
	}

	return answer
}
