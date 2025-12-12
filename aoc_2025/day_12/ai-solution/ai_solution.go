// Package main solves Advent of Code 2025 - Day 12: Christmas Tree Farm
//
// This puzzle is a 2D bin packing problem with polyomino shapes.
// Given a set of irregularly-shaped "presents" and rectangular regions,
// we must determine how many regions can fit all their assigned presents.
//
// Algorithm Overview:
// - Parse shape definitions (polyominoes) and region specifications
// - For each region, use backtracking to try placing all required shapes
// - Shapes can be rotated (4 orientations) and flipped (2 mirrors) = up to 8 unique orientations
// - Optimization: Always fill the "first empty cell" to avoid redundant placements
// - Allow empty cells when total shape area < grid area
//
// Time Complexity: O(R * (O^S * W*H)) where R=regions, O=orientations, S=shapes, W*H=grid size
// Space Complexity: O(W*H) for the grid + O(S*O) for precomputed orientations
//
// Note: This solution was created with AI assistance as part of learning Go.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// =============================================================================
// Input Parsing
// =============================================================================

// parse_input reads the puzzle input from a file
func parse_input(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// cell represents a single coordinate in a shape or grid
type cell struct {
	row int
	col int
}

// shape represents a polyomino as a list of cell coordinates
// Coordinates are relative offsets from an anchor point
type shape []cell

// region represents a rectangular area and the shapes that must fit in it
type region struct {
	width  int   // width of the region (columns)
	height int   // height of the region (rows)
	counts []int // count of each shape type needed (indexed by shape ID)
}

// parse_shapes extracts all shape definitions from input
// Shape format: "N:\n###\n##.\n##." where N is the index, # is filled, . is empty
func parse_shapes(input string) []shape {
	shapes := []shape{}
	sections := strings.Split(input, "\n\n")

	for _, section := range sections {
		section = strings.TrimSpace(section)
		if section == "" {
			continue
		}

		// Shape definitions start with "N:" where N is a digit
		if len(section) > 0 && section[0] >= '0' && section[0] <= '9' &&
			strings.Contains(section, ":") && strings.Contains(section, "#") {

			lines := strings.Split(section, "\n")
			cells := shape{}

			// Parse each row after the "N:" header
			for row_idx, line := range lines[1:] {
				for col_idx, ch := range line {
					if ch == '#' {
						cells = append(cells, cell{row: row_idx, col: col_idx})
					}
				}
			}

			if len(cells) > 0 {
				shapes = append(shapes, cells)
			}
		}
	}
	return shapes
}

// parse_regions extracts region definitions from input
// Region format: "WxH: c0 c1 c2 ..." where WxH is dimensions, cN is count of shape N
func parse_regions(input string, num_shapes int) []region {
	regions := []region{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Region lines contain "x" and ":" but no "#" (distinguishes from shapes)
		if strings.Contains(line, "x") && strings.Contains(line, ":") && !strings.Contains(line, "#") {
			parts := strings.SplitN(line, ":", 2)
			dims := strings.Split(strings.TrimSpace(parts[0]), "x")

			width, _ := strconv.Atoi(dims[0])
			height, _ := strconv.Atoi(dims[1])

			counts := make([]int, num_shapes)
			count_strs := strings.Fields(strings.TrimSpace(parts[1]))

			for i, cs := range count_strs {
				if i < num_shapes {
					counts[i], _ = strconv.Atoi(cs)
				}
			}

			regions = append(regions, region{
				width:  width,
				height: height,
				counts: counts,
			})
		}
	}
	return regions
}

// =============================================================================
// Shape Transformations
// =============================================================================

// normalize_shape shifts a shape so its minimum row and column are both 0
// This creates a canonical form for comparing shapes
func normalize_shape(s shape) shape {
	if len(s) == 0 {
		return s
	}

	// Find minimum row and column
	min_row, min_col := s[0].row, s[0].col
	for _, c := range s {
		if c.row < min_row {
			min_row = c.row
		}
		if c.col < min_col {
			min_col = c.col
		}
	}

	// Shift all cells
	normalized := make(shape, len(s))
	for i, c := range s {
		normalized[i] = cell{row: c.row - min_row, col: c.col - min_col}
	}
	return normalized
}

// rotate_shape rotates a shape 90 degrees clockwise
// Transformation: (row, col) -> (col, -row)
func rotate_shape(s shape) shape {
	rotated := make(shape, len(s))
	for i, c := range s {
		rotated[i] = cell{row: c.col, col: -c.row}
	}
	return normalize_shape(rotated)
}

// flip_shape flips a shape horizontally (mirror across vertical axis)
// Transformation: (row, col) -> (row, -col)
func flip_shape(s shape) shape {
	flipped := make(shape, len(s))
	for i, c := range s {
		flipped[i] = cell{row: c.row, col: -c.col}
	}
	return normalize_shape(flipped)
}

// shape_to_string creates a unique string representation for deduplication
func shape_to_string(s shape) string {
	// Create sorted list of "row,col" strings
	cells := make([]string, len(s))
	for i, c := range s {
		cells[i] = fmt.Sprintf("%d,%d", c.row, c.col)
	}

	// Simple bubble sort (shapes are small, typically < 10 cells)
	for i := 0; i < len(cells); i++ {
		for j := i + 1; j < len(cells); j++ {
			if cells[j] < cells[i] {
				cells[i], cells[j] = cells[j], cells[i]
			}
		}
	}
	return strings.Join(cells, ";")
}

// get_all_orientations generates all unique rotations and flips of a shape
// For each orientation, generates versions anchored at each cell position
// This allows the solver to place shapes such that any cell covers a target position
func get_all_orientations(s shape) []shape {
	orientations := []shape{}
	seen := make(map[string]bool)

	current := normalize_shape(s)

	// Try 2 flips (original + horizontally flipped)
	for flip := 0; flip < 2; flip++ {
		// Try 4 rotations (0°, 90°, 180°, 270°)
		for rot := 0; rot < 4; rot++ {
			// For each cell in the shape, create a version anchored at that cell
			// This means that cell becomes (0,0), allowing placement at any grid position
			for _, anchor := range current {
				shifted := make(shape, len(current))
				for i, c := range current {
					shifted[i] = cell{
						row: c.row - anchor.row,
						col: c.col - anchor.col,
					}
				}

				key := shape_to_string(shifted)
				if !seen[key] {
					seen[key] = true
					orientations = append(orientations, shifted)
				}
			}
			current = rotate_shape(current)
		}
		current = flip_shape(current)
	}
	return orientations
}

// =============================================================================
// Grid Operations
// =============================================================================

// can_place checks if a shape can be placed at (row, col) on the grid
// Returns false if any cell would be out of bounds or already occupied
func can_place(grid [][]bool, s shape, row, col, width, height int) bool {
	for _, c := range s {
		r, co := row+c.row, col+c.col
		if r < 0 || r >= height || co < 0 || co >= width {
			return false
		}
		if grid[r][co] {
			return false
		}
	}
	return true
}

// place_shape marks all cells of a shape as occupied on the grid
func place_shape(grid [][]bool, s shape, row, col int) {
	for _, c := range s {
		grid[row+c.row][col+c.col] = true
	}
}

// remove_shape marks all cells of a shape as unoccupied on the grid
func remove_shape(grid [][]bool, s shape, row, col int) {
	for _, c := range s {
		grid[row+c.row][col+c.col] = false
	}
}

// =============================================================================
// Backtracking Solver
// =============================================================================

// solve attempts to place all shapes using backtracking
//
// Strategy:
// 1. Find the first empty cell in the grid (scanning top-to-bottom, left-to-right)
// 2. Try placing each unused shape at that position (with all orientations)
// 3. If successful, recursively try to place remaining shapes
// 4. If no shape fits, try "skipping" the cell (if empty cells are allowed)
// 5. Backtrack if we reach a dead end
//
// The "first empty cell" strategy is crucial for efficiency:
// - It prevents trying the same placement in different orders
// - It prunes the search space significantly
func solve(grid [][]bool, shapes_to_place [][]shape, used []bool, width, height, empty_allowed int) bool {
	// Base case: all shapes placed successfully
	all_used := true
	for _, u := range used {
		if !u {
			all_used = false
			break
		}
	}
	if all_used {
		return true
	}

	// Find first empty cell (top-to-bottom, left-to-right)
	first_row, first_col := -1, -1
	for row := 0; row < height && first_row == -1; row++ {
		for col := 0; col < width; col++ {
			if !grid[row][col] {
				first_row, first_col = row, col
				break
			}
		}
	}

	// No empty cell but shapes remain - impossible
	if first_row == -1 {
		return false
	}

	// Try placing each unused shape at the first empty cell
	for shape_idx := 0; shape_idx < len(shapes_to_place); shape_idx++ {
		if used[shape_idx] {
			continue
		}

		for _, oriented := range shapes_to_place[shape_idx] {
			if can_place(grid, oriented, first_row, first_col, width, height) {
				// Place shape and mark as used
				place_shape(grid, oriented, first_row, first_col)
				used[shape_idx] = true

				// Recurse
				if solve(grid, shapes_to_place, used, width, height, empty_allowed) {
					return true
				}

				// Backtrack
				used[shape_idx] = false
				remove_shape(grid, oriented, first_row, first_col)
			}
		}
	}

	// If no shape fits, try skipping this cell (leaving it empty)
	// This is only allowed when total shape area < grid area
	if empty_allowed > 0 {
		grid[first_row][first_col] = true // Mark as "filled" temporarily
		if solve(grid, shapes_to_place, used, width, height, empty_allowed-1) {
			return true
		}
		grid[first_row][first_col] = false // Backtrack
	}

	return false
}

// can_fit_all determines if all required shapes can fit in a region
func can_fit_all(width, height int, shapes []shape, counts []int) bool {
	// Build list of shapes to place with precomputed orientations
	shapes_to_place := [][]shape{}
	total_cells := 0

	for shape_idx, count := range counts {
		if count > 0 {
			orientations := get_all_orientations(shapes[shape_idx])
			for i := 0; i < count; i++ {
				shapes_to_place = append(shapes_to_place, orientations)
				total_cells += len(shapes[shape_idx])
			}
		}
	}

	// No shapes to place - trivially true
	if len(shapes_to_place) == 0 {
		return true
	}

	// Quick check: total shape area must not exceed grid area
	grid_area := width * height
	if total_cells > grid_area {
		return false
	}

	// Create empty grid
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}

	// Track which shapes have been placed
	used := make([]bool, len(shapes_to_place))

	// Calculate how many cells can remain empty
	empty_allowed := grid_area - total_cells

	return solve(grid, shapes_to_place, used, width, height, empty_allowed)
}

// =============================================================================
// Solution
// =============================================================================

// part1 counts how many regions can fit all their required presents
func part1(input string) int {
	shapes := parse_shapes(input)
	regions := parse_regions(input, len(shapes))

	count := 0
	for _, r := range regions {
		if can_fit_all(r.width, r.height, shapes, r.counts) {
			count++
		}
	}
	return count
}

// =============================================================================
// Testing & Main
// =============================================================================

func test() {
	example_input := `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`

	expected := 2

	result := part1(example_input)
	if result != expected {
		fmt.Printf("\n❌ Test failed: got %d, expected %d\n", result, expected)
	} else {
		fmt.Printf("\n✅ Test passed: %d regions can fit all presents\n", result)
	}
}

func main() {
	test()

	data := parse_input("../input.txt")
	fmt.Println("\n🎄 Part 1 - Valid regions:", part1(data))
	fmt.Println("⭐ Part 2 is awarded automatically after completing Part 1!")
}
