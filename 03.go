package aoc

import (
	"fmt"
	"strconv"
)

func dayThreePartOne(input []string) int {
	total := 0
	rowLen := len(input[0])

	for row := 0; row < len(input); row++ {
		inNum := false
		checkNum := false
		startNum, endNum := 0, 0

		for col := 0; col < rowLen; col++ {
			switch input[row][col] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if !inNum {
					inNum = true
					startNum = col
				}
			default:
				if inNum {
					inNum = false
					endNum = col - 1
					checkNum = true
				}
			}

			// we're in a number and hit the end of the row
			if inNum && col+1 >= rowLen {
				inNum = false
				endNum = col
				checkNum = true
			}

			if !checkNum {
				continue
			}

			valid := true

			isTopRow := row-1 < 0
			isBottomRow := row+1 >= len(input)
			isLeftCol := startNum-1 < 0
			isRightCol := endNum+1 > rowLen-1

			valid = func() bool {
				if !isTopRow {
					for c := startNum; c < endNum+1; c++ {
						if isSymbol(input[row-1][c]) {
							return true
						}
					}
				}

				if !isTopRow && !isLeftCol {
					if isSymbol(input[row-1][startNum-1]) {
						return true
					}
				}

				if !isTopRow && !isRightCol {
					if isSymbol(input[row-1][endNum+1]) {
						return true
					}
				}

				if !isRightCol {
					if isSymbol(input[row][endNum+1]) {
						return true
					}
				}

				if !isLeftCol {
					if isSymbol(input[row][startNum-1]) {
						return true
					}
				}

				if !isBottomRow {
					for c := startNum; c < endNum+1; c++ {
						if isSymbol(input[row+1][c]) {
							return true
						}
					}
				}

				if !isBottomRow && !isLeftCol {
					if isSymbol(input[row+1][startNum-1]) {
						return true
					}
				}

				if !isBottomRow && !isRightCol {
					if isSymbol(input[row+1][endNum+1]) {
						return true
					}
				}

				return false
			}()

			if valid {
				n, err := strconv.Atoi(string(input[row][startNum : endNum+1]))
				if err != nil {
					panic("failed to convert string to num")
				}
				total += n
			}

			checkNum = false
		}
	}

	return total
}

func isSymbol(c byte) bool {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
		return false
	}
	return true
}

func dayThreePartTwo(input []string) int {
	rowLen := len(input[0])

	gears := map[string][]string{}

	for row := 0; row < len(input); row++ {
		inNum := false
		checkNum := false
		startNum, endNum := 0, 0

		for col := 0; col < rowLen; col++ {
			switch input[row][col] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if !inNum {
					inNum = true
					startNum = col
				}
			default:
				if inNum {
					inNum = false
					endNum = col - 1
					checkNum = true
				}
			}

			if inNum && col+1 >= rowLen {
				inNum = false
				endNum = col
				checkNum = true
			}

			if !checkNum {
				continue
			}

			isTopRow := row-1 < 0
			isBottomRow := row+1 >= len(input)
			isLeftCol := startNum-1 < 0
			isRightCol := endNum+1 > rowLen-1

			stringNum := string(input[row][startNum : endNum+1])

			if !isTopRow {
				for c := startNum; c < endNum+1; c++ {
					if input[row-1][c] == '*' {
						gKey := fmt.Sprintf("%d,%d", row-1, c)
						gears[gKey] = append(gears[gKey], stringNum)
					}
				}
			}

			if !isTopRow && !isLeftCol {
				if input[row-1][startNum-1] == '*' {
					gKey := fmt.Sprintf("%d,%d", row-1, startNum-1)
					gears[gKey] = append(gears[gKey], stringNum)
				}
			}

			if !isTopRow && !isRightCol {
				if input[row-1][endNum+1] == '*' {
					gKey := fmt.Sprintf("%d,%d", row-1, endNum+1)
					gears[gKey] = append(gears[gKey], stringNum)
				}
			}

			if !isRightCol {
				if input[row][endNum+1] == '*' {
					gKey := fmt.Sprintf("%d,%d", row, endNum+1)
					gears[gKey] = append(gears[gKey], stringNum)
				}
			}

			if !isLeftCol {
				if input[row][startNum-1] == '*' {
					gKey := fmt.Sprintf("%d,%d", row, startNum-1)
					gears[gKey] = append(gears[gKey], stringNum)
				}
			}

			if !isBottomRow {
				for c := startNum; c < endNum+1; c++ {
					if input[row+1][c] == '*' {
						gKey := fmt.Sprintf("%d,%d", row+1, c)
						gears[gKey] = append(gears[gKey], stringNum)
					}
				}
			}

			if !isBottomRow && !isLeftCol {
				if input[row+1][startNum-1] == '*' {
					gKey := fmt.Sprintf("%d,%d", row+1, startNum-1)
					gears[gKey] = append(gears[gKey], stringNum)
				}
			}

			if !isBottomRow && !isRightCol {
				if input[row+1][endNum+1] == '*' {
					gKey := fmt.Sprintf("%d,%d", row+1, endNum+1)
					gears[gKey] = append(gears[gKey], stringNum)
				}
			}

			checkNum = false
		}
	}

	total := 0
	for _, v := range gears {
		if len(v) != 2 {
			continue
		}

		ratio := 1
		for _, s := range v {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic("failed to convert string to num")
			}
			ratio *= n
		}

		total += ratio
	}

	return total
}

func dayThreeTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(3)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 4361,
		},
		"2": {
			input: []string{
				"467..114..",
				"...*......",
				"..35...633",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 4361,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 467835,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
