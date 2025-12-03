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

func parse_id_ranges(example_product_id_ranges string) []int {
	ids := []int{}
	ranges := strings.Split(example_product_id_ranges, ",")
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		for i := start; i <= end; i++ {
			ids = append(ids, i)
		}
	}
	return ids
}

func part1(product_id_ranges string) int {
	// Implement the logic for part 1 of the puzzle here
	invalid_ids_sum := 0
	ids := parse_id_ranges(product_id_ranges)
	for _, id := range ids {

		id_str := strconv.Itoa(id)
		digits := strings.Split(id_str, "")

		len_digits := len(digits)
		if len_digits%2 != 0 {
			// skip odd-length IDs
			continue
		}
		sequence_0 := id_str[0 : len_digits/2]
		sequence_1 := id_str[len_digits/2 : len_digits]

		// fmt.Printf("ID: %d Seq0: %s Seq1: %s\n", id, sequence_0, sequence_1)
		if sequence_0 == sequence_1 {
			invalid_ids_sum += id
		}
	}

	return invalid_ids_sum
}

func part2(product_id_ranges string) int {
	invalid_ids_sum := 0
	ids := parse_id_ranges(product_id_ranges)
	for _, id := range ids {
		id_str := strconv.Itoa(id)
		len_digits := len(id_str)

		isInvalid := false
		// Check for each possible pattern length from 1 to len_digits/2
		for patternLen := 1; patternLen <= len_digits/2; patternLen++ {
			if len_digits%patternLen == 0 {
				// Check if id_str is made of the pattern repeated
				pattern := id_str[:patternLen]
				allMatch := true
				for i := patternLen; i < len_digits; i += patternLen {
					if id_str[i:i+patternLen] != pattern {
						allMatch = false
						break
					}
				}
				if allMatch {
					isInvalid = true
					break
				}
			}
		}

		if isInvalid {
			invalid_ids_sum += id
		}
	}
	return invalid_ids_sum
}

func test() {
	// Test with example data from the puzzle
	example_product_id_ranges := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	excpeted_part1 := 1227775554
	excpeted_part2 := 4174379265

	result_part1 := part1(example_product_id_ranges)
	if result_part1 != excpeted_part1 {
		fmt.Printf("Part 1 failed: got %d, expected %d\n", result_part1, excpeted_part1)
	} else {
		fmt.Println("Part 1 passed")
	}

	result_part2 := part2(example_product_id_ranges)
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

	fmt.Println("Part 1 - Invalid IDs Sum (twice-pattern):", part1(data))
	fmt.Println("Part 2 - Invalid IDs Sum (repeated pattern):", part2(data))
}
