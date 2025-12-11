package main

import (
	"fmt"
	"os"
	"sort"
)

// Range represents a fresh ID range with start and end (inclusive)
type Range struct {
	start int64
	end   int64
}

// parseInput reads and parses the input file in a single pass
// Uses direct byte parsing to avoid string allocations
func parseInput(filename string) ([]Range, []int64) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return parseBytes(data)
}

// parseBytes parses raw bytes directly without string conversions
// This is the hot path - optimized for minimal allocations
func parseBytes(data []byte) ([]Range, []int64) {
	n := len(data)

	// Find the blank line separator (two consecutive newlines)
	sepIdx := -1
	for i := 0; i < n-1; i++ {
		if data[i] == '\n' && data[i+1] == '\n' {
			sepIdx = i
			break
		}
	}
	if sepIdx == -1 {
		panic("invalid input: no separator found")
	}

	// Count lines in each section for pre-allocation
	rangeCount := 1
	for i := 0; i < sepIdx; i++ {
		if data[i] == '\n' {
			rangeCount++
		}
	}
	idCount := 0
	for i := sepIdx + 2; i < n; i++ {
		if data[i] == '\n' {
			idCount++
		}
	}
	// Handle last line if no trailing newline
	if n > 0 && data[n-1] != '\n' {
		idCount++
	}

	ranges := make([]Range, 0, rangeCount)
	ids := make([]int64, 0, idCount)

	// Parse ranges section
	i := 0
	for i < sepIdx {
		// Skip whitespace
		for i < sepIdx && (data[i] == ' ' || data[i] == '\r') {
			i++
		}
		if i >= sepIdx || data[i] == '\n' {
			i++
			continue
		}

		// Parse start number
		var start int64
		for i < sepIdx && data[i] >= '0' && data[i] <= '9' {
			start = start*10 + int64(data[i]-'0')
			i++
		}

		// Skip the '-' separator
		if i < sepIdx && data[i] == '-' {
			i++
		}

		// Parse end number
		var end int64
		for i < sepIdx && data[i] >= '0' && data[i] <= '9' {
			end = end*10 + int64(data[i]-'0')
			i++
		}

		ranges = append(ranges, Range{start: start, end: end})

		// Skip to next line
		for i < sepIdx && data[i] != '\n' {
			i++
		}
		i++ // skip newline
	}

	// Skip separator (blank line)
	i = sepIdx + 2

	// Parse IDs section
	for i < n {
		// Skip whitespace
		for i < n && (data[i] == ' ' || data[i] == '\r' || data[i] == '\n') {
			i++
		}
		if i >= n {
			break
		}

		// Parse ID number
		var id int64
		for i < n && data[i] >= '0' && data[i] <= '9' {
			id = id*10 + int64(data[i]-'0')
			i++
		}
		ids = append(ids, id)

		// Skip to next line
		for i < n && data[i] != '\n' {
			i++
		}
	}

	return ranges, ids
}

// part1Optimized counts how many ingredient IDs fall within fresh ranges
// Uses merged ranges with binary search for O(n log m) complexity
// where n = number of IDs and m = number of merged ranges
func part1Optimized(ranges []Range, ids []int64) int {
	if len(ranges) == 0 {
		return 0
	}

	// Sort ranges by start
	sortedRanges := make([]Range, len(ranges))
	copy(sortedRanges, ranges)
	sort.Slice(sortedRanges, func(i, j int) bool {
		return sortedRanges[i].start < sortedRanges[j].start
	})

	// Merge overlapping ranges first - this ensures each ID is counted at most once
	// and allows us to use binary search correctly
	merged := make([]Range, 0, len(sortedRanges))
	current := sortedRanges[0]

	for i := 1; i < len(sortedRanges); i++ {
		r := sortedRanges[i]
		if r.start <= current.end+1 {
			if r.end > current.end {
				current.end = r.end
			}
		} else {
			merged = append(merged, current)
			current = r
		}
	}
	merged = append(merged, current)

	count := 0
	for _, id := range ids {
		// Binary search: find the rightmost range where start <= id
		left, right := 0, len(merged)
		for left < right {
			mid := (left + right) / 2
			if merged[mid].start <= id {
				left = mid + 1
			} else {
				right = mid
			}
		}
		// Check if the found range contains the ID
		if left > 0 && id <= merged[left-1].end {
			count++
		}
	}
	return count
}

// part2Optimized counts unique fresh IDs by merging overlapping ranges
// Uses an optimized sort-and-merge algorithm for O(n log n) complexity
func part2Optimized(ranges []Range) int64 {
	if len(ranges) == 0 {
		return 0
	}

	// Sort ranges by start
	sortedRanges := make([]Range, len(ranges))
	copy(sortedRanges, ranges)
	sort.Slice(sortedRanges, func(i, j int) bool {
		return sortedRanges[i].start < sortedRanges[j].start
	})

	// Merge overlapping ranges in a single pass
	merged := make([]Range, 0, len(sortedRanges))
	current := sortedRanges[0]

	for i := 1; i < len(sortedRanges); i++ {
		r := sortedRanges[i]
		// Check for overlap or adjacency (ranges are inclusive, so adjacent means end+1 >= start)
		if r.start <= current.end+1 {
			// Merge: extend current range if needed
			if r.end > current.end {
				current.end = r.end
			}
		} else {
			// No overlap, save current and start new
			merged = append(merged, current)
			current = r
		}
	}
	merged = append(merged, current)

	// Sum up all unique IDs
	var total int64
	for _, r := range merged {
		total += r.end - r.start + 1
	}
	return total
}

// solve computes both parts efficiently
// Ranges are only sorted once for part2, and the merged result is reused conceptually
func solve(ranges []Range, ids []int64) (int, int64) {
	p1 := part1Optimized(ranges, ids)
	p2 := part2Optimized(ranges)
	return p1, p2
}

func test() {
	example_input := []byte(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

	ranges, ids := parseBytes(example_input)

	expected_part1 := 3
	var expected_part2 int64 = 14

	result_part1 := part1Optimized(ranges, ids)
	result_part2 := part2Optimized(ranges)

	if result_part1 != expected_part1 {
		fmt.Printf("Part 1 failed: got %d, expected %d\n", result_part1, expected_part1)
	} else {
		fmt.Println("Part 1 passed")
	}

	if result_part2 != expected_part2 {
		fmt.Printf("Part 2 failed: got %d, expected %d\n", result_part2, expected_part2)
	} else {
		fmt.Println("Part 2 passed")
	}
}

func main() {
	// Run tests
	test()

	// Parse input and solve
	ranges, ids := parseInput("../input.txt")
	p1, p2 := solve(ranges, ids)
	fmt.Println("Part 1 - Fresh ID Count:", p1)
	fmt.Println("Part 2 - Unique Fresh ID Count:", p2)
}
