package main

import (
	"fmt"
	"os"
	"strings"
)

func parse_input(input string) string {
	// Parse the input string into rotation instructions
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func is_letter_digit(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') // a-Z and 0-9 matching
}

func part1(input string) int {
	// Implement the logic for part 1 of the puzzle here
	unique_antinote_locations := 0

	locations_set := make(map[string]bool)
	lines := strings.Split(input, "\n")

	is_char_set := false

	// Find the current character and its position
	for y, line := range lines {
		for x, char := range line {
			if is_letter_digit(char) && !is_char_set {
				is_char_set = true
			} else {
				continue
			}
			for i := 0; i < len(lines); i++ {
				for j := 0; j < len(lines[i]); j++ {
					other_char := lines[i][j]
					if rune(other_char) != char || (y == i && x == j) {
						continue
					}
					x_distance := j - x
					y_distance := i - y
					loc_key := fmt.Sprintf("%d,%d", x_distance, y_distance)
					locations_set[loc_key] = true
					// fmt.Printf("Distances: '%d':'%d' Chars: %s:%s\n", x_distance, y_distance, string(char), string(other_char))
				}
			}
			is_char_set = false
		}
	}

	unique_antinote_locations = len(locations_set)
	return unique_antinote_locations
}

func part2(input string) int {
	unique_antinote_locations := 0
	return unique_antinote_locations
}

func test() {
	// Test with example data from the puzzle
	example_input_str := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	excpeted_part1 := 14
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

	data := parse_input("input.txt")
	fmt.Println("Part 1 - Unique antinode locations:", part1(data))
	// fmt.Println("Part 2 - Max Joltage Sum (12):", part2(data))
}
