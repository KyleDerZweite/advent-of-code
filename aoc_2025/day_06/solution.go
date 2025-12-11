package main

import (
	"fmt"
	"os"
	"strconv"
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

func parse_string_to_matrix(input string) ([][]int, []string) {
	// Trim leading/trailing whitespace so we don't get empty first/last lines
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var operators []string
	columns := [][]int{}

	for i, ln := range lines {
		ln = strings.TrimSpace(ln)
		if ln == "" {
			continue
		}

		fields := strings.Fields(ln)
		// If this is the last line and any field is not a number, treat it as operators
		if i == len(lines)-1 {
			is_op := false
			for _, f := range fields {
				if _, err := strconv.Atoi(f); err != nil {
					is_op = true
					break
				}
			}
			if is_op {
				operators = fields
				continue
			}
		}

		// Ensure columns slice is pre-allocated based on first numeric line
		if len(columns) == 0 {
			columns = make([][]int, len(fields))
		}

		for j, f := range fields {
			if j >= len(columns) {
				extra := make([][]int, j-len(columns)+1)
				columns = append(columns, extra...)
			}
			v, _ := strconv.Atoi(f)

			columns[j] = append(columns[j], v)
		}

	}
	return columns, operators
}

func part1(input string) int {
	// Implement the logic for part 1 of the puzzle here
	grand_total := 0

	columns, operators := parse_string_to_matrix(input)

	for i, column := range columns {
		op := operators[i]
		column_total := 0
		for _, val := range column {
			// fmt.Printf("Operator '%s' - Value %d - Column Total %d - Grand Total %d\n", op, val, column_total, grand_total)
			switch op {
			case "*":
				if column_total == 0 {
					column_total = 1
				}
				column_total *= val
			case "+":
				column_total += val
			}
		}

		grand_total += column_total
	}

	return grand_total
}

func part2(input string) int {
	grand_total := 0

	string_matrix := [][]string{}
	lines := strings.Split(input, "\n")
	for _, ln := range lines {
		fields := strings.Split(ln, "")
		string_matrix = append(string_matrix, fields)
	}

	operator_index := len(string_matrix) - 1
	row_length := len(string_matrix[0])
	column_length := len(string_matrix)

	// fmt.Printf("Column Length: %d\n", column_length)
	// fmt.Printf("Row Length: %d\n", row_length)
	// fmt.Printf("Operator Index: %d\n", operator_index)

	skip_next := false
	values := []int{}
	for i := row_length - 1; i >= 0; i-- {
		if skip_next {
			skip_next = false
			continue
		}

		val_str := ""
		for j := 0; j < column_length-1; j++ {
			val_str += strings.TrimSpace(string_matrix[j][i])
		}
		val, _ := strconv.Atoi(val_str)
		values = append(values, val)

		op := string_matrix[operator_index][i]

		if op != " " {
			skip_next = true
			// fmt.Printf("Operator %s Values: %v\n", op, values)
			column_total := 0
			switch op {
			case "*":
				if column_total == 0 {
					column_total = 1
				}
				for _, val := range values {
					column_total *= val
				}
			case "+":
				for _, val := range values {
					column_total += val
				}
			}
			grand_total += column_total
			values = []int{}
		}

	}

	return grand_total
}

func test() {
	// Test with example data from the puzzle
	example_input_str := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	excpeted_part1 := 4277556
	excpeted_part2 := 3263827

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
	fmt.Println("Part 1 - Grand Total:", part1(data))
	fmt.Println("Part 2 - Cephalopod Math Grand Total:", part2(data))
}
