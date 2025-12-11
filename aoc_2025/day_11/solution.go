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

func parse_graph(input string) map[string][]string {
	graph := make(map[string][]string)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}

		node := parts[0]
		connections := strings.Split(parts[1], " ")
		graph[node] = connections
	}

	return graph
}

func count_paths(current string, graph map[string][]string, memo map[string]int) int {
	// Base case: reached the output
	if current == "out" {
		return 1
	}

	// Check memo for already computed result
	if val, exists := memo[current]; exists {
		return val
	}

	// Count all paths from current node to "out"
	total_paths := 0
	if neighbors, exists := graph[current]; exists {
		for _, neighbor := range neighbors {
			total_paths += count_paths(neighbor, graph, memo)
		}
	}

	memo[current] = total_paths
	return total_paths
}

func part1(input string) int {
	graph := parse_graph(input)
	memo := make(map[string]int)
	return count_paths("you", graph, memo)
}

func count_paths_with_required(current string, graph map[string][]string, visited_dac bool, visited_fft bool, memo map[string]map[bool]map[bool]int) int {
	// Base case: reached the output
	if current == "out" {
		// Only count if we've visited both required nodes
		if visited_dac && visited_fft {
			return 1
		}
		return 0
	}

	// Initialize nested maps if needed
	if memo[current] == nil {
		memo[current] = make(map[bool]map[bool]int)
	}
	if memo[current][visited_dac] == nil {
		memo[current][visited_dac] = make(map[bool]int)
	}

	// Check memo
	if val, exists := memo[current][visited_dac][visited_fft]; exists {
		return val
	}

	// Update flags if current node is a required node
	new_visited_dac := visited_dac || current == "dac"
	new_visited_fft := visited_fft || current == "fft"

	// Count all paths from current node to "out" with required nodes
	total_paths := 0
	if neighbors, exists := graph[current]; exists {
		for _, neighbor := range neighbors {
			total_paths += count_paths_with_required(neighbor, graph, new_visited_dac, new_visited_fft, memo)
		}
	}

	memo[current][visited_dac][visited_fft] = total_paths
	return total_paths
}

func part2(input string) int {
	graph := parse_graph(input)
	memo := make(map[string]map[bool]map[bool]int)
	return count_paths_with_required("svr", graph, false, false, memo)
}

func test() {
	// Test with example data from the puzzle
	example_input_str := `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

	expected_part1 := 5

	result_part1 := part1(example_input_str)
	if result_part1 != expected_part1 {
		fmt.Printf("Part 1 failed: got %d, expected %d, difference %d\n", result_part1, expected_part1, expected_part1-result_part1)
	} else {
		fmt.Println("Part 1 passed")
	}

	// Test with part 2 example data
	example_input_str_part2 := `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`

	expected_part2 := 2
	result_part2 := part2(example_input_str_part2)
	if result_part2 != expected_part2 {
		fmt.Printf("Part 2 failed: got %d, expected %d\n", result_part2, expected_part2)
	} else {
		fmt.Println("Part 2 passed")
	}
}

func main() {
	// Run tests
	test()

	data := parse_input("input.txt")
	fmt.Println("Part 1 - Total Paths from you to out:", part1(data))
	fmt.Println("Part 2 - Total Paths from svr to out with dac and fft visited:", part2(data))
}
