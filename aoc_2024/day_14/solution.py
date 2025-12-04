"""
Advent of Code 2024 - Day 14: Restroom Redoubt
https://adventofcode.com/2024/day/14
"""


def parse_input(filename: str) -> str:
    with open(filename, 'r') as f:
        data = f.read()
    
    return data


def part1(data: str, width: int = 101, height: int = 103) -> int:
    # TODO: Implement solution
    return 0


def part2(data: str) -> int:
    # TODO: Implement solution
    return 0


def test():
    """Test with example data from the puzzle."""
    example_data = """p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3"""

    # Part 1: Safety factor is 12 (for 11x7 grid)
    assert part1(example_data, width=11, height=7) == 12, "Part 1 failed"
    
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
