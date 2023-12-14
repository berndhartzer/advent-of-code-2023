package aoc

import (
	"fmt"
)

func dayFourteenPartOne(input []string) int {
	total := 0

	topLoad := len(input)

	col, row, empty := 0, 0, 0
	for {

		switch input[row][col] {
		case 'O':
			load := topLoad - (row - empty)
			total += load
		case '#':
			empty = 0
		case '.':
			empty++
		}

		row++

		if row >= len(input) {
			row = 0
			empty = 0
			col++
		}

		if col >= len(input[0]) {
			break
		}
	}

	return total
}

func dayFourteenPartTwo(input []string) int {
	return 0
}

func dayFourteenTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(14)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: 136,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
