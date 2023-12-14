package aoc

import (
	"encoding/json"
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
	grid := make([][]byte, len(input))

	for row, line := range input {
		gridRow := make([]byte, len(line))
		for i, c := range line {
			switch c {
			case 'O', '#':
				gridRow[i] = byte(c)
			}
		}
		grid[row] = gridRow
	}

	tiltNorth := func() {
		col, row, empty := 0, 0, 0
		for {
			switch grid[row][col] {
			case 'O':
				if empty > 0 {
					grid[row-empty][col] = 'O'
					grid[row][col] = byte(0)
				}
			case '#':
				empty = 0
			default:
				empty++
			}

			row++

			if row >= len(grid) {
				row = 0
				empty = 0
				col++
			}

			if col >= len(grid[0]) {
				break
			}
		}
	}

	tiltWest := func() {
		for row := 0; row < len(grid); row++ {
			empty := 0
			for col := 0; col < len(grid[0]); col++ {
				switch grid[row][col] {
				case 'O':
					if empty > 0 {
						grid[row][col-empty] = 'O'
						grid[row][col] = byte(0)
					}
				case '#':
					empty = 0
				default:
					empty++
				}
			}
		}
	}

	tiltSouth := func() {
		col, row, empty := 0, len(grid)-1, 0
		for {
			switch grid[row][col] {
			case 'O':
				if empty > 0 {
					grid[row+empty][col] = 'O'
					grid[row][col] = byte(0)
				}
			case '#':
				empty = 0
			default:
				empty++
			}

			row--

			if row < 0 {
				row = len(grid) - 1
				empty = 0
				col++
			}

			if col >= len(grid[0]) {
				break
			}
		}
	}

	tiltEast := func() {
		for row := 0; row < len(grid); row++ {
			empty := 0
			for col := len(grid[0]) - 1; col >= 0; col-- {
				switch grid[row][col] {
				case 'O':
					if empty > 0 {
						grid[row][col+empty] = 'O'
						grid[row][col] = byte(0)
					}
				case '#':
					empty = 0
				default:
					empty++
				}
			}
		}
	}

	cycle := func() {
		tiltNorth()
		tiltWest()
		tiltSouth()
		tiltEast()
	}

	cache := map[string]int{}
	breakpoint := -1

	for i := 0; i < 1_000_000_000; i++ {
		cycle()

		state, err := json.Marshal(grid)
		if err != nil {
			panic("failed to marshal grid")
		}

		if breakpoint == -1 {
			v, ok := cache[string(state)]
			if ok {
				breakpoint = i + (1_000_000_000-v)%(i-v) - 1
				continue
			}

			cache[string(state)] = i
		}

		if i != breakpoint {
			continue
		}

		break
	}

	// Score the grid
	total := 0

	topLoad := len(grid)

	col, row, empty := 0, 0, 0
	for {
		switch grid[row][col] {
		case 'O':
			load := topLoad - row
			total += load
		case '#':
			empty = 0
		default:
			empty++
		}

		row++

		if row >= len(grid) {
			row = 0
			empty = 0
			col++
		}

		if col >= len(grid[0]) {
			break
		}
	}

	return total
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
			expected: 64,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
