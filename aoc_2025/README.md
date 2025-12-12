# Advent of Code 2025 🎄

My solutions for [Advent of Code 2025](https://adventofcode.com/2025).

**Stars: ⭐ 24/24** — All puzzles completed!

## Progress

| Day | Puzzle | Part 1 | Part 2 | Solution |
|-----|--------|--------|--------|----------|
| 01 | [Secret Entrance](https://adventofcode.com/2025/day/1) | ⭐ | ⭐ | [solution.go](day_01/solution.go) |
| 02 | [Gift Shop](https://adventofcode.com/2025/day/2) | ⭐ | ⭐ | [solution.go](day_02/solution.go) |
| 03 | [Lobby](https://adventofcode.com/2025/day/3) | ⭐ | ⭐ | [solution.go](day_03/solution.go) |
| 04 | [Printing Department](https://adventofcode.com/2025/day/4) | ⭐ | ⭐ | [solution.go](day_04/solution.go) |
| 05 | [Cafeteria](https://adventofcode.com/2025/day/5) | ⭐ | ⭐ | [solution.go](day_05/solution.go) |
| 06 | [Trash Compactor](https://adventofcode.com/2025/day/6) | ⭐ | ⭐ | [solution.go](day_06/solution.go) |
| 07 | [Laboratories](https://adventofcode.com/2025/day/7) | ⭐ | ⭐ | [solution.go](day_07/solution.go) |
| 08 | [Playground](https://adventofcode.com/2025/day/8) | ⭐ | ⭐ | [solution.go](day_08/solution.go) |
| 09 | [Movie Theater](https://adventofcode.com/2025/day/9) | ⭐ | ⭐ | [solution.go](day_09/solution.go) |
| 10 | [Factory](https://adventofcode.com/2025/day/10) | ⭐ | ⭐ | [solution.go](day_10/solution.go) |
| 11 | [Reactor](https://adventofcode.com/2025/day/11) | ⭐ | ⭐ | [solution.go](day_11/solution.go) |
| 12 | [Christmas Tree Farm](https://adventofcode.com/2025/day/12) | ⭐ | ⭐ | [ai_solution.go](day_12/ai-solution/ai_solution.go) |

> **Note:** Day 11 required AI assistance for bug fixing, and Day 12 was solved entirely with AI help.
> I'm still learning Go and these puzzles were beyond my current skill level.
> The `ai-solution/` subfolder contains the AI-assisted solution with detailed documentation.

## ASCII Art
> I really liked the art on the website and tried to recreate it as best as i can as ascii art

```
. ..     ____ ''.    .  * ..  .  . .  '  <o '           
________/O___\__________|_________________O______   1 **
 . _______||_________ .  '         .'       .'          
.  | _@__ || _o_  '.|_ _________________________    2 **
.  |_&_%__||_oo__^=_[ \|..'  _    .. .. ..     |        
'      .' . .' '.  . \_]__--|_|___[]_[]_[]__//_|    3 **
. ..  '  .'  ..    ..  '   .   ____________//___        
__________________________  ...| \ '''''' // @@|    4 **
|_  ___ | .--.  ()   ()  |.' ..__[#]_@@__//_@@@|        
|_\_|^|_]_|==|_T_T_T_T_T_...'      ' . ..      ''   5 **
 || ' ____________    _______________________           
_||__/'...' '...' \_  |.     |~    .''.    .|  '.   6 **
|^@ |   1  2  3 () |  | '..'/ \'..'    '..' |____       
|&%;]__[]_[]_[]__<>|  |    |H_/|\   \\\\\\  | | |   7 **
. .  '' '.''  .   '...|<>__|H__|_\__|_____|_[_O_|       
 __________________________       ''    '' .  |     8 **
/'....'______'...'...'__'.|  _________________O__       
[&  @ |(_%) [ o  o  ^ ] \ |__|  [  ]  ''''''  | |   9 **
o=====|_____o=========o_|_[__]_____-/_-/_-/___|_|       
_________||______ ______________________________   10 **
|  ___       '..| |'..''..''..''..''..''..''..'|        
|_|...|_(:::::)_| |   *  ()  *  ()  *  ()  *   |   11 **
    |   '  .--.   |  <o>    <^o    <o>    o^>  |        
 '  '------'  '---#_<<^o>__<o^>>__<<^>o__<<^o>_|   12 **
```

## Running Solutions

Each day's solution is in its own folder (`day_XX/`). To run a solution:

```bash
cd day_01
go run main.go
```

Or from the repo root:

```bash
go run ./aoc_2025/day_01
```

## Structure

```
aoc_2025/
├── README.md
├── day_01/
│   ├── 01.md        # Puzzle description (git-ignored)
│   ├── input.txt    # Puzzle input (git-ignored)
│   └── solution.go   # Solution file
└── ...
```

## Legal Notice

Puzzle text and descriptions are © Advent of Code and are not included in this repository.
Links are provided to the original puzzles on [adventofcode.com](https://adventofcode.com/2025).
Input files are personal and git-ignored.

## Disclaimer

The docstrings and comments in the solution files were mostly generated with the assistance of AI (GitHub Copilot).
