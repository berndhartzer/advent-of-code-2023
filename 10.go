package aoc

import (
	"fmt"
)

func dayTenPartOne(input []string) int {
	// U, R, D, L
	pipeOptions := map[byte][4]byte{
		'|': {1, 0, 1, 0},
		'-': {0, 1, 0, 1},
		'L': {1, 1, 0, 0},
		'J': {1, 0, 0, 1},
		'7': {0, 0, 1, 1},
		'F': {0, 1, 1, 0},

		// Special case for start point
		'S': {1, 1, 1, 1},
	}

	x, y := 0, 0
	xPrev, yPrev := 0, 0

	// first find s
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[0]); col++ {
			if input[row][col] == 'S' {
				x, xPrev = col, col
				y, yPrev = row, row
				break
			}
		}
	}

	steps := 0
	currPipe := input[y][x]

	for {
		currPipe = input[y][x]
		if currPipe == 'S' && steps > 0 {
			break
		}

		options := pipeOptions[currPipe]

		if y > 0 {
			up := input[y-1][x]
			if up == '|' || up == '7' || up == 'F' || up == 'S' {
				// can go up
				options[0] = options[0] & byte(1)
			} else {
				// We need to exclude these options for the
				// initial starting S position
				options[0] = byte(0)
			}
		}
		if x < len(input[0])-1 {
			right := input[y][x+1]
			if right == '-' || right == 'J' || right == '7' || right == 'S' {
				// can go right
				options[1] = options[1] & byte(1)
			} else {
				options[1] = byte(0)
			}
		}
		if y < len(input)-1 {
			down := input[y+1][x]
			if down == '|' || down == 'L' || down == 'J' || down == 'S' {
				// can go down
				options[2] = options[2] & byte(1)
			} else {
				options[2] = byte(0)
			}
		}
		if x > 0 {
			left := input[y][x-1]
			if left == '-' || left == 'L' || left == 'F' || left == 'S' {
				// can go left
				options[3] = options[3] & byte(1)
			} else {
				options[3] = byte(0)
			}
		}

		// Get the first direction we can go which isn't backwards
		xNew, yNew := func() (int, int) {
			xN, yN := 0, 0

			if options[0] == byte(1) {
				xN, yN = x, y-1
				if xN != xPrev || yN != yPrev {
					return xN, yN
				}
			}

			if options[1] == byte(1) {
				xN, yN = x+1, y
				if xN != xPrev || yN != yPrev {
					return xN, yN
				}
			}

			if options[2] == byte(1) {
				xN, yN = x, y+1
				if xN != xPrev || yN != yPrev {
					return xN, yN
				}
			}

			if options[3] == byte(1) {
				xN, yN = x-1, y
				if xN != xPrev || yN != yPrev {
					return xN, yN
				}
			}

			return -1, -1
		}()

		xPrev, yPrev = x, y
		x, y = xNew, yNew

		steps++
	}

	return steps / 2
}

func dayTenPartTwo(input []string) int {
	return 0
}

func dayTenTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(10)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			},
			expected: 4,
		},
		"2": {
			input: []string{
				// "..F7.",
				// ".FJ|.",
				// "SJ.L7",
				// "|F--J",
				// "LJ...",

				"7-F7-",
				".FJ|7",
				"SJLL7",
				"|F--J",
				"LJ.LJ",
			},
			expected: 8,
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
