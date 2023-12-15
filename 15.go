package aoc

import (
	"fmt"
)

func dayFifteenPartOne(input string) int {
	total := 0
	subTotal := 0

	for i := 0; i < len(input); i++ {
		if input[i] == ',' {
			total += subTotal
			subTotal = 0
			continue
		}

		subTotal += int(input[i])
		subTotal *= 17
		subTotal %= 256
	}
	total += subTotal

	return total
}

func dayFifteenPartTwo(input string) int {
	total := 0
	return total
}

func dayFifteenTests() (map[string]stringToIntTestConfig, map[string]stringToIntTestConfig, error) {
	fileInput, err := getInput(15)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asString()

	partOne := map[string]stringToIntTestConfig{
		"1": {
			input: "HASH",
			expected: 52,
		},
		"2": {
			input: "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			expected: 1320,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringToIntTestConfig{
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
