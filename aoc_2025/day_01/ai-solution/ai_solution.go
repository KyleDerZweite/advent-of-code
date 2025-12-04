package main

import (
	"fmt"
	"os"
)

// Instruction represents a parsed move instruction
// Using int32 for better memory alignment and cache efficiency (8 bytes per instruction)
type Instruction struct {
	stepsRaw int32 // original steps for part 2 calculation
	stepsMod int32 // steps % 100 (can be negative for left)
}

// parseInput parses the input file directly into pre-allocated instruction slice
// Avoids string allocations and uses direct byte parsing
func parseInput(input string) []Instruction {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// Count newlines to pre-allocate slice (avoids reallocations)
	n := len(data)
	count := 0
	for i := 0; i < n; i++ {
		if data[i] == '\n' {
			count++
		}
	}
	// Add 1 for last line if no trailing newline
	if n > 0 && data[n-1] != '\n' {
		count++
	}

	instructions := make([]Instruction, 0, count)
	i := 0

	for i < n {
		// Skip any whitespace/newlines
		for i < n && (data[i] == '\n' || data[i] == '\r' || data[i] == ' ') {
			i++
		}
		if i >= n {
			break
		}

		// Parse direction (L or R)
		left := data[i] == 'L'
		i++

		// Parse number directly (inline atoi - much faster than strconv.Atoi)
		// Unrolled for common cases (1-3 digit numbers)
		var steps int32
		if i < n && data[i] >= '0' && data[i] <= '9' {
			steps = int32(data[i] - '0')
			i++
			if i < n && data[i] >= '0' && data[i] <= '9' {
				steps = steps*10 + int32(data[i]-'0')
				i++
				if i < n && data[i] >= '0' && data[i] <= '9' {
					steps = steps*10 + int32(data[i]-'0')
					i++
					// Handle 4+ digit numbers (rare)
					for i < n && data[i] >= '0' && data[i] <= '9' {
						steps = steps*10 + int32(data[i]-'0')
						i++
					}
				}
			}
		}

		// Pre-compute mod and sign for the hot loop
		stepsMod := steps % 100
		if left {
			stepsMod = -stepsMod
		}
		instructions = append(instructions, Instruction{stepsRaw: steps, stepsMod: stepsMod})
	}

	return instructions
}

// solve computes both parts in a single pass for maximum efficiency
// Using int32 arithmetic which can be faster on some CPUs
func solve(start int, instructions []Instruction) (int, int) {
	password1 := 0
	password2 := 0
	dial := int32(start)

	// Get slice header for direct access
	n := len(instructions)
	for idx := 0; idx < n; idx++ {
		stepsRaw := instructions[idx].stepsRaw
		stepsMod := instructions[idx].stepsMod

		// Part 2: count zero crossings
		if stepsMod < 0 {
			// Going left
			distTo0 := dial
			if distTo0 == 0 {
				distTo0 = 100
			}
			if stepsRaw >= distTo0 {
				password2 += 1 + int((stepsRaw-distTo0)/100)
			}
		} else {
			// Going right
			distTo0 := 100 - dial
			if dial == 0 {
				distTo0 = 100
			}
			if stepsRaw >= distTo0 {
				password2 += 1 + int((stepsRaw-distTo0)/100)
			}
		}

		// Update dial position - stepsMod already has the sign baked in
		dial += stepsMod
		// Normalize to [0, 99] range
		if dial < 0 {
			dial += 100
		} else if dial >= 100 {
			dial -= 100
		}

		// Part 1: count when dial lands on 0
		// Branchless increment: password1 += (dial == 0) ? 1 : 0
		// But Go doesn't have ternary, so we use a trick
		if dial == 0 {
			password1++
		}
	}

	return password1, password2
}

// part1 wrapper for testing compatibility
func part1(start int, rotation []string) int {
	instructions := make([]Instruction, len(rotation))
	for i, move := range rotation {
		left := move[0] == 'L'
		// Fast inline atoi
		var steps int32
		for j := 1; j < len(move); j++ {
			steps = steps*10 + int32(move[j]-'0')
		}
		stepsMod := steps % 100
		if left {
			stepsMod = -stepsMod
		}
		instructions[i] = Instruction{stepsRaw: steps, stepsMod: stepsMod}
	}
	p1, _ := solve(start, instructions)
	return p1
}

// part2 wrapper for testing compatibility
func part2(start int, rotation []string) int {
	instructions := make([]Instruction, len(rotation))
	for i, move := range rotation {
		left := move[0] == 'L'
		// Fast inline atoi
		var steps int32
		for j := 1; j < len(move); j++ {
			steps = steps*10 + int32(move[j]-'0')
		}
		stepsMod := steps % 100
		if left {
			stepsMod = -stepsMod
		}
		instructions[i] = Instruction{stepsRaw: steps, stepsMod: stepsMod}
	}
	_, p2 := solve(start, instructions)
	return p2
}

func test() {
	// Test with example data from the puzzle
	exampleStart := 50
	exampleRotation := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}

	expectedPart1 := 3
	expectedPart2 := 6

	resultPart1 := part1(exampleStart, exampleRotation)
	resultPart2 := part2(exampleStart, exampleRotation)

	if resultPart1 != expectedPart1 {
		fmt.Printf("Part 1 failed: got %d, expected %d\n", resultPart1, expectedPart1)
	} else {
		fmt.Println("Part 1 passed")
	}

	if resultPart2 != expectedPart2 {
		fmt.Printf("Part 2 failed: got %d, expected %d\n", resultPart2, expectedPart2)
	} else {
		fmt.Println("Part 2 passed")
	}
}

func main() {
	// Run tests
	test()

	// The dial starts by pointing at 50.
	start := 50
	instructions := parseInput("input.txt")

	// Solve both parts in a single pass
	p1, p2 := solve(start, instructions)
	fmt.Println("Part 1 - Password:", p1)
	fmt.Println("Part 2 - Password:", p2)
}
