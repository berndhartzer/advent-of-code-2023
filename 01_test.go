package aoc

import (
	"fmt"
	"testing"
	"time"
)

func dayOnePartOne(input []string) int {
	total := 0

	for _, line := range input {
		l := 0
		r := len(line)-1

		gotFirst, gotLast := false, false
		var first, last byte

		for {
			// 49 == 1, 57 == 9
			if !gotFirst && line[l] >= 49 && line[l] <= 57 {
				first = line[l]
				gotFirst = true
			}
			if !gotLast && line[r] >= 49 && line[r] <= 57 {
				last = line[r]
				gotLast = true
			}
			if gotFirst && gotLast {
				break
			}

			l++
			r--
		}

		// convert byte to int value by subtracting 48
		// e.g. 49 - 48 == 1
		total += (int(first)-48)*10
		total += (int(last)-48)
	}

	return total
}

func TestDayOne(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	fileInput, err := getInput(1)
	if err != nil {
		t.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	runTests := func(t *testing.T, tests map[string]testConfig, fn func([]string) int) {
		for name, cfg := range tests {
			cfg := cfg
			t.Run(name, func(t *testing.T) {
				start := time.Now()
				output := fn(cfg.input)
				finish := time.Since(start)
				if cfg.logResult {
					t.Log(fmt.Sprintf("\nsolution:\t%v\nelapsed time:\t%s", output, finish))
					return
				}

				if output != cfg.expected {
					t.Fatalf("Incorrect output - got: %v, want: %v", output, cfg.expected)
				}
			})
		}
	}

	t.Run("part one", func(t *testing.T) {
		tests := map[string]testConfig{
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, dayOnePartOne)
	})
}
