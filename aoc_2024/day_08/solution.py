"""Advent of Code 2024 - Day 08: Resonant Collinearity."""

from collections import defaultdict
from math import gcd


def parse_input(filename: str) -> str:
    with open(filename, "r", encoding="utf-8") as f:
        return f.read()


def is_letter_digit(ch: str) -> bool:
    return ("A" <= ch <= "Z") or ("a" <= ch <= "z") or ("0" <= ch <= "9")


def part1(data: str) -> int:
    lines = [line for line in data.splitlines() if line != ""]
    if not lines:
        return 0

    h = len(lines)
    antinodes: set[tuple[int, int]] = set()

    positions: dict[str, list[tuple[int, int]]] = defaultdict(list)
    for y, line in enumerate(lines):
        for x, ch in enumerate(line):
            if is_letter_digit(ch):
                positions[ch].append((x, y))

    for pts in positions.values():
        n = len(pts)
        if n < 2:
            continue
        for i in range(n):
            ax, ay = pts[i]
            for j in range(i + 1, n):
                bx, by = pts[j]

                x1, y1 = 2 * bx - ax, 2 * by - ay
                if 0 <= y1 < h and 0 <= x1 < len(lines[y1]):
                    antinodes.add((x1, y1))

                x2, y2 = 2 * ax - bx, 2 * ay - by
                if 0 <= y2 < h and 0 <= x2 < len(lines[y2]):
                    antinodes.add((x2, y2))

    return len(antinodes)


def part2(data: str) -> int:
    lines = [line for line in data.splitlines() if line != ""]
    if not lines:
        return 0

    h = len(lines)
    antinodes: set[tuple[int, int]] = set()

    positions: dict[str, list[tuple[int, int]]] = defaultdict(list)
    for y, line in enumerate(lines):
        for x, ch in enumerate(line):
            if is_letter_digit(ch):
                positions[ch].append((x, y))

    for pts in positions.values():
        n = len(pts)
        if n < 2:
            continue
        for i in range(n):
            ax, ay = pts[i]
            for j in range(i + 1, n):
                bx, by = pts[j]
                dx, dy = bx - ax, by - ay
                g = gcd(dx, dy)
                step_x, step_y = dx // g, dy // g

                # walk backward to the first in-bounds point on the line
                start_x, start_y = ax, ay
                while True:
                    nx, ny = start_x - step_x, start_y - step_y
                    if 0 <= ny < h and 0 <= nx < len(lines[ny]):
                        start_x, start_y = nx, ny
                        continue
                    break

                # sweep forward until out of bounds, marking all antinodes
                x, y = start_x, start_y
                while 0 <= y < h and 0 <= x < len(lines[y]):
                    antinodes.add((x, y))
                    x += step_x
                    y += step_y

    return len(antinodes)


def test():
    """Test with example data from the puzzle."""

    example_data = """............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............"""

    assert part1(example_data) == 14, "Part 1 failed"
    assert part2(example_data) == 34, "Part 2 failed"
    print("All tests passed!")


if __name__ == "__main__":
    test()

    data = parse_input("input.txt")

    result1 = part1(data)
    print(f"Part 1: {result1}")

    result2 = part2(data)
    print(f"Part 2: {result2}")
