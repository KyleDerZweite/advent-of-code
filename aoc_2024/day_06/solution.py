"""
Advent of Code 2024 - Day 06: Guard Gallivant
https://adventofcode.com/2024/day/6
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
    example_data = """....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..."""

    # Part 1: Guard visits 41 distinct positions
    assert part1(example_data) == 41, "Part 1 failed"
    
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
