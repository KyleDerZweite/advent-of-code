package main

import (
	"fmt"
	"os"
)

func parse_input(input string) string {
	// Parse the input string into rotation instructions
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func part1(input string) int {
	// Implement the logic for part 1 of the puzzle here
	joltage_sum := 0
	return joltage_sum
}

func part2(input string) int {
	joltage_sum := 0
	return joltage_sum
}

func test() {
	// Test with example data from the puzzle
	example_input_str := ""

	excpeted_part1 := 0
	excpeted_part2 := 0

	result_part1 := part1(example_input_str)
	if result_part1 != excpeted_part1 {
		fmt.Printf("Part 1 failed: got %d, expected %d, difference %d\n", result_part1, excpeted_part1, excpeted_part1-result_part1)
	} else {
		fmt.Println("Part 1 passed")
	}

	result_part2 := part2(example_input_str)
	if result_part2 != excpeted_part2 {
		fmt.Printf("Part 2 failed: got %d, expected %d, difference %d\n", result_part2, excpeted_part2, excpeted_part2-result_part2)
	} else {
		fmt.Println("Part 2 passed")
	}
}

func main() {
	// Run tests
	test()

	// data := parse_input("input.txt")
	// fmt.Println("Part 1 - Max Joltage Sum (2):", part1(data))
	// fmt.Println("Part 2 - Max Joltage Sum (12):", part2(data))
}
