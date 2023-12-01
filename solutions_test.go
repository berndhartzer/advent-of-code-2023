package aoc

import (
	"fmt"
	"testing"
	"time"
)

func TestSolutions(t *testing.T) {
	t.Run("day 1", func(t *testing.T) {
		runner := func(t *testing.T, tests map[string]dayOneTestConfig, fn func([]string) int) {
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

		partOne, partTwo, err := getDayOneTests()
		if err != nil {
			t.Errorf("failed to get tests: %v", err)
		}

		t.Run("part 1", func(t *testing.T) {
			runner(t, partOne, dayOnePartOne)
		})

		t.Run("part 2", func(t *testing.T) {
			runner(t, partTwo, dayOnePartTwo)
		})
	})

	t.Run("day 2", func(t *testing.T) {
		runner := func(t *testing.T, tests map[string]dayTwoTestConfig, fn func([]string) int) {
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

		partOne, _, err := getDayTwoTests()
		if err != nil {
			t.Errorf("failed to get tests: %v", err)
		}

		t.Run("part 1", func(t *testing.T) {
			runner(t, partOne, dayTwoPartOne)
		})
	})
}