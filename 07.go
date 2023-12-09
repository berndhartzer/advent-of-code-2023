package aoc

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

func daySevenPartOne(input []string) int {
	cardScores := map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}

	handScores := []int{}
	scoreToBid := map[int]int{}

	for _, line := range input {
		hand := line[:5]
		bidStr := line[6:]

		handScore := 0

		handCount := map[rune]int{}
		for i, c := range hand {
			cardValue := float64(cardScores[c]) * math.Pow(100, float64(4-i))
			handScore += int(cardValue)
			handCount[c]++
		}

		typeScore := func() int {
			switch len(handCount) {
			case 1:
				return 70_000_000_000
			case 4:
				return 20_000_000_000
			case 5:
				return 10_000_000_000
			case 2:
				for _, v := range handCount {
					if v == 4 || v == 1 {
						return 60_000_000_000
					}
					if v == 3 || v == 2 {
						return 50_000_000_000
					}
				}
			}

			// handCount is 3
			for _, v := range handCount {
				if v == 3 {
					return 40_000_000_000
				}
				if v == 2 {
					return 30_000_000_000
				}
			}

			return 0
		}()

		handScore += typeScore

		handScores = append(handScores, handScore)

		n, err := strconv.Atoi(bidStr)
		if err != nil {
			panic("failed to convert string to num")
		}
		scoreToBid[handScore] = n
	}

	slices.Sort(handScores)

	total := 0
	for i, score := range handScores {
		total += scoreToBid[score] * (i + 1)
	}

	return total
}

func daySevenPartTwo(input []string) int {
	cardScores := map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}

	handScores := []int{}
	scoreToBid := map[int]int{}

	for _, line := range input {
		hand := line[:5]
		bidStr := line[6:]

		handScore := 0

		// handScore should stay the same at this point for purpose of tie breakers
		handCount := map[rune]int{}
		for i, c := range hand {
			cardValue := float64(cardScores[c]) * math.Pow(100, float64(4-i))
			handScore += int(cardValue)
			handCount[c]++
		}

		most := 0
		var mostCard rune
		for k, v := range handCount {
			if k == 'J' {
				continue
			}

			if v > most {
				most = v
				mostCard = k
			}
		}

		jokers, ok := handCount['J']
		if ok {
			handCount[mostCard] += jokers
			most += jokers
			delete(handCount, 'J')
		}

		typeScore := func() int {
			switch most {
			case 5:
				return 70_000_000_000
			case 4:
				return 60_000_000_000
			case 3:
				for _, v := range handCount {
					if v == 2 {
						return 50_000_000_000
					}
				}
				return 40_000_000_000
			case 2:
				pairs := 0
				for _, v := range handCount {
					if v == 2 {
						pairs++
					}
				}

				if pairs == 2 {
					return 30_000_000_000
				}

				return 20_000_000_000
			}

			return 10_000_000_000
		}()

		handScore += typeScore

		handScores = append(handScores, handScore)

		n, err := strconv.Atoi(bidStr)
		if err != nil {
			panic("failed to convert string to num")
		}
		scoreToBid[handScore] = n
	}

	slices.Sort(handScores)

	total := 0
	for i, score := range handScores {
		total += scoreToBid[score] * (i + 1)
	}

	return total
}

type daySevenTestConfig struct {
	input     []string
	expected  int
	logResult bool
}

func getDaySevenTests() (map[string]daySevenTestConfig, map[string]daySevenTestConfig, error) {
	fileInput, err := getInput(7)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]daySevenTestConfig{
		"1": {
			input: []string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			expected: 6440,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]daySevenTestConfig{
		"1": {
			input: []string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			expected: 5905,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
