package aoc

import (
	"fmt"
	"strings"
)

func dayOnePartOne(input []string) int {
	total := 0

	for _, line := range input {
		l := 0
		r := len(line) - 1

		gotFirst, gotLast := false, false
		var first, last byte

		for {
			// 49 == 1, 57 == 9
			if !gotFirst && line[l] >= 49 && line[l] <= 57 {
				first = line[l]
				gotFirst = true
			}
			if !gotLast && line[r] >= 49 && line[r] <= 57 {
				last = line[r]
				gotLast = true
			}
			if gotFirst && gotLast {
				break
			}

			l++
			r--
		}

		// convert byte to int value by subtracting 48
		// e.g. 49 - 48 == 1
		total += (int(first) - 48) * 10
		total += (int(last) - 48)
	}

	return total
}

func dayOnePartTwo(input []string) int {
	total := 0

	for _, line := range input {
		l := 0
		r := len(line) - 1

		gotFirst, gotLast := false, false
		first, last := 0, 0

		for {
			// 49 == 1, 57 == 9
			if !gotFirst && line[l] >= 49 && line[l] <= 57 {
				first = int(line[l]) - 48
				gotFirst = true
			}
			if !gotLast && line[r] >= 49 && line[r] <= 57 {
				last = int(line[r]) - 48
				gotLast = true
			}

			if !gotFirst {
				tmpStr := line[l:]

				switch {
				case strings.HasPrefix(tmpStr, "one"):
					first = 1
					gotFirst = true
				case strings.HasPrefix(tmpStr, "two"):
					first = 2
					gotFirst = true
				case strings.HasPrefix(tmpStr, "three"):
					first = 3
					gotFirst = true
				case strings.HasPrefix(tmpStr, "four"):
					first = 4
					gotFirst = true
				case strings.HasPrefix(tmpStr, "five"):
					first = 5
					gotFirst = true
				case strings.HasPrefix(tmpStr, "six"):
					first = 6
					gotFirst = true
				case strings.HasPrefix(tmpStr, "seven"):
					first = 7
					gotFirst = true
				case strings.HasPrefix(tmpStr, "eight"):
					first = 8
					gotFirst = true
				case strings.HasPrefix(tmpStr, "nine"):
					first = 9
					gotFirst = true
				}
			}
			if !gotLast {
				tmpStr := line[0 : r+1]

				switch {
				case strings.HasSuffix(tmpStr, "one"):
					last = 1
					gotLast = true
				case strings.HasSuffix(tmpStr, "two"):
					last = 2
					gotLast = true
				case strings.HasSuffix(tmpStr, "three"):
					last = 3
					gotLast = true
				case strings.HasSuffix(tmpStr, "four"):
					last = 4
					gotLast = true
				case strings.HasSuffix(tmpStr, "five"):
					last = 5
					gotLast = true
				case strings.HasSuffix(tmpStr, "six"):
					last = 6
					gotLast = true
				case strings.HasSuffix(tmpStr, "seven"):
					last = 7
					gotLast = true
				case strings.HasSuffix(tmpStr, "eight"):
					last = 8
					gotLast = true
				case strings.HasSuffix(tmpStr, "nine"):
					last = 9
					gotLast = true
				}
			}

			if gotFirst && gotLast {
				break
			}

			l++
			r--
		}

		total += first * 10
		total += last
	}

	return total
}

type dayOneTestConfig struct {
	input     []string
	expected  int
	logResult bool
}

func getDayOneTests() (map[string]dayOneTestConfig, map[string]dayOneTestConfig, error) {
	fileInput, err := getInput(1)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]dayOneTestConfig{
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]dayOneTestConfig{
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
