# 🎄 Advent of Code

My solutions for [Advent of Code](https://adventofcode.com/) - an annual coding challenge with daily puzzles from December 1st to 25th.

## Years

| Year | Language | Stars | Status |
|------|----------|-------|--------|
| [2025](aoc_2025/) | Go | ⭐ 24/24 | ✅ Complete |
| [2024](aoc_2024/) | Python | ⭐ 16/50 | 🔄 In Progress |
| [2023](aoc_2023/) | Rust | ☆ 0/50 | 🔄 Not Started |
| [2022](aoc_2022/) | Zig | ☆ 0/50 | 🆕 Not Started |
| [2021](aoc_2021/) | Elixir | ☆ 0/50 | 🆕 Not Started |
| [2020](aoc_2020/) | Haskell | ☆ 0/50 | 🆕 Not Started |
| [2019](aoc_2019/) | Dart | ☆ 0/50 | 🆕 Not Started |
| [2018](aoc_2018/) | OCaml | ☆ 0/50 | 🆕 Not Started |
| [2017](aoc_2017/) | Gleam | ☆ 0/50 | 🆕 Not Started |
| [2016](aoc_2016/) | Kotlin | ☆ 0/50 | 🆕 Not Started |
| [2015](aoc_2015/) | Julia | ☆ 0/50 | 🆕 Not Started |

> **Note on 2025:** Day 11 required AI assistance for bug fixing, and Day 12 was solved with AI help.
> Still learning Go — see the [aoc_2025 README](aoc_2025/README.md) for details.

## Repository Structure

```
advent-of-code/
├── aoc_2015/          # Julia solutions
├── aoc_2016/          # Kotlin solutions
├── aoc_2017/          # Gleam solutions
├── aoc_2018/          # OCaml solutions
├── aoc_2019/          # Dart solutions
├── aoc_2020/          # Haskell solutions
├── aoc_2021/          # Elixir solutions
├── aoc_2022/          # Zig solutions
├── aoc_2023/          # Rust solutions
├── aoc_2024/          # Python solutions
├── aoc_2025/          # Go solutions
└── README.md
```

Each year folder contains its own README with detailed progress and instructions.

## About Advent of Code

> *Advent of Code is an Advent calendar of small programming puzzles for a variety of skill sets and skill levels that can be solved in any programming language you like.*
> — [Eric Wastl](https://twitter.com/ericwastl), creator of AoC

## Tooling

This repository includes [`fetch_aoc.py`](fetch_aoc.py), a helper script to fetch puzzle descriptions for personal use.

Use [uv](https://docs.astral.sh/uv/getting-started/installation/) 0.11.16 for
the helper. The project selects Python 3.14 through `.python-version`; uv downloads
it if needed. From the repository root, install the locked dependencies and
preview a fetch without changing files or making requests:

```bash
uv sync --locked
uv run --locked fetch_aoc.py --help
uv run --locked fetch_aoc.py --years 2025 --days 1 --dry-run
```

Remove `--dry-run` to fetch descriptions. Existing descriptions are skipped unless
you pass `--force`. The helper's dependencies live in `pyproject.toml` and
`uv.lock`, separate from the language-specific solution projects. Update a
dependency with `uv lock --upgrade-package requests`, then verify with
`uv sync --locked` and the commands above. Dependabot checks the uv project weekly.

> ⚠️ **Important**: This tool is for **personal use only** to help organize your own solutions.
> Please be respectful of [adventofcode.com](https://adventofcode.com) servers.
> Puzzle content is © Advent of Code / Eric Wastl and is git-ignored.

## Legal

- Puzzle descriptions are © [Advent of Code](https://adventofcode.com/about) / Eric Wastl and are not included in this repository
- Solutions are my own work (with AI-assisted documentation)
- Input files are personalized and git-ignored
- Please consider [supporting Advent of Code](https://adventofcode.com/support) if you enjoy it!

## License

Solutions are available under the MIT License. See [LICENSE](LICENSE).
