"""
Advent of Code 2024 - Day 07: Bridge Repair
https://adventofcode.com/2024/day/7
"""


def parse_input(filename: str) -> str:
    with open(filename, 'r') as f:
        data = f.read()
    
    return data


def part1(data: str) -> int:
    # TODO: Implement solution
    return 0


def part2(data: str) -> int:
    # TODO: Implement solution
    return 0


def test():
    """Test with example data from the puzzle."""
    example_data = """190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20"""

    # Part 1: Total calibration result is 3749
    assert part1(example_data) == 3749, "Part 1 failed"
    
    # Part 2: TODO
    # assert part2(example_data) == X, "Part 2 failed"
    
    print("All tests passed!")


if __name__ == "__main__":
    test()
    
    data = parse_input("input.txt")

    result1 = part1(data)
    print(f"Part 1: {result1}")
    
    result2 = part2(data)
    print(f"Part 2: {result2}")
