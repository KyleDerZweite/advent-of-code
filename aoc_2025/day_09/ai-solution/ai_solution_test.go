package main

import (
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

// ==================== Original Solution ====================

type originalSegment struct {
	x1, y1, x2, y2 int
}

func originalPart1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	maxTiles := 0
	for _, line := range lines {
		lineX, _ := strconv.Atoi(strings.Split(line, ",")[0])
		lineY, _ := strconv.Atoi(strings.Split(line, ",")[1])

		for _, otherLine := range lines {
			otherLineX, _ := strconv.Atoi(strings.Split(otherLine, ",")[0])
			otherLineY, _ := strconv.Atoi(strings.Split(otherLine, ",")[1])

			if line == otherLine {
				continue
			}
			xDistance := math.Abs(float64(otherLineX-lineX)) + 1
			yDistance := math.Abs(float64(otherLineY-lineY)) + 1

			newMaxTiles := xDistance * yDistance

			if int(newMaxTiles) > maxTiles {
				maxTiles = int(newMaxTiles)
			}
		}
	}
	return int(maxTiles)
}

func originalPart2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	redTiles := []struct{ x, y int }{}
	for _, line := range lines {
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		redTiles = append(redTiles, struct{ x, y int }{x, y})
	}

	segments := []originalSegment{}
	for i := 0; i < len(redTiles); i++ {
		rt1 := redTiles[i]
		rt2 := redTiles[(i+1)%len(redTiles)]
		segments = append(segments, originalSegment{rt1.x, rt1.y, rt2.x, rt2.y})
	}

	maxArea := 0
	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			r1, r2 := redTiles[i], redTiles[j]
			if r1.x == r2.x || r1.y == r2.y {
				continue
			}

			minX := int(math.Min(float64(r1.x), float64(r2.x)))
			maxX := int(math.Max(float64(r1.x), float64(r2.x)))
			minY := int(math.Min(float64(r1.y), float64(r2.y)))
			maxY := int(math.Max(float64(r1.y), float64(r2.y)))

			if !originalIsInside(minX, minY, segments) || !originalIsInside(maxX, minY, segments) ||
				!originalIsInside(minX, maxY, segments) || !originalIsInside(maxX, maxY, segments) {
				continue
			}

			if originalRectangleIsCut(minX, minY, maxX, maxY, segments) {
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

func originalIsInside(x, y int, segments []originalSegment) bool {
	for _, seg := range segments {
		if seg.x1 == seg.x2 {
			minY := int(math.Min(float64(seg.y1), float64(seg.y2)))
			maxY := int(math.Max(float64(seg.y1), float64(seg.y2)))
			if x == seg.x1 && y >= minY && y <= maxY {
				return true
			}
		} else {
			minX := int(math.Min(float64(seg.x1), float64(seg.x2)))
			maxX := int(math.Max(float64(seg.x1), float64(seg.x2)))
			if y == seg.y1 && x >= minX && x <= maxX {
				return true
			}
		}
	}

	crossings := 0
	for _, seg := range segments {
		if seg.x1 == seg.x2 {
			minY := int(math.Min(float64(seg.y1), float64(seg.y2)))
			maxY := int(math.Max(float64(seg.y1), float64(seg.y2)))
			if seg.x1 > x && y > minY && y < maxY {
				crossings++
			}
		}
	}
	return crossings%2 == 1
}

func originalRectangleIsCut(minX, minY, maxX, maxY int, segments []originalSegment) bool {
	for _, seg := range segments {
		if seg.x1 == seg.x2 {
			segMinY := int(math.Min(float64(seg.y1), float64(seg.y2)))
			segMaxY := int(math.Max(float64(seg.y1), float64(seg.y2)))
			if seg.x1 > minX && seg.x1 < maxX && segMinY < maxY && segMaxY > minY {
				return true
			}
		} else {
			segMinX := int(math.Min(float64(seg.x1), float64(seg.x2)))
			segMaxX := int(math.Max(float64(seg.x1), float64(seg.x2)))
			if seg.y1 > minY && seg.y1 < maxY && segMinX < maxX && segMaxX > minX {
				return true
			}
		}
	}
	return false
}

// ==================== Test Data Loading ====================

func loadTestData() ([]byte, string) {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	return data, string(data)
}

// ==================== Benchmarks ====================

// Benchmark parsing only
func BenchmarkOriginalParse(b *testing.B) {
	_, inputStr := loadTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Split(strings.TrimSpace(inputStr), "\n")
	}
}

func BenchmarkOptimizedParse(b *testing.B) {
	data, _ := loadTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseBytes(data)
	}
}

// Benchmark Part 1 only
func BenchmarkOriginalPart1(b *testing.B) {
	_, inputStr := loadTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		originalPart1(inputStr)
	}
}

func BenchmarkOptimizedPart1(b *testing.B) {
	data, _ := loadTestData()
	redTiles := parseBytes(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1Optimized(redTiles)
	}
}

// Benchmark Part 2 only
func BenchmarkOriginalPart2(b *testing.B) {
	_, inputStr := loadTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		originalPart2(inputStr)
	}
}

func BenchmarkOptimizedPart2(b *testing.B) {
	data, _ := loadTestData()
	redTiles := parseBytes(data)
	segments := buildSegments(redTiles)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2Optimized(redTiles, segments)
	}
}

// Benchmark full solve (parse + both parts)
func BenchmarkOriginalFull(b *testing.B) {
	_, inputStr := loadTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		originalPart1(inputStr)
		originalPart2(inputStr)
	}
}

func BenchmarkOptimizedFull(b *testing.B) {
	data, _ := loadTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redTiles := parseBytes(data)
		segments := buildSegments(redTiles)
		solve(redTiles, segments)
	}
}

// ==================== Correctness Tests ====================

func TestCorrectness(t *testing.T) {
	data, inputStr := loadTestData()

	// Run both versions
	origP1 := originalPart1(inputStr)
	origP2 := originalPart2(inputStr)

	redTiles := parseBytes(data)
	segments := buildSegments(redTiles)
	optP1, optP2 := solve(redTiles, segments)

	// Compare Part 1 results
	if origP1 != optP1 {
		t.Errorf("Part 1 mismatch: original=%d, optimized=%d", origP1, optP1)
	} else {
		t.Logf("Part 1: both returned %d", origP1)
	}

	// Compare Part 2 results
	if origP2 != optP2 {
		t.Errorf("Part 2 mismatch: original=%d, optimized=%d", origP2, optP2)
	} else {
		t.Logf("Part 2: both returned %d", origP2)
	}
}

func TestExampleInput(t *testing.T) {
	exampleInput := []byte(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)

	redTiles := parseBytes(exampleInput)
	segments := buildSegments(redTiles)

	p1 := part1Optimized(redTiles)
	if p1 != 50 {
		t.Errorf("Part 1 example failed: got %d, expected 50", p1)
	}

	p2 := part2Optimized(redTiles, segments)
	if p2 != 24 {
		t.Errorf("Part 2 example failed: got %d, expected 24", p2)
	}
}
