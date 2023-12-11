package aoc

import (
	"fmt"
)

type point struct {
	x, y int
}

func getGalaxyDists(input []string, multiplier int) int {
	galaxies := []point{}

	rowHasGalaxy := make([]bool, len(input))
	colHasGalaxy := make([]bool, len(input[0]))

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[0]); col++ {
			if input[row][col] == '#' {
				rowHasGalaxy[row] = true
				colHasGalaxy[col] = true
				galaxies = append(galaxies, point{x: col, y: row})
			}
		}
	}

	total := 0

	for i, a := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			b := galaxies[j]

			xStart, xEnd := a.x, b.x
			if a.x > b.x {
				xStart, xEnd = b.x, a.x
			}
			yStart, yEnd := a.y, b.y
			if a.y > b.y {
				yStart, yEnd = b.y, a.y
			}

			galaxyCols := colHasGalaxy[xStart:xEnd]
			galaxyRows := rowHasGalaxy[yStart:yEnd]

			expansions := 0
			for _, b := range galaxyCols {
				if !b {
					expansions++
				}
			}
			for _, b := range galaxyRows {
				if !b {
					expansions++
				}
			}

			dist := abs(b.x-a.x) + abs(b.y-a.y)
			dist += expansions * multiplier

			total += dist
		}
	}

	return total
}

func dayElevenPartOne(input []string) int {
	return getGalaxyDists(input, 1)
}

func dayElevenPartTwo(input []string) int {
	return getGalaxyDists(input, 1_000_000-1)
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func dayElevenTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(11)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			},
			expected: 374,
		},
		"2": {
			input: []string{
				"..........",
				"..........",
				"..........",
				"..........",
				"..........",
				".#........",
				"..........",
				"..........",
				"..........",
				"....#.....",
			},
			expected: 9,
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
