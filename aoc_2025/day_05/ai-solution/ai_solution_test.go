package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// Original solution functions for comparison
type OriginalRange struct {
	start int64
	end   int64
}

func originalParseString(input string) ([]OriginalRange, []int64) {
	fresh_id_ranges := []OriginalRange{}
	ingredient_ids := []int64{}

	fresh_id_lines := strings.Split(strings.TrimSpace(strings.Split(input, "\n\n")[0]), "\n")
	ingredient_id_lines := strings.Split(strings.TrimSpace(strings.Split(input, "\n\n")[1]), "\n")

	for _, line := range fresh_id_lines {
		parts := strings.Split(strings.TrimSpace(line), "-")
		fresh_ingredient_id_start, _ := strconv.ParseInt(parts[0], 10, 64)
		fresh_ingredient_id_end, _ := strconv.ParseInt(parts[1], 10, 64)
		fresh_id_ranges = append(fresh_id_ranges, OriginalRange{start: fresh_ingredient_id_start, end: fresh_ingredient_id_end})
	}

	for _, line := range ingredient_id_lines {
		ingredient_id, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		ingredient_ids = append(ingredient_ids, ingredient_id)
	}
	return fresh_id_ranges, ingredient_ids
}

func originalCheckIdBetweenRange(check_id int64, range_start int64, range_end int64) bool {
	return check_id >= range_start && check_id <= range_end
}

func originalPart1(fresh_id_ranges []OriginalRange, ingredient_ids []int64) int {
	fresh_id_count := 0
	for _, ingredient_id := range ingredient_ids {
		for _, r := range fresh_id_ranges {
			if originalCheckIdBetweenRange(ingredient_id, r.start, r.end) {
				fresh_id_count++
				break
			}
		}
	}
	return fresh_id_count
}

func originalPart2(fresh_id_ranges []OriginalRange, ingredient_ids []int64) int64 {
	_ = ingredient_ids

	unique_fresh_id_ranges := []OriginalRange{}
	for _, r := range fresh_id_ranges {
		new_range := r
		non_overlapping := []OriginalRange{}

		for _, ur := range unique_fresh_id_ranges {
			if !(new_range.end < ur.start || new_range.start > ur.end) {
				if ur.start < new_range.start {
					new_range.start = ur.start
				}
				if ur.end > new_range.end {
					new_range.end = ur.end
				}
			} else {
				non_overlapping = append(non_overlapping, ur)
			}
		}
		non_overlapping = append(non_overlapping, new_range)
		unique_fresh_id_ranges = non_overlapping
	}

	var unique_fresh_id_count int64 = 0
	for _, r := range unique_fresh_id_ranges {
		unique_fresh_id_count += r.end - r.start + 1
	}
	return unique_fresh_id_count
}

// Load test data
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
		originalParseString(inputStr)
	}
}

func BenchmarkOptimizedParse(b *testing.B) {
	data, _ := loadTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseBytes(data)
	}
}

// Benchmark Part 1 only (excluding parse)
func BenchmarkOriginalPart1(b *testing.B) {
	_, inputStr := loadTestData()
	ranges, ids := originalParseString(inputStr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		originalPart1(ranges, ids)
	}
}

func BenchmarkOptimizedPart1(b *testing.B) {
	data, _ := loadTestData()
	ranges, ids := parseBytes(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1Optimized(ranges, ids)
	}
}

// Benchmark Part 2 only (excluding parse)
func BenchmarkOriginalPart2(b *testing.B) {
	_, inputStr := loadTestData()
	ranges, ids := originalParseString(inputStr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		originalPart2(ranges, ids)
	}
}

func BenchmarkOptimizedPart2(b *testing.B) {
	data, _ := loadTestData()
	ranges, _ := parseBytes(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2Optimized(ranges)
	}
}

// Benchmark full solve (parse + both parts)
func BenchmarkOriginalFull(b *testing.B) {
	_, inputStr := loadTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ranges, ids := originalParseString(inputStr)
		originalPart1(ranges, ids)
		originalPart2(ranges, ids)
	}
}

func BenchmarkOptimizedFull(b *testing.B) {
	data, _ := loadTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ranges, ids := parseBytes(data)
		solve(ranges, ids)
	}
}

// ==================== Correctness Tests ====================

func TestCorrectness(t *testing.T) {
	data, inputStr := loadTestData()

	// Parse with both methods
	origRanges, origIds := originalParseString(inputStr)
	optRanges, optIds := parseBytes(data)

	// Verify parse results match
	if len(origRanges) != len(optRanges) {
		t.Fatalf("Range count mismatch: original=%d, optimized=%d", len(origRanges), len(optRanges))
	}
	if len(origIds) != len(optIds) {
		t.Fatalf("ID count mismatch: original=%d, optimized=%d", len(origIds), len(optIds))
	}

	// Sort both range slices for comparison (order may differ)
	sort.Slice(origRanges, func(i, j int) bool {
		return origRanges[i].start < origRanges[j].start
	})
	sort.Slice(optRanges, func(i, j int) bool {
		return optRanges[i].start < optRanges[j].start
	})

	for i := range origRanges {
		if origRanges[i].start != optRanges[i].start || origRanges[i].end != optRanges[i].end {
			t.Fatalf("Range mismatch at %d: original={%d,%d}, optimized={%d,%d}",
				i, origRanges[i].start, origRanges[i].end, optRanges[i].start, optRanges[i].end)
		}
	}

	// Compare Part 1 results
	origP1 := originalPart1(origRanges, origIds)
	optP1 := part1Optimized(optRanges, optIds)
	if origP1 != optP1 {
		t.Errorf("Part 1 mismatch: original=%d, optimized=%d", origP1, optP1)
	}

	// Compare Part 2 results
	origP2 := originalPart2(origRanges, origIds)
	optP2 := part2Optimized(optRanges)
	if origP2 != optP2 {
		t.Errorf("Part 2 mismatch: original=%d, optimized=%d", origP2, optP2)
	}
}

func TestExampleInput(t *testing.T) {
	exampleInput := []byte(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

	ranges, ids := parseBytes(exampleInput)

	p1 := part1Optimized(ranges, ids)
	if p1 != 3 {
		t.Errorf("Part 1 example failed: got %d, expected 3", p1)
	}

	p2 := part2Optimized(ranges)
	if p2 != 14 {
		t.Errorf("Part 2 example failed: got %d, expected 14", p2)
	}
}
