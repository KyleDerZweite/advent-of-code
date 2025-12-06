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

def part2_get_incorrect_lists(data: str):
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
    return rules, incorrect_lists

def part2_bogo(data: str) -> int:
    """
    This was my first thought that i directly had on part2, but i also implemented a more efficient one, still wanted to test this.
    It somehow works... At least for the example input... But yeah can't recommend for full input :D
    Still found it amusing.
    """
    rules, incorrect_lists = part2_get_incorrect_lists(data)
    middle_sum = 0  
  
    # For each incorrect list, try to fix it by random shuffling until correct
    import random
    failsafe = 10000000  # Prevent infinite loops
    for values in incorrect_lists:
        attempts = 0
        while True:
            random.shuffle(values)
            attempts += 1
            print(f"Attempt {attempts}: {','.join(map(str, values))}")
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
                break
            if attempts > failsafe:
                raise RuntimeError("Failsafe triggered: too many attempts to fix list")
        middle_index = len(values) // 2
        middle_sum += values[middle_index]
    return middle_sum

def part2_optimized_bogo(data: str) -> int:
    """
    Full-list bogo with faster validation: convert rules once, pre-filter relevant rules per list,
    and avoid repeated ``index`` calls. Still intentionally shuffles the whole list.
    """
    rules, incorrect_lists = part2_get_incorrect_lists(data)
    rule_pairs = [(int(a), int(b)) for a, b in rules]
    middle_sum = 0

    import random
    failsafe = 1000000  # Prevent infinite loops
    for original_values in incorrect_lists:
        values = list(original_values)
        relevant_rules = [rule for rule in rule_pairs if rule[0] in values and rule[1] in values]
        attempts = 0

        while True:
            random.shuffle(values)
            attempts += 1
            print(f"Attempt {attempts}: {','.join(map(str, values))}")

            positions = {value: idx for idx, value in enumerate(values)}
            is_valid = True
            for first, second in relevant_rules:
                if positions[first] > positions[second]:
                    is_valid = False
                    break

            if is_valid:
                break
            if attempts > failsafe:
                raise RuntimeError("Failsafe triggered: too many attempts to fix list")

        middle_index = len(values) // 2
        middle_sum += values[middle_index]
    return middle_sum

def part2_swapp_adjacent(data: str) -> int:
    middle_sum = 0
    rules, incorrect_lists = part2_get_incorrect_lists(data)

    # For each incorrect list, try to fix it by swapping adjacent elements
    for values in incorrect_lists:
        changed = True
        while changed:
            changed = False
            for rule in rules:
                first, second = int(rule[0]), int(rule[1])
                if first in values and second in values:
                    pos1 = values.index(first)
                    pos2 = values.index(second)
                    if pos1 > pos2:
                        values[pos1], values[pos2] = values[pos2], values[pos1]
                        changed = True
        middle_index = len(values) // 2
        middle_sum += values[middle_index]
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
    print("Running Part 2 tests (may take some time due to bogo methods)...")
    assert part2_bogo(example_data) == 123, "Part (BOGO) 2 failed"
    
    print("Running Part 2 OPTIMIZED BOGO test...")
    assert part2_optimized_bogo(example_data) == 123, "Part (OPTIMIZED BOGO) 2 failed"

    print("Running Part 2 SWAPP ADJACENT test...")
    assert part2_swapp_adjacent(example_data) == 123, "Part (SWAPP ADJACENT) 2 failed"
    
    print("All tests passed!")


if __name__ == "__main__":
    test()
    
    data = parse_input("input.txt")

    result1 = part1(data)
    print(f"Part 1: {result1}")
    
    result2 = part2_swapp_adjacent(data)
    print(f"Part 2: {result2}")
