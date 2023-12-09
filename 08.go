package aoc

import (
	"fmt"
)

type node struct {
	id    string
	left  *node
	right *node
}

func buildNetwork(network []string) map[string]*node {
	allNodes := map[string]*node{}

	for _, line := range network {
		parent := line[:3]
		left := line[7:10]
		right := line[12:15]

		var leftNode, rightNode, parentNode *node

		gotLeft, ok := allNodes[left]
		if ok {
			leftNode = gotLeft
		} else {
			leftNode = &node{
				id: left,
			}
			allNodes[left] = leftNode
		}

		gotRight, ok := allNodes[right]
		if ok {
			rightNode = gotRight
		} else {
			rightNode = &node{
				id: right,
			}
			allNodes[right] = rightNode
		}

		gotParent, ok := allNodes[parent]
		if ok {
			parentNode = gotParent
			parentNode.left = leftNode
			parentNode.right = rightNode
		} else {
			parentNode = &node{
				id:    parent,
				left:  leftNode,
				right: rightNode,
			}
			allNodes[parent] = parentNode
		}
	}

	return allNodes
}

func dayEightPartOne(input []string) int {
	instructions := input[0]
	network := input[2:]

	allNodes := buildNetwork(network)

	curr := allNodes["AAA"]
	steps := 0

	looper := 0
	for {
		steps++

		switch instructions[looper] {
		case 'L':
			curr = curr.left
		case 'R':
			curr = curr.right
		}

		if curr.id == "ZZZ" {
			break
		}

		looper++
		if looper > len(instructions)-1 {
			looper = 0
		}
	}

	return steps
}

func dayEightPartTwo(input []string) int {
	instructions := input[0]
	network := input[2:]

	allNodes := buildNetwork(network)

	ghosts := []*node{}
	ghostSteps := []int{}
	for k, v := range allNodes {
		if k[2] == 'A' {
			ghosts = append(ghosts, v)
			ghostSteps = append(ghostSteps, 0)
		}
	}

	looper := 0
	for {
		activeGhosts := false

		for i := 0; i < len(ghosts); i++ {
			if ghosts[i] == nil {
				continue
			}
			activeGhosts = true

			ghostSteps[i]++

			switch instructions[looper] {
			case 'L':
				ghosts[i] = ghosts[i].left
			case 'R':
				ghosts[i] = ghosts[i].right
			}

			if ghosts[i].id[2] != 'Z' {
				continue
			}

			ghosts[i] = nil
		}

		if !activeGhosts {
			break
		}

		looper++
		if looper > len(instructions)-1 {
			looper = 0
		}
	}

	return LCM(ghostSteps[0], ghostSteps[1], ghostSteps[2:]...)
}

func dayEightTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(8)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			expected: 6,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"LR",
				"",
				"11A = (11B, XXX)",
				"11B = (XXX, 11Z)",
				"11Z = (11B, XXX)",
				"22A = (22B, XXX)",
				"22B = (22C, 22C)",
				"22C = (22Z, 22Z)",
				"22Z = (22B, 22B)",
				"XXX = (XXX, XXX)",
			},
			expected: 6,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}

// Credit due:
// https://go.dev/play/p/SmzvkDjYlb

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
