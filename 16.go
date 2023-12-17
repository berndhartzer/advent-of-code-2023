package aoc

import (
	"fmt"
)

type beam struct {
	x, y, xDir, yDir int
}

func (b *beam) inBounds(xBound, yBound int) bool {
	if b.x < 0 {
		return false
	}
	if b.y < 0 {
		return false
	}
	if b.x >= xBound {
		return false
	}
	if b.y >= yBound {
		return false
	}

	return true
}

func (b *beam) move() {
	b.x += b.xDir
	b.y += b.yDir
}

func (b *beam) split(splitter byte) *beam {
	if splitter == '|' && b.yDir != 0 {
		b.move()
		return nil
	}
	if splitter == '-' && b.xDir != 0 {
		b.move()
		return nil
	}

	if splitter == '|' && b.xDir != 0 {
		b.xDir = 0
		b.yDir = -1

		newBeam := &beam{
			x:    b.x,
			y:    b.y,
			xDir: 0,
			yDir: 1,
		}

		b.move()
		newBeam.move()

		return newBeam
	}

	if splitter == '-' && b.yDir != 0 {
		b.yDir = 0
		b.xDir = -1

		newBeam := &beam{
			x:    b.x,
			y:    b.y,
			xDir: 1,
			yDir: 0,
		}

		b.move()
		newBeam.move()

		return newBeam
	}

	return nil
}

func (b *beam) mirror(mirror byte) {
	if mirror == '/' {
		if b.xDir == -1 {
			b.xDir = 0
			b.yDir = 1
		} else if b.xDir == 1 {
			b.xDir = 0
			b.yDir = -1
		} else if b.yDir == -1 {
			b.yDir = 0
			b.xDir = 1
		} else if b.yDir == 1 {
			b.yDir = 0
			b.xDir = -1
		}
	} else if mirror == '\\' {
		if b.xDir == -1 {
			b.xDir = 0
			b.yDir = -1
		} else if b.xDir == 1 {
			b.xDir = 0
			b.yDir = 1
		} else if b.yDir == -1 {
			b.yDir = 0
			b.xDir = -1
		} else if b.yDir == 1 {
			b.yDir = 0
			b.xDir = 1
		}
	}

	b.move()
}

func simulateBeams(input []string, start *beam) int {
	beams := []*beam{start}

	// U, R, D, L
	tileCache := map[string][4]byte{}

	for len(beams) > 0 {
		newBeams := []*beam{}
		removeBeams := []int{}

		for i, beam := range beams {

			if !beam.inBounds(len(input[0]), len(input)) {
				removeBeams = append(removeBeams, i)
				continue
			}

			thisDir := [4]byte{}
			switch {
			case beam.xDir == 1:
				thisDir[1] = byte(1)
			case beam.xDir == -1:
				thisDir[3] = byte(1)
			case beam.yDir == 1:
				thisDir[2] = byte(1)
			case beam.yDir == -1:
				thisDir[0] = byte(1)
			}

			cacheKey := fmt.Sprintf("%d,%d", beam.x, beam.y)
			dirs, ok := tileCache[cacheKey]
			if ok {
				newDirs := [4]byte{}

				for i, b := range dirs {
					newDirs[i] = b | thisDir[i]
				}

				if newDirs == dirs {
					removeBeams = append(removeBeams, i)
					continue
				}

				tileCache[cacheKey] = newDirs

			} else {
				tileCache[cacheKey] = thisDir
			}

			tile := input[beam.y][beam.x]

			switch tile {
			case '.':
				beam.move()
			case '|', '-':
				newBeam := beam.split(tile)
				if newBeam != nil {
					newBeams = append(newBeams, newBeam)
				}
			case '/', '\\':
				beam.mirror(tile)
			}
		}

		for _, beamIdx := range removeBeams {
			beams[beamIdx] = nil
		}
		r := 0
		for _, beam := range beams {
			if beam != nil {
				beams[r] = beam
				r++
			}
		}
		beams = beams[:r]

		for _, newBeam := range newBeams {
			beams = append(beams, newBeam)
		}
	}

	return len(tileCache)
}

func daySixteenPartOne(input []string) int {
	startBeam := &beam{0, 0, 1, 0}
	return simulateBeams(input, startBeam)
}

func daySixteenPartTwo(input []string) int {
	startBeams := []*beam{}

	for i := 0; i < len(input[0]); i++ {
		topBeam := &beam{
			x: i,
			y: 0,
			xDir: 0,
			yDir: 1,
		}
		bottomBeam := &beam{
			x: i,
			y: len(input)-1,
			xDir: 0,
			yDir: -1,
		}
		startBeams = append(startBeams, topBeam, bottomBeam)
	}

	for i := 0; i < len(input); i++ {
		leftBeam := &beam{
			x: 0,
			y: i,
			xDir: 1,
			yDir: 0,
		}
		rightBeam := &beam{
			x: len(input[0])-1,
			y: i,
			xDir: -1,
			yDir: 0,
		}
		startBeams = append(startBeams, leftBeam, rightBeam)
	}

	biggest := 0
	for _, beam := range startBeams {
		n := simulateBeams(input, beam)
		if n > biggest {
			biggest = n
		}
	}

	return biggest
}

func daySixteenTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(16)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				`.|...\....`,
				`|.-.\.....`,
				`.....|-...`,
				`........|.`,
				`..........`,
				`.........\`,
				`..../.\\..`,
				`.-.-/..|..`,
				`.|....-|.\`,
				`..//.|....`,
			},
			expected: 46,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				`.|...\....`,
				`|.-.\.....`,
				`.....|-...`,
				`........|.`,
				`..........`,
				`.........\`,
				`..../.\\..`,
				`.-.-/..|..`,
				`.|....-|.\`,
				`..//.|....`,
			},
			expected: 51,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
