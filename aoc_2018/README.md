# Advent of Code 2018 🎄

My solutions for [Advent of Code 2018](https://adventofcode.com/2018).

## Progress

| Day | Puzzle | Part 1 | Part 2 | Solution |
|-----|--------|--------|--------|----------|
| 01 | [Chronal Calibration](https://adventofcode.com/2018/day/1) | | | [solution](day_01/) |
| 02 | [Inventory Management System](https://adventofcode.com/2018/day/2) | | | [solution](day_02/) |
| 03 | [No Matter How You Slice It](https://adventofcode.com/2018/day/3) | | | [solution](day_03/) |
| 04 | [Repose Record](https://adventofcode.com/2018/day/4) | | | [solution](day_04/) |
| 05 | [Alchemical Reduction](https://adventofcode.com/2018/day/5) | | | [solution](day_05/) |
| 06 | [Chronal Coordinates](https://adventofcode.com/2018/day/6) | | | [solution](day_06/) |
| 07 | [The Sum of Its Parts](https://adventofcode.com/2018/day/7) | | | [solution](day_07/) |
| 08 | [Memory Maneuver](https://adventofcode.com/2018/day/8) | | | [solution](day_08/) |
| 09 | [Marble Mania](https://adventofcode.com/2018/day/9) | | | [solution](day_09/) |
| 10 | [The Stars Align](https://adventofcode.com/2018/day/10) | | | [solution](day_10/) |
| 11 | [Chronal Charge](https://adventofcode.com/2018/day/11) | | | [solution](day_11/) |
| 12 | [Subterranean Sustainability](https://adventofcode.com/2018/day/12) | | | [solution](day_12/) |
| 13 | [Mine Cart Madness](https://adventofcode.com/2018/day/13) | | | [solution](day_13/) |
| 14 | [Chocolate Charts](https://adventofcode.com/2018/day/14) | | | [solution](day_14/) |
| 15 | [Beverage Bandits](https://adventofcode.com/2018/day/15) | | | [solution](day_15/) |
| 16 | [Chronal Classification](https://adventofcode.com/2018/day/16) | | | [solution](day_16/) |
| 17 | [Reservoir Research](https://adventofcode.com/2018/day/17) | | | [solution](day_17/) |
| 18 | [Settlers of The North Pole](https://adventofcode.com/2018/day/18) | | | [solution](day_18/) |
| 19 | [Go With The Flow](https://adventofcode.com/2018/day/19) | | | [solution](day_19/) |
| 20 | [A Regular Map](https://adventofcode.com/2018/day/20) | | | [solution](day_20/) |
| 21 | [Chronal Conversion](https://adventofcode.com/2018/day/21) | | | [solution](day_21/) |
| 22 | [Mode Maze](https://adventofcode.com/2018/day/22) | | | [solution](day_22/) |
| 23 | [Experimental Emergency Teleportation](https://adventofcode.com/2018/day/23) | | | [solution](day_23/) |
| 24 | [Immune System Simulator 20XX](https://adventofcode.com/2018/day/24) | | | [solution](day_24/) |
| 25 | [Four-Dimensional Adventure](https://adventofcode.com/2018/day/25) | | | [solution](day_25/) |

## Running Solutions

Each day's solution is in its own folder (`day_XX/`). To run a solution:

```bash
cd day_01
ocaml solution.ml
# Or compile and run:
ocamlfind ocamlopt -package str -linkpkg solution.ml -o solution && ./solution
```

## Structure

```
aoc_2018/
├── README.md
├── day_01/
│   ├── 01.md        # Puzzle description (git-ignored)
│   ├── input.txt    # Puzzle input (git-ignored)
│   └── solution.ml  # Solution file
└── ...
```

## Solution Template

> **Note:** This template was automatically.

```ocaml
(*
  Advent of Code 2018 - Day XX: Puzzle Name
  https://adventofcode.com/2018/day/X
*)

let parse_input filename =
  let ic = open_in filename in
  let n = in_channel_length ic in
  let s = really_input_string ic n in
  close_in ic;
  String.trim s

let part1 data =
  (* TODO: Implement solution *)
  0

let part2 data =
  (* TODO: Implement solution *)
  0

let test () =
  let example_data = "" in
  
  (* Part 1: TODO - add expected value *)
  (* assert (part1 example_data = X); *)
  
  (* Part 2: TODO - add expected value *)
  (* assert (part2 example_data = X); *)
  
  print_endline "All tests passed!"

let () =
  test ();
  
  let data = parse_input "input.txt" in
  
  let result1 = part1 data in
  Printf.printf "Part 1: %d\n" result1;
  
  let result2 = part2 data in
  Printf.printf "Part 2: %d\n" result2
```

## Legal Notice

Puzzle text and descriptions are © Advent of Code and are not included in this repository.
Links are provided to the original puzzles on [adventofcode.com](https://adventofcode.com/2018).
Input files are personal and git-ignored.

## Disclaimer

The docstrings and comments in the solution files were mostly generated with the assistance of AI (GitHub Copilot).
