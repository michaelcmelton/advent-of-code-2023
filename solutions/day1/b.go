package day1

import (
	"regexp"
)

func SolvePartB(input string) int {
	substitutedInput := input
	substitutions := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5v",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
		"zero":  "z0o",
	}

	for key, val := range substitutions {
		re := regexp.MustCompile(key)
		substitutedInput = re.ReplaceAllString(substitutedInput, val)
	}

	return SolvePartA([]byte(substitutedInput))
}
