package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func dayTwoPartOne(input []string) int {
	redLimit, greenLimit, blueLimit := 12, 13, 14

	total := 0

	for idx, line := range input {
		valid := true

		split := strings.Split(line, " ")
		for i := 2; i < len(split); i += 2 {
			n, err := strconv.Atoi(split[i])
			if err != nil {
				panic("failed to convert string to num")
			}

			switch {
			case strings.HasPrefix(split[i+1], "red"):
				if n > redLimit {
					valid = false
				}
			case strings.HasPrefix(split[i+1], "green"):
				if n > greenLimit {
					valid = false
				}
			case strings.HasPrefix(split[i+1], "blue"):
				if n > blueLimit {
					valid = false
				}
			}

			if !valid {
				break
			}
		}

		if valid {
			total += idx + 1
		}
	}

	return total
}

func dayTwoPartTwo(input []string) int {
	total := 0

	for _, line := range input {
		minRed, minGreen, minBlue := -1, -1, -1

		split := strings.Split(line, " ")
		for i := 2; i < len(split); i += 2 {
			n, err := strconv.Atoi(split[i])
			if err != nil {
				panic("failed to convert string to num")
			}

			switch {
			case strings.HasPrefix(split[i+1], "red"):
				if minRed == -1 || n > minRed {
					minRed = n
				}
			case strings.HasPrefix(split[i+1], "green"):
				if minGreen == -1 || n > minGreen {
					minGreen = n
				}
			case strings.HasPrefix(split[i+1], "blue"):
				if minBlue == -1 || n > minBlue {
					minBlue = n
				}
			}
		}

		total += (minRed * minGreen * minBlue)
	}

	return total
}

type dayTwoTestConfig struct {
	input     []string
	expected  int
	logResult bool
}

func getDayTwoTests() (map[string]dayTwoTestConfig, map[string]dayTwoTestConfig, error) {
	fileInput, err := getInput(2)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]dayTwoTestConfig{
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]dayTwoTestConfig{
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
