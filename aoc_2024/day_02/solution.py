"""
Advent of Code 2024 - Day 02: Red-Nosed Reports
https://adventofcode.com/2024/day/2
"""

from collections import Counter


def parse_input(filename: str) -> str:
    with open(filename, 'r') as f:
        data = f.read()
    
    return data

def part1(data: str) -> int:
    min_diff = 1
    max_diff = 3
    safe_report_count = 0
    lines = data.strip().split('\n')

    for line in lines:
        last_digit = None
        last_increasing = None
        is_safe = True
        digits = line.split(' ') 
        for digit in digits:
            # print(f"Processing digit: {digit} in line: {line} with last_digit: {last_digit} and last_increasing: {last_increasing}")
            if last_digit is not None:
                diff = abs(int(digit) - int(last_digit))
                # print(f"  Difference: {diff} between {digit} and {last_digit}")
                if diff < min_diff or diff > max_diff:
                    is_safe = False
                    break
                current_increasing = int(digit) > int(last_digit)
                if last_increasing is not None:
                    if last_increasing != current_increasing:
                        is_safe = False
                        break
                last_increasing = current_increasing
            last_digit = digit
        if is_safe:
            safe_report_count += 1
        # print(f"Line: {line} is {'safe' if is_safe else 'not safe'} - Total safe reports so far: {safe_report_count}")
    return safe_report_count

def is_report_safe(digits: list) -> bool:
    min_diff = 1
    max_diff = 3
    last_digit = None
    last_increasing = None
    for digit in digits:
        if last_digit is not None:
            diff = abs(int(digit) - int(last_digit))
            if diff < min_diff or diff > max_diff:
                return False
            current_increasing = int(digit) > int(last_digit)
            if last_increasing is not None:
                if last_increasing != current_increasing:
                    return False
            last_increasing = current_increasing
        last_digit = digit
    return True

def part2(data: str) -> int:
    safe_report_count = 0
    lines = data.strip().split('\n')

    for line in lines:
        digits = line.split(' ')
        # Check if already safe without removing any digit
        if is_report_safe(digits):
            safe_report_count += 1
            continue
        # Try removing each digit one at a time
        for i in range(len(digits)):
            digits_without_i = digits[:i] + digits[i+1:]
            if is_report_safe(digits_without_i):
                safe_report_count += 1
                break
    return safe_report_count


def test():
    """Test with example data from the puzzle."""
    example_data = """7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"""

    # Part 1: 2 Safe Reports
    assert part1(example_data) == 2, "Part 1 failed"
    
    # Part 2: 4 Safe Reports
    assert part2(example_data) == 4, "Part 2 failed"
    
    print("All tests passed!")


if __name__ == "__main__":
    test()
    
    data = parse_input("input.txt")

    result1 = part1(data)
    print(f"Part 1 - Safe Reports: {result1}")
    result2 = part2(data)
    print(f"Part 2 - Safe Reports: {result2}")