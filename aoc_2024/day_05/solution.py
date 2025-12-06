"""
Advent of Code 2024 - Day 05: Print Queue
https://adventofcode.com/2024/day/5
"""


def parse_input(filename: str) -> str:
    with open(filename, 'r') as f:
        data = f.read()
    
    return data

def parse_input_str(data: str):
    sections = data.strip().split("\n\n")
    rules = [line.split('|') for line in sections[0].strip().splitlines()]
    value_lists = [list(map(int, line.split(','))) for line in sections[1].strip().splitlines()]
    return rules, value_lists


def part1(data: str):
    middle_sum = 0
    rules, value_lists = parse_input_str(data)

    for values in value_lists:
        is_valid = True
        # Check if values are in correct order according to rules
        for rule in rules:
            if not ((rule[0] in map(str, values)) and (rule[1] in map(str, values))):
                continue
            pos1 = values.index(int(rule[0]))
            pos2 = values.index(int(rule[1]))
            if pos1 > pos2:
                is_valid = False
                break
        if is_valid:
            middle_index = len(values) // 2
            middle_sum += values[middle_index]
    return middle_sum

def part2(data: str) -> int:
    middle_sum = 0
    rules, value_lists = parse_input_str(data)

    incorrect_lists = []

    for values in value_lists:
        # Check if values are in correct order according to rules
        for rule in rules:
            if not ((rule[0] in map(str, values)) and (rule[1] in map(str, values))):
                continue
            pos1 = values.index(int(rule[0]))
            pos2 = values.index(int(rule[1]))
            if pos1 > pos2:
                incorrect_lists.append(values)
                break

    # For each incorrect list, try to fix it by swapping adjacent elements
    for values in incorrect_lists:
        n = len(values)
        fixed = False
        for i in range(n - 1):
            # Swap adjacent elements
            values[i], values[i + 1] = values[i + 1], values[i]
            # Check if now valid
            is_valid = True
            for rule in rules:
                if not ((rule[0] in map(str, values)) and (rule[1] in map(str, values))):
                    continue
                pos1 = values.index(int(rule[0]))
                pos2 = values.index(int(rule[1]))
                if pos1 > pos2:
                    is_valid = False
                    break
            if is_valid:
                middle_index = len(values) // 2
                middle_sum += values[middle_index]
                fixed = True
                break
            # Swap back if not valid
            values[i], values[i + 1] = values[i + 1], values[i]
        if not fixed:
            raise ValueError("Could not fix the list:", values)
    

    return middle_sum


def test():
    """Test with example data from the puzzle."""
    example_data = """47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47"""

    # Part 1: Sum of middle page numbers from correctly-ordered updates is 143
    assert part1(example_data) == 143, "Part 1 failed"
    
    # Part 2: TODO
    assert part2(example_data) == 123, "Part 2 failed"
    
    print("All tests passed!")


if __name__ == "__main__":
    test()
    
    data = parse_input("input.txt")

    result1 = part1(data)
    print(f"Part 1: {result1}")
    
    result2 = part2(data)
    print(f"Part 2: {result2}")
