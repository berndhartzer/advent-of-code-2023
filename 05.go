package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func dayFivePartOne(input []string) int {
	seedMaps := [][8]int{}

	seedsStr := strings.Split(input[0], " ")
	seedsStr = seedsStr[1:]

	for _, seed := range seedsStr {
		newSeed := [8]int{-1, -1, -1, -1, -1, -1, -1, -1}

		n, err := strconv.Atoi(seed)
		if err != nil {
			panic("could not convert str to number")
		}
		newSeed[0] = n

		seedMaps = append(seedMaps, newSeed)
	}

	inputLooper := 3

	// There are 7 rounds of mappings
	for mappingLooper := 0; mappingLooper < 7; mappingLooper++ {
		for ; inputLooper < len(input); inputLooper++ {
			if input[inputLooper] == "" {
				break
			}

			split := strings.Split(input[inputLooper], " ")
			mapping := [3]int{}

			for i, s := range split {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic("could not convert str to number")
				}
				mapping[i] = n
			}

			dest := mapping[0]
			source := mapping[1]
			length := mapping[2]

			sourceUpper := source + (length - 1)
			mappingDiff := dest - source

			for i := 0; i < len(seedMaps); i++ {
				if seedMaps[i][mappingLooper+1] != -1 {
					continue
				}

				if seedMaps[i][mappingLooper] >= source && seedMaps[i][mappingLooper] <= sourceUpper {
					seedMaps[i][mappingLooper+1] = seedMaps[i][mappingLooper] + mappingDiff
				}
			}
		}

		// skip newline and heading
		inputLooper += 2

		// Map seeds that haven't matched an input mapping
		for i := 0; i < len(seedMaps); i++ {
			if seedMaps[i][mappingLooper+1] != -1 {
				continue
			}

			seedMaps[i][mappingLooper+1] = seedMaps[i][mappingLooper]
		}
	}

	lowest := -1
	for _, seedMap := range seedMaps {
		if lowest == -1 || seedMap[7] < lowest {
			lowest = seedMap[7]
		}
	}

	return lowest
}

// Struggled with this one. Special shoutout to this solution:
// https://github.com/conorgolden1/Jsaoc/blob/solutions/day_5/day_5.js
// which looked most similar to what I was trying to achieve, so I leaned on it
// heavily to get across the line
func dayFivePartTwo(input []string) int {
	seedsStr := strings.Split(input[0], " ")
	seedsStr = seedsStr[1:]

	seeds := []int{}
	for _, seed := range seedsStr {
		n, err := strconv.Atoi(seed)
		if err != nil {
			panic("could not convert str to number")
		}
		seeds = append(seeds, n)
	}

	mappings := [][][3]int{}
	mappingSet := [][3]int{}

	skipNext := false
	for i := 3; i < len(input); i++ {
		if input[i] == "" {
			skipNext = true
			continue
		}
		if skipNext {
			skipNext = false
			mappings = append(mappings, mappingSet)
			mappingSet = [][3]int{}
			continue
		}

		// destination, source, range
		mapping := [3]int{}

		split := strings.Split(input[i], " ")
		for j, s := range split {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic("could not convert str to number")
			}
			mapping[j] = n
		}

		mappingSet = append(mappingSet, mapping)
	}
	mappings = append(mappings, mappingSet)

	lowest := -1

	for seedIdx := 0; seedIdx < len(seeds); seedIdx += 2 {
		allSeeds := []int{seeds[seedIdx]}
		allRanges := []int{seeds[seedIdx+1]}

		for _, mappingSet := range mappings {
			newSeeds := []int{}
			newRanges := []int{}

			partialSeeds := allSeeds
			partialRanges := allRanges

			for len(partialSeeds) != 0 {
				seedFrom := partialSeeds[len(partialSeeds)-1]
				partialSeeds = partialSeeds[:len(partialSeeds)-1]

				seedRange := partialRanges[len(partialRanges)-1]
				partialRanges = partialRanges[:len(partialRanges)-1]

				seedTo := seedFrom + seedRange
				seedsLen := len(newSeeds)

				for _, mapping := range mappingSet {
					dest := mapping[0]
					mapFrom := mapping[1]
					mapRange := mapping[2]
					mapTo := mapFrom + mapRange

					if seedFrom < mapFrom && seedTo > mapTo {
						// range = []
						// seed = ()
						// ---(---[---]---)---

						partialSeeds = append(partialSeeds, seedFrom)
						partialRanges = append(partialRanges, (mapFrom - seedFrom))

						newSeeds = append(newSeeds, dest)
						newRanges = append(newRanges, (mapTo - mapFrom))

						partialSeeds = append(partialSeeds, mapTo)
						partialRanges = append(partialRanges, (seedTo - mapTo))
					} else if seedFrom < mapFrom && seedTo > mapFrom {
						// ---(---[---)---]---

						partialSeeds = append(partialSeeds, seedFrom)
						partialRanges = append(partialRanges, (mapFrom - seedFrom))

						newSeeds = append(newSeeds, dest)
						newRanges = append(newRanges, (seedTo - mapFrom))
					} else if seedFrom >= mapFrom && seedFrom < mapTo && seedTo > mapTo {
						// ---[---(---]---)---

						newSeeds = append(newSeeds, (seedFrom - mapFrom + dest))
						newRanges = append(newRanges, (mapTo - seedFrom - 1))

						partialSeeds = append(partialSeeds, mapTo)
						partialRanges = append(partialRanges, (seedTo - mapTo))
					} else if seedFrom >= mapFrom && seedTo <= mapTo {
						// ---[---(---)---]---

						newSeeds = append(newSeeds, (seedFrom - mapFrom + dest))
						newRanges = append(newRanges, seedRange)
					}
				}

				if len(newSeeds) == seedsLen {
					newSeeds = append(newSeeds, seedFrom)
					newRanges = append(newRanges, seedRange)
				}
			}

			allSeeds = newSeeds
			allRanges = newRanges
		}

		for _, seed := range allSeeds {
			if lowest == -1 || seed < lowest {
				lowest = seed
			}
		}
	}

	return lowest
}

func dayFiveTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(5)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			expected: 35,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			expected: 46,
		},
		"2": {
			input: []string{
				"seeds: 79 14",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			expected: 46,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
