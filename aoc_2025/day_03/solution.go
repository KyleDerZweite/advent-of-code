package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func parse_joltage_str(joltage_str string) []string {
	lines := strings.Split(strings.TrimSpace(joltage_str), "\n")
	joltages := []string{}
	for _, line := range lines {
		joltages = append(joltages, line)
	}
	return joltages
}

func parse_input(input string) []string {
	// Parse the input string into rotation instructions
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return parse_joltage_str(string(data))
}

func part1(joltages []string) int {
	// Implement the logic for part 1 of the puzzle here
	joltage_sum := 0

	for _, joltage := range joltages {
		len_joltage := len(joltage)

		joltage_max := 0
		joltage_max_index := -1

		for i := 0; i < len_joltage-1; i++ {
			joltage_digit := int(joltage[i] - '0')
			if joltage_digit > joltage_max {
				joltage_max = joltage_digit
				joltage_max_index = i
			}
		}
		joltage_sum += joltage_max * 10

		joltage_max = 0
		for i := joltage_max_index + 1; i < len_joltage; i++ {
			joltage_digit := int(joltage[i] - '0')
			if joltage_digit > joltage_max {
				joltage_max = joltage_digit
				joltage_max_index = i
			}
		}
		joltage_sum += joltage_max
	}
	return joltage_sum
}

func part2(joltages []string) int {
	joltage_sum := 0

	for _, joltage := range joltages {
		len_joltage := len(joltage)

		battery_digits := 12
		joltage_index := 0
		for i := 0; i < 12; i++ {
			joltage_max := 0
			joltage_factor := math.Pow(float64(10), float64(battery_digits-1-i))

			for j := joltage_index + 1; j < len_joltage-1; j++ {
				joltage_digit := int(joltage[j] - '0')
				if joltage_digit > joltage_max {
					joltage_max = joltage_digit
					joltage_index = j
				}
			}
			fmt.Printf("Joltage max: %d at index %d with factor %d\n", joltage_max, joltage_index, int(joltage_factor))
			joltage_sum += (joltage_max * int(joltage_factor))
		}
	}
	return joltage_sum
}

func test() {
	// Test with example data from the puzzle
	example_joltage_str := `987654321111111
811111111111119
234234234234278
818181911112111`
	example_joltage := parse_joltage_str(example_joltage_str)

	excpeted_part1 := 357
	excpeted_part2 := 3121910778619

	result_part1 := part1(example_joltage)
	if result_part1 != excpeted_part1 {
		fmt.Printf("Part 1 failed: got %d, expected %d, difference %d\n", result_part1, excpeted_part1, excpeted_part1-result_part1)
	} else {
		fmt.Println("Part 1 passed")
	}

	result_part2 := part2(example_joltage)
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
