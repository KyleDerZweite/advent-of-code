"""
Advent of Code 2024 - Day 03: Mull It Over
https://adventofcode.com/2024/day/3
"""

import re


def parse_input(filename: str) -> str:
    with open(filename, 'r') as f:
        data = f.read()
    
    return data


def part1(data: str) -> int:
    total = 0
    # Find all valid mul(X,Y) patterns where X and Y are 1-3 digit numbers
    pattern = r'mul\((\d{1,3}),(\d{1,3})\)'
    matches = re.findall(pattern, data)
    for match in matches:
        x = int(match[0])
        y = int(match[1])
        total += x * y
    return total


def part2(data: str) -> int:
    total = 0
    enabled = True
    # Find all mul(X,Y), do(), and don't() instructions in order
    pattern = r"mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)"
    matches = re.finditer(pattern, data)
    for match in matches:
        instruction = match.group()
        if instruction == "do()":
            enabled = True
        elif instruction == "don't()":
            enabled = False
        elif enabled:
            x = int(match.group(1))
            y = int(match.group(2))
            total += x * y
    return total


def test():
    """Test with example data from the puzzle."""
    example_data = """xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"""

    # Part 1: 161 (2*4 + 5*5 + 11*8 + 8*5)
    assert part1(example_data) == 161, "Part 1 failed"
        
    # Part 2: 48 (2*4 + 8*5)
    assert part2(example_data) == 48, "Part 2 failed"
    
    print("All tests passed!")


if __name__ == "__main__":
    test()
    
    data = parse_input("input.txt")

    result1 = part1(data)
    print(f"Part 1 - Sum of multiplications: {result1}")
    
    result2 = part2(data)
    print(f"Part 2 - Sum of enabled multiplications: {result2}")
