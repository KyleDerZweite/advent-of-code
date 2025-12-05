# Advent of Code 2022 🎄

My solutions for [Advent of Code 2022](https://adventofcode.com/2022).

## Progress

| Day | Puzzle | Part 1 | Part 2 | Solution |
|-----|--------|--------|--------|----------|
| 01 | [Calorie Counting](https://adventofcode.com/2022/day/1) | | | [solution](day_01/) |
| 02 | [Rock Paper Scissors](https://adventofcode.com/2022/day/2) | | | [solution](day_02/) |
| 03 | [Rucksack Reorganization](https://adventofcode.com/2022/day/3) | | | [solution](day_03/) |
| 04 | [Camp Cleanup](https://adventofcode.com/2022/day/4) | | | [solution](day_04/) |
| 05 | [Supply Stacks](https://adventofcode.com/2022/day/5) | | | [solution](day_05/) |
| 06 | [Tuning Trouble](https://adventofcode.com/2022/day/6) | | | [solution](day_06/) |
| 07 | [No Space Left On Device](https://adventofcode.com/2022/day/7) | | | [solution](day_07/) |
| 08 | [Treetop Tree House](https://adventofcode.com/2022/day/8) | | | [solution](day_08/) |
| 09 | [Rope Bridge](https://adventofcode.com/2022/day/9) | | | [solution](day_09/) |
| 10 | [Cathode-Ray Tube](https://adventofcode.com/2022/day/10) | | | [solution](day_10/) |
| 11 | [Monkey in the Middle](https://adventofcode.com/2022/day/11) | | | [solution](day_11/) |
| 12 | [Hill Climbing Algorithm](https://adventofcode.com/2022/day/12) | | | [solution](day_12/) |
| 13 | [Distress Signal](https://adventofcode.com/2022/day/13) | | | [solution](day_13/) |
| 14 | [Regolith Reservoir](https://adventofcode.com/2022/day/14) | | | [solution](day_14/) |
| 15 | [Beacon Exclusion Zone](https://adventofcode.com/2022/day/15) | | | [solution](day_15/) |
| 16 | [Proboscidea Volcanium](https://adventofcode.com/2022/day/16) | | | [solution](day_16/) |
| 17 | [Pyroclastic Flow](https://adventofcode.com/2022/day/17) | | | [solution](day_17/) |
| 18 | [Boiling Boulders](https://adventofcode.com/2022/day/18) | | | [solution](day_18/) |
| 19 | [Not Enough Minerals](https://adventofcode.com/2022/day/19) | | | [solution](day_19/) |
| 20 | [Grove Positioning System](https://adventofcode.com/2022/day/20) | | | [solution](day_20/) |
| 21 | [Monkey Math](https://adventofcode.com/2022/day/21) | | | [solution](day_21/) |
| 22 | [Monkey Map](https://adventofcode.com/2022/day/22) | | | [solution](day_22/) |
| 23 | [Unstable Diffusion](https://adventofcode.com/2022/day/23) | | | [solution](day_23/) |
| 24 | [Blizzard Basin](https://adventofcode.com/2022/day/24) | | | [solution](day_24/) |
| 25 | [Full of Hot Air](https://adventofcode.com/2022/day/25) | | | [solution](day_25/) |

## Running Solutions

Each day's solution is in its own folder (`day_XX/`). To run a solution:

```bash
cd day_01
zig run solution.zig
# Or build and run:
zig build-exe solution.zig && ./solution
```

## Structure

```
aoc_2022/
├── README.md
├── day_01/
│   ├── 01.md         # Puzzle description (git-ignored)
│   ├── input.txt     # Puzzle input (git-ignored)
│   └── solution.zig  # Solution file
└── ...
```

## Solution Template

> **Note:** This template was automatically.

```zig
// Advent of Code 2022 - Day XX: Puzzle Name
// https://adventofcode.com/2022/day/X

const std = @import("std");

fn parseInput(allocator: std.mem.Allocator, filename: []const u8) ![]const u8 {
    const file = try std.fs.cwd().openFile(filename, .{});
    defer file.close();
    return try file.readToEndAlloc(allocator, std.math.maxInt(usize));
}

fn part1(data: []const u8) i64 {
    // TODO: Implement solution
    _ = data;
    return 0;
}

fn part2(data: []const u8) i64 {
    // TODO: Implement solution
    _ = data;
    return 0;
}

fn test_solutions() !void {
    const example_data = "";

    // Part 1: TODO - add expected value
    // try std.testing.expectEqual(@as(i64, x), part1(example_data));

    // Part 2: TODO - add expected value
    // try std.testing.expectEqual(@as(i64, x), part2(example_data));

    _ = example_data;
    std.debug.print("All tests passed!\n", .{});
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    try test_solutions();

    const data = try parseInput(allocator, "input.txt");
    defer allocator.free(data);

    const result1 = part1(data);
    std.debug.print("Part 1: {}\n", .{result1});

    const result2 = part2(data);
    std.debug.print("Part 2: {}\n", .{result2});
}
```

## Legal Notice

Puzzle text and descriptions are © Advent of Code and are not included in this repository.
Links are provided to the original puzzles on [adventofcode.com](https://adventofcode.com/2022).
Input files are personal and git-ignored.

## Disclaimer

The docstrings and comments in the solution files were mostly generated with the assistance of AI (GitHub Copilot).
