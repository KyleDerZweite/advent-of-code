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
	// Accept uppercase letters, lowercase letters, and digits 0-9
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}

// gcd returns greatest common divisor of two non-zero integers.
func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	if a == 0 {
		return 1
	}
	return a
}

func part1(input string) int {
	// Implement the logic for part 1 of the puzzle here
	unique_antinote_locations := 0
	locations_set := make(map[string]bool)
	lines := strings.Split(input, "\n")

	// Collect all antenna positions grouped by character
	type point struct{ x, y int }
	positions := make(map[rune][]point)
	for y, line := range lines {
		for x, char := range line {
			if is_letter_digit(char) {
				positions[char] = append(positions[char], point{x, y})
			}
		}
	}

	height := len(lines)

	// For each character, produce antinode positions from each pair of antennas
	for _, pts := range positions {
		n := len(pts)
		if n < 2 {
			continue
		}
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				ax, ay := pts[i].x, pts[i].y
				bx, by := pts[j].x, pts[j].y

				// Two antinode positions: 2*B - A and 2*A - B
				x1, y1 := 2*bx-ax, 2*by-ay
				x2, y2 := 2*ax-bx, 2*ay-by

				// Bound checks
				if y1 >= 0 && y1 < height && x1 >= 0 && x1 < len(lines[y1]) {
					key := fmt.Sprintf("%d,%d", x1, y1)
					locations_set[key] = true
				}
				if y2 >= 0 && y2 < height && x2 >= 0 && x2 < len(lines[y2]) {
					key := fmt.Sprintf("%d,%d", x2, y2)
					locations_set[key] = true
				}
			}
		}
	}

	unique_antinote_locations = len(locations_set)
	return unique_antinote_locations
}

func part2(input string) int {
	unique_antinote_locations := 0
	locations_set := make(map[string]bool)
	lines := strings.Split(input, "\n")

	// Collect all antenna positions grouped by character
	type point struct{ x, y int }
	positions := make(map[rune][]point)
	for y, line := range lines {
		for x, char := range line {
			if is_letter_digit(char) {
				positions[char] = append(positions[char], point{x, y})
			}
		}
	}

	height := len(lines)

	// For each character, mark every in-bounds point collinear with any pair of its antennas
	for _, pts := range positions {
		n := len(pts)
		if n < 2 {
			continue
		}
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				ax, ay := pts[i].x, pts[i].y
				bx, by := pts[j].x, pts[j].y
				dx, dy := bx-ax, by-ay
				stepGCD := gcd(dx, dy)
				sx, sy := dx/stepGCD, dy/stepGCD

				// Move backward from A to the first in-bounds point on this line
				startX, startY := ax, ay
				for {
					nx, ny := startX-sx, startY-sy
					if ny >= 0 && ny < height && nx >= 0 && nx < len(lines[ny]) {
						startX, startY = nx, ny
						continue
					}
					break
				}

				// Sweep forward marking all positions on the line until out of bounds
				xCur, yCur := startX, startY
				for {
					if yCur < 0 || yCur >= height || xCur < 0 || xCur >= len(lines[yCur]) {
						break
					}
					key := fmt.Sprintf("%d,%d", xCur, yCur)
					locations_set[key] = true
					xCur += sx
					yCur += sy
				}
			}
		}
	}

	unique_antinote_locations = len(locations_set)
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
	excpeted_part2 := 34

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
	fmt.Println("Part 2 - Unique antinode locations (till bounds):", part2(data))
}
