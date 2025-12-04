"""
Advent of Code 2024 - Day 18: RAM Run
https://adventofcode.com/2024/day/18
"""


def parse_input(filename: str) -> str:
    with open(filename, 'r') as f:
        data = f.read()
    
    return data


def part1(data: str, grid_size: int = 71, num_bytes: int = 1024) -> int:
    # TODO: Implement solution
    return 0


def part2(data: str) -> str:
    # TODO: Implement solution
    return ""


def test():
    """Test with example data from the puzzle."""
    example_data = """5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0"""

    # Part 1: Minimum steps is 22 (for 7x7 grid with first 12 bytes)
    assert part1(example_data, grid_size=7, num_bytes=12) == 22, "Part 1 failed"
    
    # Part 2: TODO
    # assert part2(example_data) == "X,Y", "Part 2 failed"
    
    print("All tests passed!")


if __name__ == "__main__":
    test()
    
    data = parse_input("input.txt")

    result1 = part1(data)
    print(f"Part 1: {result1}")
    
    result2 = part2(data)
    print(f"Part 2: {result2}")
