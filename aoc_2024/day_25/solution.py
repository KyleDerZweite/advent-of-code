"""
Advent of Code 2024 - Day 25: Code Chronicle
https://adventofcode.com/2024/day/25
"""


def parse_input(filename: str) -> str:
    with open(filename, 'r') as f:
        data = f.read()
    
    return data


def part1(data: str) -> int:
    # TODO: Implement solution
    return 0


def part2(data: str) -> int:
    # Day 25 typically only has one part
    return 0


def test():
    """Test with example data from the puzzle."""
    example_data = """#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####"""

    # Part 1: 3 unique lock/key pairs fit together
    assert part1(example_data) == 3, "Part 1 failed"
    
    print("All tests passed!")


if __name__ == "__main__":
    test()
    
    data = parse_input("input.txt")

    result1 = part1(data)
    print(f"Part 1: {result1}")
