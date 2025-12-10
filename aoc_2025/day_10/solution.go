package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parse_input(input string) string {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return string(data)
}

type machine struct {
	indicator_lights []bool
	current_state    []bool
	buttons          [][]int
	joltages         []int
	current_joltages []int
}

func (m *machine) init_state() {
	m.current_state = make([]bool, len(m.indicator_lights))
	m.current_joltages = make([]int, len(m.joltages))
	for i := range m.current_joltages {
		m.current_joltages[i] = 0
	}
	// Start with all lights off
}

func (m *machine) press_button(index int) {
	if index < 0 || index >= len(m.buttons) {
		return
	}
	// Toggle each light that this button affects
	for _, light_idx := range m.buttons[index] {
		if light_idx >= 0 && light_idx < len(m.current_state) {
			m.current_state[light_idx] = !m.current_state[light_idx]
			m.current_joltages[light_idx] += 1
		}
	}
}

func (m *machine) is_correct() bool {
	if len(m.current_state) != len(m.indicator_lights) {
		return false
	}
	for i := range m.indicator_lights {
		if m.current_state[i] != m.indicator_lights[i] {
			return false
		}
	}
	return true
}

func (m *machine) is_joltage_correct() bool {
	if len(m.current_joltages) != len(m.joltages) {
		return false
	}
	for i := range m.joltages {
		if m.current_joltages[i] != m.joltages[i] {
			return false
		}
	}
	return true
}

func parse_machines(input string) []machine {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	machines := []machine{}

	// Regex patterns
	indicator_re := regexp.MustCompile(`\[([.#]+)\]`)
	button_re := regexp.MustCompile(`\(([0-9,]+)\)`)
	joltage_re := regexp.MustCompile(`\{([0-9,]+)\}`)

	for _, line := range lines {
		m := machine{}

		// Parse indicator lights [.##.]
		if match := indicator_re.FindStringSubmatch(line); match != nil {
			for _, c := range match[1] {
				m.indicator_lights = append(m.indicator_lights, c == '#')
			}
		}

		// Parse buttons (0,2,3,4) (2,3) etc.
		button_matches := button_re.FindAllStringSubmatch(line, -1)
		for _, match := range button_matches {
			var button []int
			for _, numStr := range strings.Split(match[1], ",") {
				num, _ := strconv.Atoi(numStr)
				button = append(button, num)
			}
			m.buttons = append(m.buttons, button)
		}

		// Parse joltages {3,5,4,7}
		if match := joltage_re.FindStringSubmatch(line); match != nil {
			for _, numStr := range strings.Split(match[1], ",") {
				num, _ := strconv.Atoi(numStr)
				m.joltages = append(m.joltages, num)
			}
		}

		machines = append(machines, m)
	}

	return machines
}

func recursive_solve_part1(m *machine, depth int, max_depth int) bool {
	if m.is_correct() {
		return true
	}
	if depth >= max_depth {
		return false
	}

	for i := range m.buttons {
		m.press_button(i)
		if recursive_solve_part1(m, depth+1, max_depth) {
			return true
		}
		m.press_button(i) // undo
	}
	return false
}

func solve_machine_part1(m *machine) int {
	m.init_state()

	if m.is_correct() {
		return 0
	}

	for max_depth := 1; max_depth <= 15; max_depth++ {
		m.init_state() // Reset before each attempt
		if recursive_solve_part1(m, 0, max_depth) {
			return max_depth
		}
	}

	return -1
}

func part1(input string) int {
	machines := parse_machines(input)

	min_button_presses := 0

	for i := range machines {
		m := &machines[i]
		result := solve_machine_part1(m)
		if result != -1 {
			min_button_presses += result
		}
	}

	return min_button_presses
}

func (m *machine) press_button_joltage(index int) {
	if index < 0 || index >= len(m.buttons) {
		return
	}
	// Increase each counter that this button affects
	for _, counter_idx := range m.buttons[index] {
		if counter_idx >= 0 && counter_idx < len(m.current_joltages) {
			m.current_joltages[counter_idx] += 1
		}
	}
}

func (m *machine) is_joltage_over() bool {
	for i := range m.joltages {
		if m.current_joltages[i] > m.joltages[i] {
			return true
		}
	}
	return false
}

func solve_machine_part2_gaussian(m *machine) int {
	num_counters := len(m.joltages)
	num_buttons := len(m.buttons)

	matrix := make([][]float64, num_counters)
	for i := 0; i < num_counters; i++ {
		matrix[i] = make([]float64, num_buttons+1)
		matrix[i][num_buttons] = float64(m.joltages[i]) // target
	}

	// Fill in coefficients
	for j, button := range m.buttons {
		for _, counter_idx := range button {
			if counter_idx >= 0 && counter_idx < num_counters {
				matrix[counter_idx][j] = 1.0
			}
		}
	}

	// Gaussian elimination (row reduction)
	pivot_row := 0
	pivot_cols := make([]int, 0) // Track which columns are pivot columns

	for col := 0; col < num_buttons && pivot_row < num_counters; col++ {
		// Find pivot
		max_row := pivot_row
		for row := pivot_row + 1; row < num_counters; row++ {
			if absFloat(matrix[row][col]) > absFloat(matrix[max_row][col]) {
				max_row = row
			}
		}

		if absFloat(matrix[max_row][col]) < 1e-9 {
			continue // No pivot in this column
		}

		// Swap rows
		matrix[pivot_row], matrix[max_row] = matrix[max_row], matrix[pivot_row]

		// Scale pivot row
		scale := matrix[pivot_row][col]
		for j := col; j <= num_buttons; j++ {
			matrix[pivot_row][j] /= scale
		}

		// Eliminate column in other rows
		for row := 0; row < num_counters; row++ {
			if row != pivot_row && absFloat(matrix[row][col]) > 1e-9 {
				factor := matrix[row][col]
				for j := col; j <= num_buttons; j++ {
					matrix[row][j] -= factor * matrix[pivot_row][j]
				}
			}
		}

		pivot_cols = append(pivot_cols, col)
		pivot_row++
	}

	// Check for inconsistency (row of form [0 0 ... 0 | nonzero])
	for row := pivot_row; row < num_counters; row++ {
		if absFloat(matrix[row][num_buttons]) > 1e-9 {
			return -1 // No solution
		}
	}

	// Now we have a reduced row echelon form
	// Free variables are columns not in pivotCols
	free_vars := make([]int, 0)
	pivot_set := make(map[int]bool)
	for _, col := range pivot_cols {
		pivot_set[col] = true
	}
	for col := 0; col < num_buttons; col++ {
		if !pivot_set[col] {
			free_vars = append(free_vars, col)
		}
	}

	return searchMinPresses(matrix, pivot_cols, free_vars, num_buttons, num_counters)
}

func absFloat(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// Search for minimum total presses
func searchMinPresses(matrix [][]float64, pivot_cols []int, free_vars []int, num_buttons int, num_counters int) int {
	// Determine reasonable upper bounds for free variables
	max_target := 0.0
	for row := 0; row < num_counters; row++ {
		if matrix[row][num_buttons] > max_target {
			max_target = matrix[row][num_buttons]
		}
	}

	max_val := int(max_target) + 1

	num_free := len(free_vars)
	if num_free == 0 {
		// Unique solution - check if valid
		solution := make([]float64, num_buttons)
		for i, col := range pivot_cols {
			solution[col] = matrix[i][num_buttons]
		}
		sum := 0.0
		for _, v := range solution {
			if v < -1e-9 || absFloat(v-float64(int(v+0.5))) > 1e-6 {
				return -1 // Not a non-negative integer solution
			}
			sum += float64(int(v + 0.5))
		}
		return int(sum + 0.5)
	}

	// Search over all free variable assignments and find minimum sum
	free_values := make([]int, num_free)
	min_sum := -1

	var search func(idx int)
	search = func(idx int) {
		if idx == num_free {
			// Evaluate this assignment
			solution := make([]float64, num_buttons)
			for i := 0; i < num_free; i++ {
				solution[free_vars[i]] = float64(free_values[i])
			}

			// Compute pivot variables from matrix
			for i, col := range pivot_cols {
				val := matrix[i][num_buttons]
				for j := 0; j < num_buttons; j++ {
					if j != col {
						val -= matrix[i][j] * solution[j]
					}
				}
				solution[col] = val
			}

			// Check validity and compute sum
			sum := 0
			for _, v := range solution {
				if v < -1e-9 {
					return // Negative value - invalid
				}
				rounded := int(v + 0.5)
				if absFloat(v-float64(rounded)) > 1e-6 {
					return // Not an integer - invalid
				}
				sum += rounded
			}

			if min_sum == -1 || sum < min_sum {
				min_sum = sum
			}
			return
		}

		for v := 0; v <= max_val; v++ {
			free_values[idx] = v
			search(idx + 1)
		}
	}

	search(0)
	return min_sum
}

func recursive_solve_part2(m *machine, depth int, max_depth int) bool {
	if m.is_joltage_correct() {
		return true
	}
	if depth >= max_depth {
		return false
	}
	if m.is_joltage_over() {
		return false // Pruning: can't decrease counters
	}

	for i := range m.buttons {
		m.press_button_joltage(i)
		if recursive_solve_part2(m, depth+1, max_depth) {
			return true
		}
		// Undo: decrease counters back
		for _, counter_idx := range m.buttons[i] {
			if counter_idx >= 0 && counter_idx < len(m.current_joltages) {
				m.current_joltages[counter_idx] -= 1
			}
		}
	}
	return false
}

func solve_machine_part2(m *machine) int {
	return solve_machine_part2_gaussian(m)
}

func part2(input string) int {
	machines := parse_machines(input)

	min_button_presses := 0

	for i := range machines {
		m := &machines[i]
		result := solve_machine_part2(m)
		if result != -1 {
			min_button_presses += result
		}
	}

	return min_button_presses
}

func test() {
	// Test with example data from the puzzle
	example_input_str := `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

	expected_part1 := 7
	expected_part2 := 33

	result_part1 := part1(example_input_str)
	if result_part1 != expected_part1 {
		fmt.Printf("Part 1 failed: got %d, expected %d, difference %d\n", result_part1, expected_part1, expected_part1-result_part1)
	} else {
		fmt.Println("Part 1 passed")
	}

	result_part2 := part2(example_input_str)
	if result_part2 != expected_part2 {
		fmt.Printf("Part 2 failed: got %d, expected %d, difference %d\n", result_part2, expected_part2, expected_part2-result_part2)
	} else {
		fmt.Println("Part 2 passed")
	}
}

func main() {
	// Run tests
	test()

	data := parse_input("input.txt")
	fmt.Println("Part 1 - fewest button presses:", part1(data))
	fmt.Println("Part 2 - joltage button presses:", part2(data))
}
