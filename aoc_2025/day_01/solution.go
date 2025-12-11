package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_input(input string) []string {
	// Parse the input string into rotation instructions
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	// Trim whitespace and split by newline
	content := strings.TrimSpace(string(data))
	rotation := strings.Split(content, "\n")
	return rotation
}

func part1(start int, rotation []string) int {
	// Implement the logic for part 1 of the puzzle here
	password := 0
	dial := start
	for _, move := range rotation {
		direction := move[0]
		steps, _ := strconv.Atoi(move[1:])

		if direction == 'L' {
			dial = ((dial-steps)%100 + 100) % 100
		} else if direction == 'R' {
			dial = (dial + steps) % 100
		}

		if dial == 0 {
			password += 1
		}
		// fmt.Printf("%c %d\n", direction, steps)
	}

	return password
}

func part2(start int, rotation []string) int {
	// Count how many times the dial points at 0 (during or at end of rotation)
	// Starting position doesn't count - only when we ARRIVE at 0
	password := 0
	dial := start
	for _, move := range rotation {
		direction := move[0]
		steps, _ := strconv.Atoi(move[1:])

		if direction == 'L' {
			// Going left: we hit 0 at step=dial, dial+100, dial+200, etc.
			// So we hit 0 if steps >= dial (when dial > 0), or steps >= 100 (when dial == 0)
			dist_to_0 := dial
			if dist_to_0 == 0 {
				dist_to_0 = 100 // if at 0, next 0 is 100 steps away going left
			}
			if steps >= dist_to_0 {
				password += 1 + (steps-dist_to_0)/100
			}
			dial = ((dial-steps)%100 + 100) % 100
		} else if direction == 'R' {
			// Going right: we hit 0 at step=(100-dial), (200-dial), etc.
			// So we hit 0 if steps >= (100-dial) (when dial > 0), or steps >= 100 (when dial == 0)
			dist_to_0 := (100 - dial) % 100
			if dist_to_0 == 0 {
				dist_to_0 = 100 // if at 0, next 0 is 100 steps away going right
			}
			if steps >= dist_to_0 {
				password += 1 + (steps-dist_to_0)/100
			}
			dial = (dial + steps) % 100
		}
		// fmt.Printf("%c %d -> %d (total: %d)\n", direction, steps, dial, password)
	}

	return password
}

func test() {
	// Test with example data from the puzzle
	example_start := 50
	example_rotation := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}

	excpeted_part1 := 3
	excpeted_part2 := 6

	result_part1 := part1(example_start, example_rotation)
	result_part2 := part2(example_start, example_rotation)

	if result_part1 != excpeted_part1 {
		fmt.Printf("Part 1 failed: got %d, expected %d\n", result_part1, excpeted_part1)
	} else {
		fmt.Println("Part 1 passed")
	}

	if result_part2 != excpeted_part2 {
		fmt.Printf("Part 2 failed: got %d, expected %d\n", result_part2, excpeted_part2)
	} else {
		fmt.Println("Part 2 passed")
	}
}

func main() {
	// Run tests
	test()

	// The dial starts by pointing at 50.
	start := 50
	rotation := parse_input("input.txt")

	fmt.Println("Part 1 - Password:", part1(start, rotation))
	fmt.Println("Part 2 - Password:", part2(start, rotation))
}
