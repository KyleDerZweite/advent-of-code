package main

import (
	"fmt"
	"os"
)

// segment represents a line segment connecting two red tiles
type segment struct {
	x1, y1, x2, y2 int
}

// point represents a coordinate
type point struct {
	x, y int
}

// parseInput reads the input file
func parseInput(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

// parseBytes parses raw bytes directly without string conversions
func parseBytes(data []byte) []point {
	points := make([]point, 0, 256)
	n := len(data)
	i := 0

	for i < n {
		// Skip whitespace
		for i < n && (data[i] == ' ' || data[i] == '\r' || data[i] == '\n') {
			i++
		}
		if i >= n {
			break
		}

		// Parse x
		x := 0
		for i < n && data[i] >= '0' && data[i] <= '9' {
			x = x*10 + int(data[i]-'0')
			i++
		}

		// Skip comma
		if i < n && data[i] == ',' {
			i++
		}

		// Parse y
		y := 0
		for i < n && data[i] >= '0' && data[i] <= '9' {
			y = y*10 + int(data[i]-'0')
			i++
		}

		points = append(points, point{x, y})
	}

	return points
}

// buildSegments creates segments connecting consecutive red tiles
func buildSegments(redTiles []point) []segment {
	segments := make([]segment, len(redTiles))
	for i := 0; i < len(redTiles); i++ {
		rt1 := redTiles[i]
		rt2 := redTiles[(i+1)%len(redTiles)]
		segments[i] = segment{rt1.x, rt1.y, rt2.x, rt2.y}
	}
	return segments
}

// minMax returns min and max of two integers
func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// part1Optimized finds the largest rectangle using any two red tiles as corners
func part1Optimized(redTiles []point) int {
	maxArea := 0
	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			r1, r2 := redTiles[i], redTiles[j]
			if r1.x == r2.x || r1.y == r2.y {
				continue
			}

			minX, maxX := minMax(r1.x, r2.x)
			minY, maxY := minMax(r1.y, r2.y)

			area := (maxX - minX + 1) * (maxY - minY + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// isInside checks if a point is inside or on the boundary of the polygon
func isInside(x, y int, segments []segment) bool {
	// Check if on boundary first
	for _, seg := range segments {
		if seg.x1 == seg.x2 { // vertical
			minY, maxY := minMax(seg.y1, seg.y2)
			if x == seg.x1 && y >= minY && y <= maxY {
				return true
			}
		} else { // horizontal
			minX, maxX := minMax(seg.x1, seg.x2)
			if y == seg.y1 && x >= minX && x <= maxX {
				return true
			}
		}
	}

	// Ray casting: count vertical segments to the right
	crossings := 0
	for _, seg := range segments {
		if seg.x1 == seg.x2 { // vertical segment
			minY, maxY := minMax(seg.y1, seg.y2)
			if seg.x1 > x && y > minY && y < maxY {
				crossings++
			}
		}
	}
	return crossings%2 == 1
}

// rectangleIsCut checks if any segment passes through the interior of the rectangle
func rectangleIsCut(minX, minY, maxX, maxY int, segments []segment) bool {
	for _, seg := range segments {
		if seg.x1 == seg.x2 { // vertical segment
			segMinY, segMaxY := minMax(seg.y1, seg.y2)
			// Segment is strictly inside x-range and crosses y-range
			if seg.x1 > minX && seg.x1 < maxX && segMinY < maxY && segMaxY > minY {
				return true
			}
		} else { // horizontal segment
			segMinX, segMaxX := minMax(seg.x1, seg.x2)
			// Segment is strictly inside y-range and crosses x-range
			if seg.y1 > minY && seg.y1 < maxY && segMinX < maxX && segMaxX > minX {
				return true
			}
		}
	}
	return false
}

// part2Optimized finds the largest rectangle with red corners containing only red/green tiles
func part2Optimized(redTiles []point, segments []segment) int {
	maxArea := 0
	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			r1, r2 := redTiles[i], redTiles[j]
			if r1.x == r2.x || r1.y == r2.y {
				continue
			}

			minX, maxX := minMax(r1.x, r2.x)
			minY, maxY := minMax(r1.y, r2.y)

			// Check all 4 corners are inside or on boundary
			if !isInside(minX, minY, segments) || !isInside(maxX, minY, segments) ||
				!isInside(minX, maxY, segments) || !isInside(maxX, maxY, segments) {
				continue
			}

			// Check that no segment crosses through our rectangle
			if rectangleIsCut(minX, minY, maxX, maxY, segments) {
				continue
			}

			area := (maxX - minX + 1) * (maxY - minY + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// solve runs both parts and returns results
func solve(redTiles []point, segments []segment) (int, int) {
	return part1Optimized(redTiles), part2Optimized(redTiles, segments)
}

func main() {
	// Test with example
	exampleData := []byte(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)

	redTiles := parseBytes(exampleData)
	segments := buildSegments(redTiles)

	p1 := part1Optimized(redTiles)
	p2 := part2Optimized(redTiles, segments)

	if p1 != 50 {
		fmt.Printf("Part 1 failed: got %d, expected 50\n", p1)
	} else {
		fmt.Println("Part 1 passed")
	}

	if p2 != 24 {
		fmt.Printf("Part 2 failed: got %d, expected 24\n", p2)
	} else {
		fmt.Println("Part 2 passed")
	}

	// Run on actual input
	data := parseInput("../input.txt")
	redTiles = parseBytes(data)
	segments = buildSegments(redTiles)

	p1, p2 = solve(redTiles, segments)
	fmt.Println("Part 1 - Max Tiles Between Points:", p1)
	fmt.Println("Part 2 - Max Red and Green Tiles:", p2)
}
