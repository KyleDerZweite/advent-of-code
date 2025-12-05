# Advent of Code 2016 🎄

My solutions for [Advent of Code 2016](https://adventofcode.com/2016).

## Progress

| Day | Puzzle | Part 1 | Part 2 | Solution |
|-----|--------|--------|--------|----------|
| 01 | [No Time for a Taxicab](https://adventofcode.com/2016/day/1) | | | [solution](day_01/) |
| 02 | [Bathroom Security](https://adventofcode.com/2016/day/2) | | | [solution](day_02/) |
| 03 | [Squares With Three Sides](https://adventofcode.com/2016/day/3) | | | [solution](day_03/) |
| 04 | [Security Through Obscurity](https://adventofcode.com/2016/day/4) | | | [solution](day_04/) |
| 05 | [How About a Nice Game of Chess?](https://adventofcode.com/2016/day/5) | | | [solution](day_05/) |
| 06 | [Signals and Noise](https://adventofcode.com/2016/day/6) | | | [solution](day_06/) |
| 07 | [Internet Protocol Version 7](https://adventofcode.com/2016/day/7) | | | [solution](day_07/) |
| 08 | [Two-Factor Authentication](https://adventofcode.com/2016/day/8) | | | [solution](day_08/) |
| 09 | [Explosives in Cyberspace](https://adventofcode.com/2016/day/9) | | | [solution](day_09/) |
| 10 | [Balance Bots](https://adventofcode.com/2016/day/10) | | | [solution](day_10/) |
| 11 | [Radioisotope Thermoelectric Generators](https://adventofcode.com/2016/day/11) | | | [solution](day_11/) |
| 12 | [Leonardo's Monorail](https://adventofcode.com/2016/day/12) | | | [solution](day_12/) |
| 13 | [A Maze of Twisty Little Cubicles](https://adventofcode.com/2016/day/13) | | | [solution](day_13/) |
| 14 | [One-Time Pad](https://adventofcode.com/2016/day/14) | | | [solution](day_14/) |
| 15 | [Timing is Everything](https://adventofcode.com/2016/day/15) | | | [solution](day_15/) |
| 16 | [Dragon Checksum](https://adventofcode.com/2016/day/16) | | | [solution](day_16/) |
| 17 | [Two Steps Forward](https://adventofcode.com/2016/day/17) | | | [solution](day_17/) |
| 18 | [Like a Rogue](https://adventofcode.com/2016/day/18) | | | [solution](day_18/) |
| 19 | [An Elephant Named Joseph](https://adventofcode.com/2016/day/19) | | | [solution](day_19/) |
| 20 | [Firewall Rules](https://adventofcode.com/2016/day/20) | | | [solution](day_20/) |
| 21 | [Scrambled Letters and Hash](https://adventofcode.com/2016/day/21) | | | [solution](day_21/) |
| 22 | [Grid Computing](https://adventofcode.com/2016/day/22) | | | [solution](day_22/) |
| 23 | [Safe Cracking](https://adventofcode.com/2016/day/23) | | | [solution](day_23/) |
| 24 | [Air Duct Spelunking](https://adventofcode.com/2016/day/24) | | | [solution](day_24/) |
| 25 | [Clock Signal](https://adventofcode.com/2016/day/25) | | | [solution](day_25/) |

## Running Solutions

Each day's solution is in its own folder (`day_XX/`). To run a solution:

```bash
cd day_01
kotlinc solution.kt -include-runtime -d solution.jar && java -jar solution.jar
# Or with Kotlin script:
kotlin solution.main.kts
```

## Structure

```
aoc_2016/
├── README.md
├── day_01/
│   ├── 01.md        # Puzzle description (git-ignored)
│   ├── input.txt    # Puzzle input (git-ignored)
│   └── solution.kt  # Solution file
└── ...
```

## Solution Template

> **Note:** This template was automatically.

```kotlin
/**
 * Advent of Code 2016 - Day XX: Puzzle Name
 * https://adventofcode.com/2016/day/X
 */

import java.io.File

fun parseInput(filename: String): String {
    return File(filename).readText().trim()
}

fun part1(data: String): Int {
    // TODO: Implement solution
    return 0
}

fun part2(data: String): Int {
    // TODO: Implement solution
    return 0
}

fun test() {
    val exampleData = """"""

    // Part 1: TODO - add expected value
    // check(part1(exampleData) == X) { "Part 1 failed" }
    
    // Part 2: TODO - add expected value
    // check(part2(exampleData) == X) { "Part 2 failed" }
    
    println("All tests passed!")
}

fun main() {
    test()
    
    val data = parseInput("input.txt")
    
    val result1 = part1(data)
    println("Part 1: $result1")
    
    val result2 = part2(data)
    println("Part 2: $result2")
}

main()
```

## Legal Notice

Puzzle text and descriptions are © Advent of Code and are not included in this repository.
Links are provided to the original puzzles on [adventofcode.com](https://adventofcode.com/2016).
Input files are personal and git-ignored.

## Disclaimer

The docstrings and comments in the solution files were mostly generated with the assistance of AI (GitHub Copilot).
