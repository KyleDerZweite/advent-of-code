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

func part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return 0
	}

	start_x, start_y := -1, -1
	for y, line := range lines {
		if idx := strings.Index(line, "S"); idx != -1 {
			start_x, start_y = idx, y
			break
		}
	}
	if start_x == -1 {
		return 0
	}

	splits := 0
	beams := map[int]struct{}{start_x: {}}

	for y := start_y + 1; y < len(lines); y++ {
		line := lines[y]
		next := make(map[int]struct{})

		for x := range beams {
			if x < 0 || x >= len(line) {
				continue
			}
			if line[x] == '^' {
				splits++
				if x-1 >= 0 {
					next[x-1] = struct{}{}
				}
				if x+1 < len(line) {
					next[x+1] = struct{}{}
				}
			} else {
				next[x] = struct{}{}
			}
		}

		beams = next
		if len(beams) == 0 {
			break
		}
	}

	return splits
}

func part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return 0
	}

	start_x, start_y := -1, -1
	for y, line := range lines {
		if idx := strings.Index(line, "S"); idx != -1 {
			start_x, start_y = idx, y
			break
		}
	}
	if start_x == -1 {
		return 0
	}

	timelines := 0
	beams := map[int]int{start_x: 1}

	for y := start_y + 1; y < len(lines); y++ {
		line := lines[y]
		next := make(map[int]int)
		for x := range beams {
			if x < 0 || x >= len(line) {
				continue
			}
			if line[x] == '^' {
				if x-1 >= 0 {
					next[x-1] += beams[x]
				}
				if x+1 < len(line) {
					next[x+1] += beams[x]
				}
			} else {
				next[x] += beams[x]
			}
		}
		// fmt.Println(len(next))
		beams = next
		if len(beams) == 0 {
			break
		}
	}
	// fmt.Println(beams)

	for _, count := range beams {
		timelines += count
	}

	return timelines
}

func test() {
	// Test with example data from the puzzle
	example_input_str := `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

	excpeted_part1 := 21
	excpeted_part2 := 40

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
	fmt.Println("Part 1 - Total Splits:", part1(data))
	fmt.Println("Part 2 - Total different timelines:", part2(data))
}
