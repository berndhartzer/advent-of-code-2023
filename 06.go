package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func getRaceScores(times, dists []int) int {
	total := 1

	for race := 0; race < len(times); race++ {
		winCount := 0

		for hold := 0; hold < times[race]; hold++ {
			remaining := times[race] - hold
			travelled := remaining * hold

			if travelled > dists[race] {
				winCount++
			}
		}

		total *= winCount
	}

	return total
}

func daySixPartOne(input []string) int {
	var times, dists []int

	strTimes := getNumbersFromString(input[0])
	strDists := getNumbersFromString(input[1])

	for i := 0; i < len(strTimes); i++ {
		n, err := strconv.Atoi(strTimes[i])
		if err != nil {
			panic("failed to convert string to num")
		}
		times = append(times, n)

		n, err = strconv.Atoi(strDists[i])
		if err != nil {
			panic("failed to convert string to num")
		}
		dists = append(dists, n)
	}

	return getRaceScores(times, dists)
}

func daySixPartTwo(input []string) int {
	strTimes := getNumbersFromString(input[0])
	strDists := getNumbersFromString(input[1])

	nTime, err := strconv.Atoi(strings.Join(strTimes, ""))
	if err != nil {
		panic("failed to convert string to num")
	}

	nDist, err := strconv.Atoi(strings.Join(strDists, ""))
	if err != nil {
		panic("failed to convert string to num")
	}

	return getRaceScores([]int{nTime}, []int{nDist})
}

type daySixTestConfig struct {
	input     []string
	expected  int
	logResult bool
}

func getDaySixTests() (map[string]daySixTestConfig, map[string]daySixTestConfig, error) {
	fileInput, err := getInput(6)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]daySixTestConfig{
		"1": {
			input: []string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			expected: 288,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]daySixTestConfig{
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
