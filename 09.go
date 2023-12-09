package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func dayNine(input []string, nextValueFn func(i int, diffs [][]int) int) int {
	total := 0

	for _, line := range input {
		split := strings.Split(line, " ")

		diffs := [][]int{}
		thisDiff := []int{}

		for _, s := range split {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic("failed to convert string to num")
			}
			thisDiff = append(thisDiff, n)
		}

		diffs = append(diffs, thisDiff)
		thisDiff = []int{}

		for d := 0; ; d++ {
			allZero := true

			prevNum := 0
			for i, n := range diffs[d] {
				if i == 0 {
					prevNum = n
					continue
				}

				diff := n - prevNum
				thisDiff = append(thisDiff, diff)
				prevNum = n

				if diff != 0 {
					allZero = false
				}
			}

			diffs = append(diffs, thisDiff)
			thisDiff = []int{}

			if allZero {
				break
			}
		}

		for i := len(diffs) - 1; i >= 0; i-- {
			if i == len(diffs)-1 {
				diffs[i] = append(diffs[i], 0)
				continue
			}

			nextValue := nextValueFn(i, diffs)

			diffs[i] = append(diffs[i], nextValue)
		}

		total += diffs[0][len(diffs[0])-1]
	}

	return total
}

func dayNinePartOne(input []string) int {
	nextValueFn := func(i int, diffs [][]int) int {
		left := diffs[i][len(diffs[i])-1]
		below := diffs[i+1][len(diffs[i+1])-1]
		return left + below
	}

	return dayNine(input, nextValueFn)
}

func dayNinePartTwo(input []string) int {
	nextValueFn := func(i int, diffs [][]int) int {
		right := diffs[i][0]

		// we'll keep using the end of the array to store the new values
		below := diffs[i+1][len(diffs[i+1])-1]
		return right - below
	}

	return dayNine(input, nextValueFn)
}

type dayNineTestConfig struct {
	input     []string
	expected  int
	logResult bool
}

func getDayNineTests() (map[string]dayNineTestConfig, map[string]dayNineTestConfig, error) {
	fileInput, err := getInput(9)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]dayNineTestConfig{
		"1": {
			input: []string{
				"0 3 6 9 12 15",
				"1 3 6 10 15 21",
				"10 13 16 21 30 45",
			},
			expected: 114,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]dayNineTestConfig{
		"1": {
			input: []string{
				"0 3 6 9 12 15",
				"1 3 6 10 15 21",
				"10 13 16 21 30 45",
			},
			expected: 2,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
