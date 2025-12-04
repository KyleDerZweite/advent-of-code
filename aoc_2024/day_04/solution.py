"""
Advent of Code 2024 - Day 04: Ceres Search
https://adventofcode.com/2024/day/4
"""


def parse_input(filename: str) -> str:
    with open(filename, 'r') as f:
        data = f.read()
    
    return data


def part1(data: str) -> int:
    # TODO: Implement solution
    search_pattern1 = "XMAS"
    search_pattern2 = "SAMX"
    occurences = 0
    # Checks for left & right overlaps
    for i in range(len(data) - 3):
        if data[i:i+4] == search_pattern1 or data[i:i+4] == search_pattern2:
            occurences += 1
    # Checks for up & down overlaps
    data_strings = data.split("\n")
    for n in range(len(data_strings)):
        for m in range(len(data_strings[n]) - 3):
            if data_strings[n][m:m+4] == search_pattern1 or data_strings[n][m:m+4] == search_pattern2:
                occurences += 1
    # Checks for diagonal overlaps
    rows = len(data_strings)
    cols = len(data_strings[0])
    for r in range(rows - 3):
        for c in range(cols - 3):
            diag1 = data_strings[r][c] + data_strings[r+1][c+1] + data_strings[r+2][c+2] + data_strings[r+3][c+3]
            diag2 = data_strings[r][c+3] + data_strings[r+1][c+2] + data_strings[r+2][c+1] + data_strings[r+3][c]
            if diag1 == search_pattern1 or diag1 == search_pattern2:
                occurences += 1
            if diag2 == search_pattern1 or diag2 == search_pattern2:
                occurences += 1

    print("Occurrences of XMAS or SAMX:", occurences)
    return occurences


def part2(data: str) -> int:
    # TODO: Implement solution
    return 0


def test():
    """Test with example data from the puzzle."""
    example_data = """MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX"""

    # Part 1: XMAS appears 18 times
    assert part1(example_data) == 18, "Part 1 failed"
    
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
