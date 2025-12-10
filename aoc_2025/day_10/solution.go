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
}

func (m *machine) init_state() {
	m.current_state = make([]bool, len(m.indicator_lights))
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

func (m *machine) reset_state() {
	for i := range m.current_state {
		m.current_state[i] = false
	}
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
		buttonMatches := button_re.FindAllStringSubmatch(line, -1)
		for _, match := range buttonMatches {
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

func recursive_button_press(m *machine, presses []int, depth int, max_depth int, min_presses *int) {
	if depth > max_depth {
		return
	}
	if m.is_correct() {
		if *min_presses == -1 || depth < *min_presses {
			*min_presses += depth
		}
		return
	}

	for i := range m.buttons {
		m.press_button(i)
		recursive_button_press(m, append(presses, i), depth+1, max_depth, min_presses)
		m.reset_state()
		for _, p := range presses {
			m.press_button(p)
		}
	}
}

func part1(input string) int {
	machines := parse_machines(input)

	min_button_presses := 0
	combo_length := 0
	_ = combo_length

	for i := range machines {
		m := &machines[i]
		m.init_state()

		recursive_button_press(m, []int{}, 0, 10, &min_button_presses)
	}

	return min_button_presses
}

func part2(input string) int {
	machines := parse_machines(input)
	_ = machines

	joltage_sum := 0
	return joltage_sum
}

func test() {
	// Test with example data from the puzzle
	example_input_str := `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

	excpeted_part1 := 7
	excpeted_part2 := 0

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

	// data := parse_input("input.txt")
	// fmt.Println("Part 1 - fewest button presses:", part1(data))
	// fmt.Println("Part 2 - Max Joltage Sum (12):", part2(data))
}
