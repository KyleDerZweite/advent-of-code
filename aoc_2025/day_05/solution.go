package main

import (
	"fmt"
	"os"
	"strconv"
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

func parse_string(input string) (*Set, *Set) {
	fresh_ids_set := NewSet()
	available_ingredient_ids_set := NewSet()

	fresh_id_lines := strings.Split(strings.TrimSpace(strings.Split(input, "\n\n")[0]), "\n")
	available_id_lines := strings.Split(strings.TrimSpace(strings.Split(input, "\n\n")[1]), "\n")
	for _, line := range fresh_id_lines {
		fresh_ingredient_id_start, _ := strconv.Atoi(strings.Split(line, "-")[0])
		fresh_ingredient_id_end, _ := strconv.Atoi(strings.Split(line, "-")[1])
		for i := fresh_ingredient_id_start; i <= fresh_ingredient_id_end; i++ {
			fresh_ids_set.Add(strconv.Itoa(i))
		}
	}

	for _, line := range available_id_lines {
		available_ingredient_ids_set.Add(strings.TrimSpace(line))
	}
	return fresh_ids_set, available_ingredient_ids_set
}

func part1(expected_input_fresh_ids *Set, expected_input_available_ids *Set) int {
	return expected_input_fresh_ids.Intersection(expected_input_available_ids).Size()
}

func part2(expected_input_fresh_ids *Set, expected_input_available_ids *Set) int {
	return 0
}

func test() {
	// Test with example data from the puzzle
	example_input_str := `
	3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	expected_input_fresh_ids, expected_input_available_ids := parse_string(example_input_str)

	excpeted_part1 := 3
	excpeted_part2 := 0

	result_part1 := part1(expected_input_fresh_ids, expected_input_available_ids)
	if result_part1 != excpeted_part1 {
		fmt.Printf("Part 1 failed: got %d, expected %d, difference %d\n", result_part1, excpeted_part1, excpeted_part1-result_part1)
	} else {
		fmt.Println("Part 1 passed")
	}

	result_part2 := part2(expected_input_fresh_ids, expected_input_available_ids)
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
	expected_input_fresh_ids, expected_input_available_ids := parse_string(data)
	fmt.Println("Part 1 - Max Joltage Sum (2):", part1(expected_input_fresh_ids, expected_input_available_ids))
	fmt.Println("Part 2 - Max Joltage Sum (12):", part2(expected_input_fresh_ids, expected_input_available_ids))
}
