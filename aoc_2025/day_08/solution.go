package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parse_input(input string) string {
	// Parse the input file to a raw string
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return string(data)
}

type point struct {
	x, y, z int
}

type edge struct {
	dist2 int64
	a, b  int
}

func parse_points(input string) []point {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	pts := make([]point, 0, len(lines))
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		pts = append(pts, point{x, y, z})
	}
	return pts
}

// build_edges generates all pairwise edges with squared distance
func build_edges(pts []point) []edge {
	n := len(pts)
	edges := make([]edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := int64(pts[i].x - pts[j].x)
			dy := int64(pts[i].y - pts[j].y)
			dz := int64(pts[i].z - pts[j].z)
			dist2 := dx*dx + dy*dy + dz*dz
			edges = append(edges, edge{dist2: dist2, a: i, b: j})
		}
	}
	return edges
}

func part1(input string, connections int) int {
	pts := parse_points(input)
	edges := build_edges(pts)

	// Sort by squared distance only; smallest first.
	sort.Slice(edges, func(i, j int) bool { return edges[i].dist2 < edges[j].dist2 })

	d := newDSU(len(pts))
	limit := connections
	if limit > len(edges) {
		limit = len(edges)
	}
	for i := 0; i < limit; i++ {
		d.union(edges[i].a, edges[i].b)
	}

	sizes := d.componentSizes()
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })
	return sizes[0] * sizes[1] * sizes[2]
}

func part2(input string, connections int) int {
	_ = connections // Ignore connections for part 2 in this example, limit = len(edges)
	pts := parse_points(input)
	edges := build_edges(pts)
	sort.Slice(edges, func(i, j int) bool { return edges[i].dist2 < edges[j].dist2 })

	d := newDSU(len(pts))
	components := len(pts)
	for _, e := range edges {
		if d.find(e.a) != d.find(e.b) {
			d.union(e.a, e.b) // union happened
			components--
			if components == 1 {
				// e is the edge that connected everything
				return pts[e.a].x * pts[e.b].x
			}
		}
	}
	return 0
}

func test() {
	// Test with example data from the puzzle

	//X-Y-Z Coordinates of Junction Boxes
	example_input_str := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

	example_shortest_connections := 10

	excpeted_part1 := 40
	excpeted_part2 := 25272

	result_part1 := part1(example_input_str, example_shortest_connections)
	if result_part1 != excpeted_part1 {
		fmt.Printf("Part 1 failed: got %d, expected %d, difference %d\n", result_part1, excpeted_part1, excpeted_part1-result_part1)
	} else {
		fmt.Println("Part 1 passed")
	}

	result_part2 := part2(example_input_str, example_shortest_connections)
	if result_part2 != excpeted_part2 {
		fmt.Printf("Part 2 failed: got %d, expected %d, difference %d\n", result_part2, excpeted_part2, excpeted_part2-result_part2)
	} else {
		fmt.Println("Part 2 passed")
	}
}

func main() {
	// Run tests
	test()

	// Example usage with real input (adjust connections as needed, e.g., 1000 per puzzle spec)
	data := parse_input("input.txt")
	fmt.Println("Part 1 - Product of top 3 sizes:", part1(data, 1000))
	fmt.Println("Part 2 - Product of top 3 sizes:", part2(data, 1000))
}
