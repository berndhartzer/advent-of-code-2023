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

type lens struct {
	label       string
	focalLength byte
}

func dayFifteenPartTwo(input string) int {
	boxes := make([][]lens, 256)

	labelStart, labelEnd, hash := 0, 0, 0
	label := ""
	var operation byte
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case ',':
			labelStart = i + 1
			continue
		case '=', '-':
			labelEnd = i

			operation = input[i]
			label = input[labelStart:labelEnd]
			hash = dayFifteenPartOne(input[labelStart:labelEnd])
		default:
			continue
		}

		switch operation {
		case '=':
			newLens := lens{
				label:       label,
				focalLength: input[i+1],
			}

			replaced := false
			for i, lens := range boxes[hash] {
				if lens.label == label {
					replaced = true
					boxes[hash][i] = newLens
				}
			}

			if !replaced {
				boxes[hash] = append(boxes[hash], newLens)
			}

		case '-':
			i := 0
			for _, lens := range boxes[hash] {
				if lens.label != label {
					boxes[hash][i] = lens
					i++
				}
			}
			boxes[hash] = boxes[hash][:i]
		}
	}

	total := 0
	for i, box := range boxes {
		for j, lens := range box {
			score := (i + 1) * (j + 1) * (int(lens.focalLength) - 48)
			total += score
		}
	}

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
			input:    "HASH",
			expected: 52,
		},
		"2": {
			input:    "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			expected: 1320,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringToIntTestConfig{
		"1": {
			input:    "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			expected: 145,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
