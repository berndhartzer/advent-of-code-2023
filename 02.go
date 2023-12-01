package aoc

import (
	"fmt"
)

func dayTwoPartOne(input []string) int {
	return 0
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
