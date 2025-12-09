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

type segment struct {
	x1, y1, x2, y2 int
}

func part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Parse red tiles
	red_tiles := []struct{ x, y int }{}
	for _, line := range lines {
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		red_tiles = append(red_tiles, struct{ x, y int }{x, y})
	}

	// Build segments connecting consecutive red tiles
	segments := []segment{}
	for i := 0; i < len(red_tiles); i++ {
		rt1 := red_tiles[i]
		rt2 := red_tiles[(i+1)%len(red_tiles)]
		segments = append(segments, segment{rt1.x, rt1.y, rt2.x, rt2.y})
	}

	// For each pair of red tiles, check if the rectangle is valid
	max_area := 0
	for i := 0; i < len(red_tiles); i++ {
		for j := i + 1; j < len(red_tiles); j++ {
			r1, r2 := red_tiles[i], red_tiles[j]
			if r1.x == r2.x || r1.y == r2.y {
				continue
			}

			min_x := int(math.Min(float64(r1.x), float64(r2.x)))
			max_x := int(math.Max(float64(r1.x), float64(r2.x)))
			min_y := int(math.Min(float64(r1.y), float64(r2.y)))
			max_y := int(math.Max(float64(r1.y), float64(r2.y)))

			// Check all 4 corners are inside or on boundary
			if !is_inside(min_x, min_y, segments) || !is_inside(max_x, min_y, segments) ||
				!is_inside(min_x, max_y, segments) || !is_inside(max_x, max_y, segments) {
				continue
			}

			// Check that no segment crosses through our rectangle (cutting it)
			if rectangle_is_cut(min_x, min_y, max_x, max_y, segments) {
				continue
			}

			area := (max_x - min_x + 1) * (max_y - min_y + 1)
			if area > max_area {
				max_area = area
			}
		}
	}

	return max_area
}

func is_inside(x, y int, segments []segment) bool {
	// Check if on boundary first
	for _, seg := range segments {
		if seg.x1 == seg.x2 { // vertical
			min_y := int(math.Min(float64(seg.y1), float64(seg.y2)))
			max_y := int(math.Max(float64(seg.y1), float64(seg.y2)))
			if x == seg.x1 && y >= min_y && y <= max_y {
				return true
			}
		} else { // horizontal
			min_x := int(math.Min(float64(seg.x1), float64(seg.x2)))
			max_x := int(math.Max(float64(seg.x1), float64(seg.x2)))
			if y == seg.y1 && x >= min_x && x <= max_x {
				return true
			}
		}
	}

	// Ray casting: count vertical segments to the right
	crossings := 0
	for _, seg := range segments {
		if seg.x1 == seg.x2 { // vertical segment
			min_y := int(math.Min(float64(seg.y1), float64(seg.y2)))
			max_y := int(math.Max(float64(seg.y1), float64(seg.y2)))
			if seg.x1 > x && y > min_y && y < max_y {
				crossings++
			}
		}
	}
	return crossings%2 == 1
}

func rectangle_is_cut(min_x, min_y, max_x, max_y int, segments []segment) bool {
	// Check if any segment passes through the interior of our rectangle
	for _, seg := range segments {
		if seg.x1 == seg.x2 { // vertical segment
			seg_min_y := int(math.Min(float64(seg.y1), float64(seg.y2)))
			seg_max_y := int(math.Max(float64(seg.y1), float64(seg.y2)))
			// Segment is strictly inside x-range and crosses y-range
			if seg.x1 > min_x && seg.x1 < max_x && seg_min_y < max_y && seg_max_y > min_y {
				return true
			}
		} else { // horizontal segment
			seg_min_x := int(math.Min(float64(seg.x1), float64(seg.x2)))
			seg_max_x := int(math.Max(float64(seg.x1), float64(seg.x2)))
			// Segment is strictly inside y-range and crosses x-range
			if seg.y1 > min_y && seg.y1 < max_y && seg_min_x < max_x && seg_max_x > min_x {
				return true
			}
		}
	}
	return false
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

	data := parse_input("input.txt")
	fmt.Println("Part 1 - Max Tiles Between Points:", part1(data))
	fmt.Println("Part 2 - Max Red and Green Tiles:", part2(data))
}
