package main

import (
	"fmt"
	"os"
)

func parse_input(input string) [][]rune {
	// Parse the input string into rotation instructions
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return parse_paper_grid(string(data))
}

func parse_paper_grid(input string) [][]rune {
	lines := []rune{}
	grid := [][]rune{}

	for _, char := range input {
		if char == '\n' {
			if len(lines) > 0 {
				grid = append(grid, lines)
				lines = []rune{}
			}
		} else {
			lines = append(lines, char)
		}
	}
	if len(lines) > 0 {
		grid = append(grid, lines)
	}

	return grid
}

func print_grid(grid [][]rune) {
	for _, line := range grid {
		for _, char := range line {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}
}

func compare_grids(grid1, grid2 [][]rune) bool {
	if len(grid1) != len(grid2) {
		return false
	}
	for i := range grid1 {
		if len(grid1[i]) != len(grid2[i]) {
			return false
		}
		for j := range grid1[i] {
			if grid1[i][j] != grid2[i][j] {
				return false
			}
		}
	}
	return true
}

func part1(input [][]rune) (int, [][]rune) {
	// Implement the logic for part 1 of the puzzle here
	accesible_paper_rolls := 0
	resulting_grid := [][]rune{}
	max_adjecent := 4

	for i := range input {
		row := []rune{}
		for j := range input[i] {
			if input[i][j] == '@' {
				adjecent := 0
				// Check up
				if i > 0 && input[i-1][j] == '@' {
					adjecent++
				}
				// Check down
				if i < len(input)-1 && input[i+1][j] == '@' {
					adjecent++
				}
				// Check left
				if j > 0 && input[i][j-1] == '@' {
					adjecent++
				}
				// Check right
				if j < len(input[i])-1 && input[i][j+1] == '@' {
					adjecent++
				}
				// Check up - right
				if i > 0 && j < len(input[i])-1 && input[i-1][j+1] == '@' {
					adjecent++
				}
				// Check down - right
				if i < len(input)-1 && j < len(input[i])-1 && input[i+1][j+1] == '@' {
					adjecent++
				}
				// Check up - left
				if i > 0 && j > 0 && input[i-1][j-1] == '@' {
					adjecent++
				}
				// Check down - left
				if i < len(input)-1 && j > 0 && input[i+1][j-1] == '@' {
					adjecent++
				}

				if adjecent >= max_adjecent {
					row = append(row, input[i][j])
				} else {
					row = append(row, 'x')
					accesible_paper_rolls++
				}
			} else {
				row = append(row, input[i][j])
			}
		}
		resulting_grid = append(resulting_grid, row)
	}

	return accesible_paper_rolls, resulting_grid
}

func part2(input [][]rune) int {
	// Implement the logic for part 2 of the puzzle here
	accesible_paper_rolls := 0
	max_adjecent := 4

	removed_paper_rolls := 1
	safety_counter := 0
	for removed_paper_rolls > 0 {
		removed_paper_rolls = 0
		for i := range input {
			for j := range input[i] {
				if input[i][j] == '@' {
					adjecent := 0
					// Check up
					if i > 0 && input[i-1][j] == '@' {
						adjecent++
					}
					// Check down
					if i < len(input)-1 && input[i+1][j] == '@' {
						adjecent++
					}
					// Check left
					if j > 0 && input[i][j-1] == '@' {
						adjecent++
					}
					// Check right
					if j < len(input[i])-1 && input[i][j+1] == '@' {
						adjecent++
					}
					// Check up - right
					if i > 0 && j < len(input[i])-1 && input[i-1][j+1] == '@' {
						adjecent++
					}
					// Check down - right
					if i < len(input)-1 && j < len(input[i])-1 && input[i+1][j+1] == '@' {
						adjecent++
					}
					// Check up - left
					if i > 0 && j > 0 && input[i-1][j-1] == '@' {
						adjecent++
					}
					// Check down - left
					if i < len(input)-1 && j > 0 && input[i+1][j-1] == '@' {
						adjecent++
					}

					if adjecent >= max_adjecent {
						// Do nothing, stay as '@'
					} else {
						input[i][j] = 'x'
						accesible_paper_rolls++
						removed_paper_rolls++
					}
				}
			}
		}
		safety_counter++
		if safety_counter > 50 {
			fmt.Println("Safety break in part 2")
			break
		}
	}
	return accesible_paper_rolls
}

func test() {
	// Test with example data from the puzzle
	example_input_str := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	example_output_str := `..xx.xx@x.
x@@.@.@.@@
@@@@@.x.@@
@.@@@@..@.
x@.@@@@.@x
.@@@@@@@.@
.@.@.@.@@@
x.@@@.@@@@
.@@@@@@@@.
x.x.@@@.x.`

	example_rune := parse_paper_grid(example_input_str)

	excpeted_part1 := 13
	result_part1, grid1 := part1(example_rune)
	if result_part1 != excpeted_part1 {
		fmt.Printf("Part 1 failed: got %d, expected %d, difference %d\n", result_part1, excpeted_part1, excpeted_part1-result_part1)
	} else {
		fmt.Println("Part 1 passed")
	}
	if !compare_grids(grid1, parse_paper_grid(example_output_str)) {
		fmt.Println("Part 1 grid output failed:")
		fmt.Println("Expected:")
		print_grid(parse_paper_grid(example_output_str))
		fmt.Println("Got:")
		print_grid(grid1)
	}

	excpeted_part2 := 43
	result_part2 := part2(example_rune)
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
	accessible_paper_rolls, _ := part1(data)
	fmt.Println("Part 1 - Accessible Paper Rolls (max. 4 adjecent):", accessible_paper_rolls)

	fmt.Println("Part 2 - Accessible Paper Rolls (max. 4 adjecent, repeating):", part2(data))
}
