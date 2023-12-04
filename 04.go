package aoc

import (
	"fmt"
	"strings"
)

func dayFourPartOne(input []string) int {
	total := 0

	for _, line := range input {
		split := strings.Split(line, " ")

		winningNums := map[string]bool{}
		checkNums := false
		cardTotal := 0

		for i := 2; i < len(split); i++ {
			if split[i] == "" {
				continue
			}
			if split[i] == "|" {
				checkNums = true
				continue
			}

			if !checkNums {
				winningNums[split[i]] = true
				continue
			}

			_, ok := winningNums[split[i]]
			if ok {
				if cardTotal == 0 {
					cardTotal = 1
				} else {
					cardTotal = cardTotal * 2
				}
			}
		}

		total += cardTotal
	}

	return total
}

func dayFourPartTwo(input []string) int {
	copiesCount := map[int]int{}

	for lineNum, line := range input {
		split := strings.Split(line, " ")

		winningNums := map[string]bool{}
		checkNums := false
		cardTotal := 0

		for i := 2; i < len(split); i++ {
			if split[i] == "" {
				continue
			}
			if split[i] == "|" {
				checkNums = true
				continue
			}

			if !checkNums {
				winningNums[split[i]] = true
				continue
			}

			_, ok := winningNums[split[i]]
			if ok {
				cardTotal++
			}
		}

		if cardTotal == 0 {
			continue
		}

		// Check copies we have of this card
		cardsToProcess, ok := copiesCount[lineNum]
		if !ok {
			cardsToProcess = 0
		}

		// Add the original card for processing
		cardsToProcess++

		// Process all the cards, our original and the copies
		for k := 0; k < cardsToProcess; k++ {
			copyIdx := lineNum + 1
			for j := 0; j < cardTotal; j++ {
				copiesCount[copyIdx]++
				copyIdx++
			}
		}
	}

	total := len(input)
	for _, v := range copiesCount {
		total += v
	}

	return total
}

type dayFourTestConfig struct {
	input     []string
	expected  int
	logResult bool
}

func getDayFourTests() (map[string]dayFourTestConfig, map[string]dayFourTestConfig, error) {
	fileInput, err := getInput(4)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]dayFourTestConfig{
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]dayFourTestConfig{
		"1": {
			input: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			expected: 30,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
