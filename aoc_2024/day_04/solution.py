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
    data_strings = data.split("\n")
    rows = len(data_strings)
    cols = len(data_strings[0]) if rows > 0 else 0
    
    # Checks for horizontal (left & right) overlaps
    for r in range(rows):
        for c in range(cols - 3):
            horizontal = data_strings[r][c:c+4]
            if horizontal == search_pattern1 or horizontal == search_pattern2:
                occurences += 1
    
    # Checks for vertical (up & down) overlaps
    for c in range(cols):
        for r in range(rows - 3):
            vertical = data_strings[r][c] + data_strings[r+1][c] + data_strings[r+2][c] + data_strings[r+3][c]
            if vertical == search_pattern1 or vertical == search_pattern2:
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
    return occurences


def part2(data: str) -> int:
    search_pattern1 = "MAS"
    search_pattern2 = "SAM"
    occurences = 0
    data_strings = data.split("\n")
    rows = len(data_strings)
    cols = len(data_strings[0]) if rows > 0 else 0

    # Check for X-MAS pattern in each 3x3 window
    # Both diagonals must spell MAS or SAM (independently)
    for r in range(rows - 2):
        for c in range(cols - 2):
            # Diagonal from top-left to bottom-right
            diag1 = data_strings[r][c] + data_strings[r+1][c+1] + data_strings[r+2][c+2]
            # Diagonal from top-right to bottom-left
            diag2 = data_strings[r][c+2] + data_strings[r+1][c+1] + data_strings[r+2][c]

            # Both diagonals must be MAS or SAM (each can be either)
            if (diag1 == search_pattern1 or diag1 == search_pattern2) and \
               (diag2 == search_pattern1 or diag2 == search_pattern2):
                occurences += 1

    return occurences


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
    
    # Part 2: X-MAS appears 9 times
    assert part2(example_data) == 9, "Part 2 failed"
    
    print("All tests passed!")


if __name__ == "__main__":
    test()
    
    data = parse_input("input.txt")

    result1 = part1(data)
    print(f"Part 1: {result1}")
    
    result2 = part2(data)
    print(f"Part 2: {result2}")
