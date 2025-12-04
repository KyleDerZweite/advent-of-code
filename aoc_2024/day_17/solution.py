"""
Advent of Code 2024 - Day 17: Chronospatial Computer
https://adventofcode.com/2024/day/17
"""


def parse_input(filename: str) -> str:
    with open(filename, 'r') as f:
        data = f.read()
    
    return data


def part1(data: str) -> str:
    # TODO: Implement solution
    return ""


def part2(data: str) -> int:
    # TODO: Implement solution
    return 0


def test():
    """Test with example data from the puzzle."""
    example_data = """Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0"""

    # Part 1: Output is 4,6,3,5,6,3,5,2,1,0
    assert part1(example_data) == "4,6,3,5,6,3,5,2,1,0", "Part 1 failed"
    
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
