# Advent of Code 2021 🎄

My solutions for [Advent of Code 2021](https://adventofcode.com/2021).

## Progress

| Day | Puzzle | Part 1 | Part 2 | Solution |
|-----|--------|--------|--------|----------|
| 01 | [Sonar Sweep](https://adventofcode.com/2021/day/1) | | | [solution](day_01/) |
| 02 | [Dive!](https://adventofcode.com/2021/day/2) | | | [solution](day_02/) |
| 03 | [Binary Diagnostic](https://adventofcode.com/2021/day/3) | | | [solution](day_03/) |
| 04 | [Giant Squid](https://adventofcode.com/2021/day/4) | | | [solution](day_04/) |
| 05 | [Hydrothermal Venture](https://adventofcode.com/2021/day/5) | | | [solution](day_05/) |
| 06 | [Lanternfish](https://adventofcode.com/2021/day/6) | | | [solution](day_06/) |
| 07 | [The Treachery of Whales](https://adventofcode.com/2021/day/7) | | | [solution](day_07/) |
| 08 | [Seven Segment Search](https://adventofcode.com/2021/day/8) | | | [solution](day_08/) |
| 09 | [Smoke Basin](https://adventofcode.com/2021/day/9) | | | [solution](day_09/) |
| 10 | [Syntax Scoring](https://adventofcode.com/2021/day/10) | | | [solution](day_10/) |
| 11 | [Dumbo Octopus](https://adventofcode.com/2021/day/11) | | | [solution](day_11/) |
| 12 | [Passage Pathing](https://adventofcode.com/2021/day/12) | | | [solution](day_12/) |
| 13 | [Transparent Origami](https://adventofcode.com/2021/day/13) | | | [solution](day_13/) |
| 14 | [Extended Polymerization](https://adventofcode.com/2021/day/14) | | | [solution](day_14/) |
| 15 | [Chiton](https://adventofcode.com/2021/day/15) | | | [solution](day_15/) |
| 16 | [Packet Decoder](https://adventofcode.com/2021/day/16) | | | [solution](day_16/) |
| 17 | [Trick Shot](https://adventofcode.com/2021/day/17) | | | [solution](day_17/) |
| 18 | [Snailfish](https://adventofcode.com/2021/day/18) | | | [solution](day_18/) |
| 19 | [Beacon Scanner](https://adventofcode.com/2021/day/19) | | | [solution](day_19/) |
| 20 | [Trench Map](https://adventofcode.com/2021/day/20) | | | [solution](day_20/) |
| 21 | [Dirac Dice](https://adventofcode.com/2021/day/21) | | | [solution](day_21/) |
| 22 | [Reactor Reboot](https://adventofcode.com/2021/day/22) | | | [solution](day_22/) |
| 23 | [Amphipod](https://adventofcode.com/2021/day/23) | | | [solution](day_23/) |
| 24 | [Arithmetic Logic Unit](https://adventofcode.com/2021/day/24) | | | [solution](day_24/) |
| 25 | [Sea Cucumber](https://adventofcode.com/2021/day/25) | | | [solution](day_25/) |

## Running Solutions

Each day's solution is in its own folder (`day_XX/`). To run a solution:

```bash
cd day_01
elixir solution.exs
```

## Structure

```
aoc_2021/
├── README.md
├── day_01/
│   ├── 01.md         # Puzzle description (git-ignored)
│   ├── input.txt     # Puzzle input (git-ignored)
│   └── solution.exs  # Solution file
└── ...
```

## Solution Template

> **Note:** This template was automatically.

```elixir
# Advent of Code 2021 - Day XX: Puzzle Name
# https://adventofcode.com/2021/day/X

defmodule Solution do
  def parse_input(filename) do
    File.read!(filename) |> String.trim()
  end

  def part1(_data) do
    # TODO: Implement solution
    0
  end

  def part2(_data) do
    # TODO: Implement solution
    0
  end

  def test do
    example_data = ""

    # Part 1: TODO - add expected value
    # assert part1(example_data) == x, "Part 1 failed"
    
    # Part 2: TODO - add expected value
    # assert part2(example_data) == x, "Part 2 failed"

    IO.puts("All tests passed!")
  end

  def main do
    test()

    data = parse_input("input.txt")

    result1 = part1(data)
    IO.puts("Part 1: #{result1}")

    result2 = part2(data)
    IO.puts("Part 2: #{result2}")
  end
end

Solution.main()
```

## Legal Notice

Puzzle text and descriptions are © Advent of Code and are not included in this repository.
Links are provided to the original puzzles on [adventofcode.com](https://adventofcode.com/2021).
Input files are personal and git-ignored.

## Disclaimer

The docstrings and comments in the solution files were mostly generated with the assistance of AI (GitHub Copilot).
