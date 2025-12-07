"""
Advent of Code 2024 - Day 06: Guard Gallivant
https://adventofcode.com/2024/day/6
"""


def parse_input(filename: str) -> str:
    with open(filename, 'r') as f:
        data = f.read()
    
    return data


def part1(data: str) -> int:
    grid = data.strip().splitlines()
    rows, cols = len(grid), len(grid[0])

    # Locate starting position and facing direction.
    facing_map = {'^': 0, '>': 1, 'v': 2, '<': 3}
    start_r = start_c = direction = None
    for r, line in enumerate(grid):
        for c, ch in enumerate(line):
            if ch in facing_map:
                start_r, start_c, direction = r, c, facing_map[ch]
                break
        if start_r is not None:
            break

    # Direction order: up, right, down, left.
    deltas = [(-1, 0), (0, 1), (1, 0), (0, -1)]

    visited = set()
    r, c = start_r, start_c
    if r is None or c is None:
        raise ValueError("Starting position not found in the grid.")
    while 0 <= r < rows and 0 <= c < cols:
        visited.add((r, c))
        dr, dc = deltas[direction] # type: ignore
        nr, nc = r + dr, c + dc

        # Rotate right if the next step is blocked; otherwise move forward.
        if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == '#':
            direction = (direction + 1) % 4 # type: ignore
            continue

        r, c = nr, nc

    return len(visited)


def part2(data: str) -> int:
    grid = data.strip().splitlines()
    rows, cols = len(grid), len(grid[0])

    facing_map = {'^': 0, '>': 1, 'v': 2, '<': 3}
    start_r = start_c = direction = None
    for r, line in enumerate(grid):
        for c, ch in enumerate(line):
            if ch in facing_map:
                start_r, start_c, direction = r, c, facing_map[ch]
                break
        if start_r is not None:
            break

    deltas = [(-1, 0), (0, 1), (1, 0), (0, -1)]

    def causes_loop(block: tuple[int, int] | None) -> bool:
        seen_states = set()
        r, c, d = start_r, start_c, direction
        if r is None or c is None:
            raise ValueError("Starting position not found in the grid.")
        while 0 <= r < rows and 0 <= c < cols:
            state = (r, c, d)
            if state in seen_states:
                return True
            seen_states.add(state)

            dr, dc = deltas[d] # type: ignore
            nr, nc = r + dr, c + dc

            blocked = False
            if 0 <= nr < rows and 0 <= nc < cols:
                blocked = grid[nr][nc] == '#' or (block is not None and (nr, nc) == block)

            if blocked:
                d = (d + 1) % 4 # type: ignore
                continue

            r, c = nr, nc

        return False

    loop_count = 0
    for r in range(rows):
        for c in range(cols):
            if (r, c) == (start_r, start_c):
                continue
            if grid[r][c] == '#':
                continue
            if causes_loop((r, c)):
                loop_count += 1

    return loop_count


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
    
    assert part2(example_data) == 6, "Part 2 failed"
    
    print("All tests passed!")


if __name__ == "__main__":
    test()
    
    data = parse_input("input.txt")

    result1 = part1(data)
    print(f"Part 1: {result1}")
    
    result2 = part2(data)
    print(f"Part 2: {result2}")
