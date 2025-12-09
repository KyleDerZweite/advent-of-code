package main

import (
	"fmt"
	"math"
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

func part1(input string) int {
	// Implement the logic for part 1 of the puzzle here
	lines := strings.Split(strings.TrimSpace(input), "\n")

	max_tiles := 0
	for _, line := range lines {
		line_x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		line_y, _ := strconv.Atoi(strings.Split(line, ",")[1])

		for _, other_line := range lines {
			other_line_x, _ := strconv.Atoi(strings.Split(other_line, ",")[0])
			other_line_y, _ := strconv.Atoi(strings.Split(other_line, ",")[1])

			if line == other_line {
				continue
			}
			x_distance := math.Abs(float64(other_line_x-line_x)) + 1
			y_distance := math.Abs(float64(other_line_y-line_y)) + 1

			new_max_tiles := x_distance * y_distance

			if int(new_max_tiles) > max_tiles {
				max_tiles = int(new_max_tiles)
			}
		}

	}
	return int(max_tiles)
}

type tile struct {
	x, y int
}

func part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	valid_tiles := make(map[tile]bool) // sparse set: red + green + interior green

	red_tiles := []tile{}
	min_x, min_y, max_x, max_y := 10000, 10000, 0, 0

	// Parse red tiles and find bounds
	for _, line := range lines {
		line_x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		line_y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		if line_x > max_x {
			max_x = line_x
		}
		if line_y > max_y {
			max_y = line_y
		}
		if line_x < min_x {
			min_x = line_x
		}
		if line_y < min_y {
			min_y = line_y
		}
		red_tile := tile{x: line_x, y: line_y}
		red_tiles = append(red_tiles, red_tile)
		valid_tiles[red_tile] = true
	}

	// Add green tiles connecting consecutive red tiles
	for i := 0; i < len(red_tiles); i++ {
		rt1 := red_tiles[i]
		rt2 := red_tiles[(i+1)%len(red_tiles)]
		if rt1.x == rt2.x {
			min_y := int(math.Min(float64(rt1.y), float64(rt2.y)))
			max_y := int(math.Max(float64(rt1.y), float64(rt2.y)))
			for y := min_y + 1; y < max_y; y++ {
				valid_tiles[tile{x: rt1.x, y: y}] = true
			}
		} else if rt1.y == rt2.y {
			min_x := int(math.Min(float64(rt1.x), float64(rt2.x)))
			max_x := int(math.Max(float64(rt1.x), float64(rt2.x)))
			for x := min_x + 1; x < max_x; x++ {
				valid_tiles[tile{x: x, y: rt1.y}] = true
			}
		}
	}

	print_tiles(valid_tiles, min_y-1, min_x-1, max_y+1, max_x+1)

	// Find largest valid rectangle between two red tiles
	max_area := 0

	return max_area
}

func print_tiles(tile_map map[tile]bool, min_y int, min_x int, max_y int, max_x int) {
	for y := min_y; y <= max_y; y++ {
		for x := min_x; x <= max_x; x++ {
			t := tile{x: x, y: y}
			val, _ := tile_map[t]
			if val {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func test() {
	// Test with example data from the puzzle
	example_input_str := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

	excpeted_part1 := 50
	excpeted_part2 := 24

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
	// fmt.Println("Part 1 - Max Tiles Between Points:", part1(data))
	// fmt.Println("Part 2 - Max Red and Green Tiles:", part2(data))
}
