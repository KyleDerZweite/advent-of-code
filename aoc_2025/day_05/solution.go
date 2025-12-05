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

// func parse_string(input string) (*Set, *Set) {
// 	// Is shit because takes ages!
// 	fresh_ids_set := NewSet()
// 	available_ingredient_ids_set := NewSet()
//
// 	fresh_id_lines := strings.Split(strings.TrimSpace(strings.Split(input, "\n\n")[0]), "\n")
// 	available_id_lines := strings.Split(strings.TrimSpace(strings.Split(input, "\n\n")[1]), "\n")
// 	for _, line := range fresh_id_lines {
// 		fresh_ingredient_id_start, _ := strconv.Atoi(strings.Split(line, "-")[0])
// 		fresh_ingredient_id_end, _ := strconv.Atoi(strings.Split(line, "-")[1])
// 		for i := fresh_ingredient_id_start; i <= fresh_ingredient_id_end; i++ {
// 			fresh_ids_set.Add(strconv.Itoa(i))
// 		}
// 	}
//
// 	for _, line := range available_id_lines {
// 		available_ingredient_ids_set.Add(strings.TrimSpace(line))
// 	}
// 	return fresh_ids_set, available_ingredient_ids_set
// }

func parse_string(input string) (map[int]int, []int) {
	// Is shit because takes ages!
	fresh_id_ranges := map[int]int{}
	ingredient_ids := []int{}

	fresh_id_lines := strings.Split(strings.TrimSpace(strings.Split(input, "\n\n")[0]), "\n")
	ingredient_id_lines := strings.Split(strings.TrimSpace(strings.Split(input, "\n\n")[1]), "\n")

	for _, line := range fresh_id_lines {
		fresh_ingredient_id_start, _ := strconv.Atoi(strings.Split(line, "-")[0])
		fresh_ingredient_id_end, _ := strconv.Atoi(strings.Split(line, "-")[1])
		fresh_id_ranges[fresh_ingredient_id_start] = fresh_ingredient_id_end
	}

	for _, line := range ingredient_id_lines {
		ingredient_id, _ := strconv.Atoi(strings.TrimSpace(line))
		ingredient_ids = append(ingredient_ids, ingredient_id)
	}
	return fresh_id_ranges, ingredient_ids
}

func check_id_between_range(check_id int, range_start int, range_end int) bool {
	return check_id >= range_start && check_id <= range_end
}

func part1(fresh_id_ranges map[int]int, ingredient_ids []int) int {
	fresh_id_count := 0
	for _, ingredient_id := range ingredient_ids {
		for range_start, range_end := range fresh_id_ranges {
			if check_id_between_range(ingredient_id, range_start, range_end) {
				fresh_id_count++
				break
			}
		}
	}
	return fresh_id_count
}

func part2(fresh_id_ranges map[int]int, ingredient_ids []int) int {
	return 0
}

func test() {
	// Test with example data from the puzzle
	example_input_str := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	fresh_id_ranges, ingredient_ids := parse_string(example_input_str)

	excpeted_part1 := 3
	excpeted_part2 := 0

	result_part1 := part1(fresh_id_ranges, ingredient_ids)
	if result_part1 != excpeted_part1 {
		fmt.Printf("Part 1 failed: got %d, expected %d, difference %d\n", result_part1, excpeted_part1, excpeted_part1-result_part1)
	} else {
		fmt.Println("Part 1 passed")
	}

	result_part2 := part2(fresh_id_ranges, ingredient_ids)
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
	fresh_id_ranges, ingredient_ids := parse_string(data)
	fmt.Println("Part 1 - Fresh ID Count:", part1(fresh_id_ranges, ingredient_ids))
	fmt.Println("Part 2 - Max Joltage Sum (12):", part2(fresh_id_ranges, ingredient_ids))
}
