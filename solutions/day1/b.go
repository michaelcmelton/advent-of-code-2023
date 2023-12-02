package day1

import (
	"fmt"
	"regexp"
)

func SolvePartB(input string) int {
	substitutedInput := input
	substitutions := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}

	for key, val := range substitutions {
		re := regexp.MustCompile(key)
		substitutedInput = re.ReplaceAllString(substitutedInput, val)
	}

	fmt.Printf("substituted input:\n\n%s\n", substitutedInput)

	return SolvePartA(input)
}
