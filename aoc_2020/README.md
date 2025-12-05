# Advent of Code 2020 🎄

My solutions for [Advent of Code 2020](https://adventofcode.com/2020).

## Progress

| Day | Puzzle | Part 1 | Part 2 | Solution |
|-----|--------|--------|--------|----------|
| 01 | [Report Repair](https://adventofcode.com/2020/day/1) | | | [solution](day_01/) |
| 02 | [Password Philosophy](https://adventofcode.com/2020/day/2) | | | [solution](day_02/) |
| 03 | [Toboggan Trajectory](https://adventofcode.com/2020/day/3) | | | [solution](day_03/) |
| 04 | [Passport Processing](https://adventofcode.com/2020/day/4) | | | [solution](day_04/) |
| 05 | [Binary Boarding](https://adventofcode.com/2020/day/5) | | | [solution](day_05/) |
| 06 | [Custom Customs](https://adventofcode.com/2020/day/6) | | | [solution](day_06/) |
| 07 | [Handy Haversacks](https://adventofcode.com/2020/day/7) | | | [solution](day_07/) |
| 08 | [Handheld Halting](https://adventofcode.com/2020/day/8) | | | [solution](day_08/) |
| 09 | [Encoding Error](https://adventofcode.com/2020/day/9) | | | [solution](day_09/) |
| 10 | [Adapter Array](https://adventofcode.com/2020/day/10) | | | [solution](day_10/) |
| 11 | [Seating System](https://adventofcode.com/2020/day/11) | | | [solution](day_11/) |
| 12 | [Rain Risk](https://adventofcode.com/2020/day/12) | | | [solution](day_12/) |
| 13 | [Shuttle Search](https://adventofcode.com/2020/day/13) | | | [solution](day_13/) |
| 14 | [Docking Data](https://adventofcode.com/2020/day/14) | | | [solution](day_14/) |
| 15 | [Rambunctious Recitation](https://adventofcode.com/2020/day/15) | | | [solution](day_15/) |
| 16 | [Ticket Translation](https://adventofcode.com/2020/day/16) | | | [solution](day_16/) |
| 17 | [Conway Cubes](https://adventofcode.com/2020/day/17) | | | [solution](day_17/) |
| 18 | [Operation Order](https://adventofcode.com/2020/day/18) | | | [solution](day_18/) |
| 19 | [Monster Messages](https://adventofcode.com/2020/day/19) | | | [solution](day_19/) |
| 20 | [Jurassic Jigsaw](https://adventofcode.com/2020/day/20) | | | [solution](day_20/) |
| 21 | [Allergen Assessment](https://adventofcode.com/2020/day/21) | | | [solution](day_21/) |
| 22 | [Crab Combat](https://adventofcode.com/2020/day/22) | | | [solution](day_22/) |
| 23 | [Crab Cups](https://adventofcode.com/2020/day/23) | | | [solution](day_23/) |
| 24 | [Lobby Layout](https://adventofcode.com/2020/day/24) | | | [solution](day_24/) |
| 25 | [Combo Breaker](https://adventofcode.com/2020/day/25) | | | [solution](day_25/) |

## Running Solutions

Each day's solution is in its own folder (`day_XX/`). To run a solution:

```bash
cd day_01
runhaskell solution.hs
# Or compile and run:
ghc -o solution solution.hs && ./solution
```

## Structure

```
aoc_2020/
├── README.md
├── day_01/
│   ├── 01.md        # Puzzle description (git-ignored)
│   ├── input.txt    # Puzzle input (git-ignored)
│   └── solution.hs  # Solution file
└── ...
```

## Solution Template

> **Note:** This template was automatically.

```haskell
-- Advent of Code 2020 - Day XX: Puzzle Name
-- https://adventofcode.com/2020/day/X

import System.IO

parseInput :: String -> String
parseInput = id

part1 :: String -> Int
part1 _ = 0  -- TODO: Implement solution

part2 :: String -> Int
part2 _ = 0  -- TODO: Implement solution

test :: IO ()
test = do
    let exampleData = ""
    
    -- Part 1: TODO - add expected value
    -- assert (part1 exampleData == x) "Part 1 failed"
    
    -- Part 2: TODO - add expected value
    -- assert (part2 exampleData == x) "Part 2 failed"
    
    putStrLn "All tests passed!"

main :: IO ()
main = do
    test
    
    contents <- readFile "input.txt"
    let input = parseInput contents
    
    putStrLn $ "Part 1: " ++ show (part1 input)
    putStrLn $ "Part 2: " ++ show (part2 input)
```

## Legal Notice

Puzzle text and descriptions are © Advent of Code and are not included in this repository.
Links are provided to the original puzzles on [adventofcode.com](https://adventofcode.com/2020).
Input files are personal and git-ignored.

## Disclaimer

The docstrings and comments in the solution files were mostly generated with the assistance of AI (GitHub Copilot).
