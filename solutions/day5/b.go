package day5

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"time"
)

func SolvePartB(input []byte) int {
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
			var interim []int
			for i := range seedsMatches {
				seed := seedsMatches[i]
				interim = append(interim, convertToInt(seed))
				if len(interim) == 2 {
					seeds = append(seeds, interim)
					interim = []int{}
				}
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
		seedRange := seeds[i][1]
		fmt.Printf("%#v processing seed pair %d of %d... %#v\n", time.Now().Format(time.DateTime), i, len(seeds), seeds[i])

		for j := 0; j < seedRange; j++ {
			seedNumber := seeds[i][0] + j
			fmt.Printf("%#v processing seed number %d, starting value: %d, ending value: %d\n", time.Now().Format(time.DateTime), seedNumber, seeds[i][0], seeds[i][0]+seeds[i][1]-1)
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
	}

	return answer
}
