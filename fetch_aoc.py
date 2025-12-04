#!/usr/bin/env python3
"""
fetch_aoc.py - Fetch Advent of Code puzzle descriptions.

Creates directory structures and downloads puzzle descriptions (Part 1 only)
for specified years and days from adventofcode.com.

Usage:
    python fetch_aoc.py --years 2015-2025
    python fetch_aoc.py --years 2024,2025 --days 1-10
    python fetch_aoc.py --years 2023 --force --delay 3

⚠️  IMPORTANT DISCLAIMER:
    This script is intended for PERSONAL USE ONLY to help organize your own
    Advent of Code solutions. Please be respectful of adventofcode.com servers.
    
    Advent of Code is a labor of love by Eric Wastl. If you enjoy it, please:
    - Don't hammer the servers with excessive requests
    - Consider supporting AoC: https://adventofcode.com/support
    - Read the about page: https://adventofcode.com/about
    
    Puzzle content is © Advent of Code / Eric Wastl and is NOT included in
    version control. This script only fetches publicly available Part 1
    descriptions for personal reference.
"""

import argparse
import os
import re
import sys
import time
from datetime import datetime
from pathlib import Path

try:
    import requests
    from bs4 import BeautifulSoup
    import html2text
except ImportError as e:
    print(f"Missing dependency: {e}")
    print("Install with: pip install requests beautifulsoup4 html2text")
    sys.exit(1)


# Configuration
AOC_BASE_URL = "https://adventofcode.com"
AOC_ABOUT_URL = "https://adventofcode.com/about"
AOC_SUPPORT_URL = "https://adventofcode.com/support"
USER_AGENT = "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/109.0"
FIRST_AOC_YEAR = 2015
MAX_RETRIES = 3
BACKOFF_FACTOR = 2
CONFIRM_THRESHOLD_DAYS = 5  # Require confirmation when fetching more than this many days


def show_disclaimer_and_confirm(total_fetches: int, years: list[int], days: list[int]) -> bool:
    """
    Show disclaimer and request user confirmation for bulk fetches.
    Returns True if user confirms, False otherwise.
    """
    disclaimer = f"""
╔══════════════════════════════════════════════════════════════════════════════╗
║                           ⚠️  IMPORTANT NOTICE ⚠️                              ║
╠══════════════════════════════════════════════════════════════════════════════╣
║  Advent of Code is created and maintained by Eric Wastl as a labor of love.  ║
║  Please be respectful of the servers and don't abuse this tool.              ║
║                                                                              ║
║  • This script is for PERSONAL USE to organize your own AoC solutions        ║
║  • Puzzle content is © Advent of Code and should not be redistributed        ║
║  • Consider supporting AoC: {AOC_SUPPORT_URL:<43}      ║
║  • Learn more: {AOC_ABOUT_URL:<55}       ║
╚══════════════════════════════════════════════════════════════════════════════╝

📊 You are about to fetch up to {total_fetches} puzzle descriptions.
   Years: {min(years)}-{max(years)} ({len(years)} years)
   Days:  {min(days)}-{max(days)} ({len(days)} days per year)
"""
    print(disclaimer)
    
    try:
        response = input("Do you accept these terms and wish to continue? [y/N]: ").strip().lower()
        return response in ('y', 'yes')
    except (KeyboardInterrupt, EOFError):
        print("\nAborted.")
        return False


def parse_range(range_str: str) -> list[int]:
    """Parse a range string like '2015-2025' or '2024,2025' into a list of integers."""
    result = []
    for part in range_str.split(","):
        part = part.strip()
        if "-" in part:
            start, end = part.split("-", 1)
            result.extend(range(int(start), int(end) + 1))
        else:
            result.append(int(part))
    return sorted(set(result))


def get_available_days(year: int) -> int:
    """Get the number of available days for a given year."""
    now = datetime.now()
    if year < now.year:
        return 25
    elif year == now.year and now.month == 12:
        return min(now.day, 25)
    elif year == now.year and now.month > 12:
        return 25
    else:
        return 0


def fetch_puzzle(year: int, day: int, delay: float) -> tuple[str, str] | None:
    """
    Fetch puzzle description from AoC website.
    Returns (title, content) or None on failure.
    """
    url = f"{AOC_BASE_URL}/{year}/day/{day}"
    headers = {"User-Agent": USER_AGENT}

    if delay < 1.0:
        delay = 1.0

    for attempt in range(MAX_RETRIES):
        try:
            time.sleep(delay)
            response = requests.get(url, headers=headers, timeout=30)

            if response.status_code == 404:
                print(f"  ✗ Day {day} not found (404)")
                return None
            elif response.status_code == 429:
                wait_time = BACKOFF_FACTOR ** (attempt + 2)
                print(f"  ⚠ Rate limited, waiting {wait_time}s...")
                time.sleep(wait_time)
                continue
            elif response.status_code != 200:
                print(f"  ✗ HTTP {response.status_code}")
                return None

            soup = BeautifulSoup(response.text, "html.parser")
            articles = soup.find_all("article", class_="day-desc")

            if not articles:
                print(f"  ✗ No puzzle content found")
                return None

            # Extract title from first article
            title_elem = articles[0].find("h2")
            if title_elem:
                # Title format: "--- Day X: Title ---"
                title_match = re.search(r"Day \d+: (.+?) ---", title_elem.get_text())
                title = title_match.group(1) if title_match else f"Day {day}"
            else:
                title = f"Day {day}"

            # Convert HTML to Markdown (Part 1 only - first article)
            h2t = html2text.HTML2Text()
            h2t.body_width = 0  # Don't wrap lines
            h2t.ignore_links = False
            h2t.ignore_images = True

            content = h2t.handle(str(articles[0]))

            return title, content.strip()

        except requests.RequestException as e:
            wait_time = BACKOFF_FACTOR ** (attempt + 1)
            print(f"  ⚠ Request error: {e}, retrying in {wait_time}s...")
            time.sleep(wait_time)

    print(f"  ✗ Failed after {MAX_RETRIES} retries")
    return None


def update_readme_day(year: int, year_dir: Path, day: int, title: str) -> None:
    """Update a single day's entry in an existing README.md."""
    readme_path = year_dir / "README.md"
    content = readme_path.read_text()
    dd = f"{day:02d}"
    
    # Pattern to match the day's row: | 04 | [Day 4](...) | ... | ... | ... |
    # or | 04 | [Some Title](...) | ... | ... | ... |
    pattern = rf"^\| {dd} \| \[.*?\]\({AOC_BASE_URL}/{year}/day/{day}\) \|(.*)$"
    
    # Build the new puzzle link with the title
    puzzle_link = f"[{title}]({AOC_BASE_URL}/{year}/day/{day})"
    
    # Find and replace the line
    new_content = re.sub(
        pattern,
        f"| {dd} | {puzzle_link} |\\1",
        content,
        flags=re.MULTILINE
    )
    
    if new_content != content:
        readme_path.write_text(new_content)
        print(f"  📝 Updated README.md (Day {day}: {title})")
    else:
        print(f"  ⚠ Could not find Day {day} entry in README.md")


def create_readme(year: int, year_dir: Path, puzzles: dict[int, str]) -> None:
    """Create or update README.md for a year with puzzle links."""
    readme_path = year_dir / "README.md"

    # If README exists, update only the specific days
    if readme_path.exists():
        for day, title in puzzles.items():
            update_readme_day(year, year_dir, day, title)
        return

    # Build progress table for new README
    rows = []
    for day in range(1, 26):
        dd = f"{day:02d}"
        if day in puzzles:
            title = puzzles[day]
            puzzle_link = f"[{title}]({AOC_BASE_URL}/{year}/day/{day})"
            solution_link = f"[solution](day_{dd}/)"
        else:
            puzzle_link = f"[Day {day}]({AOC_BASE_URL}/{year}/day/{day})"
            solution_link = ""
        rows.append(f"| {dd} | {puzzle_link} | | | {solution_link} |")

    table = "\n".join(rows)

    content = f"""# Advent of Code {year} 🎄

My solutions for [Advent of Code {year}]({AOC_BASE_URL}/{year}).

## Progress

| Day | Puzzle | Part 1 | Part 2 | Solution |
|-----|--------|--------|--------|----------|
{table}

## Running Solutions

Each day's solution is in its own folder (`day_XX/`). To run a solution:

```bash
cd day_01
# Run with appropriate command for the language
```

## Structure

```
aoc_{year}/
├── README.md
├── day_01/
│   ├── 01.md        # Puzzle description (git-ignored)
│   ├── input.txt    # Puzzle input (git-ignored)
│   └── solution.*   # Solution file
└── ...
```

## Legal Notice

Puzzle text and descriptions are © Advent of Code and are not included in this repository.
Links are provided to the original puzzles on [adventofcode.com]({AOC_BASE_URL}/{year}).
Input files are personal and git-ignored.

## Disclaimer

The docstrings and comments in the solution files were mostly generated with the assistance of AI (GitHub Copilot).
"""

    readme_path.write_text(content)
    print(f"  📝 Created README.md")


def main():
    parser = argparse.ArgumentParser(
        description="Fetch Advent of Code puzzle descriptions",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  python fetch_aoc.py --years 2015-2025
  python fetch_aoc.py --years 2024,2025 --days 1-10
  python fetch_aoc.py --years 2023 --force --delay 3
  python fetch_aoc.py --years 2015-2025 --dry-run
        """,
    )
    parser.add_argument(
        "--years",
        required=True,
        help="Years to fetch (e.g., '2015-2025' or '2024,2025')",
    )
    parser.add_argument(
        "--days",
        default="1-25",
        help="Days to fetch (e.g., '1-25' or '1,2,3'). Default: 1-25",
    )
    parser.add_argument(
        "--force",
        action="store_true",
        help="Overwrite existing puzzle description files",
    )
    parser.add_argument(
        "--delay",
        type=float,
        default=4.0,
        help="Delay between requests in seconds. Default: 2.0",
    )
    parser.add_argument(
        "--dry-run",
        action="store_true",
        help="Show what would be fetched without writing files",
    )
    parser.add_argument(
        "-y", "--yes",
        action="store_true",
        help="Skip confirmation prompt (use responsibly)",
    )

    args = parser.parse_args()

    # Parse years and days
    years = parse_range(args.years)
    days = parse_range(args.days)

    # Validate years
    current_year = datetime.now().year
    years = [y for y in years if FIRST_AOC_YEAR <= y <= current_year]
    if not years:
        print(f"No valid years specified. AoC runs from {FIRST_AOC_YEAR} to {current_year}.")
        sys.exit(1)

    # Validate days
    days = [d for d in days if 1 <= d <= 25]
    if not days:
        print("No valid days specified. Days must be 1-25.")
        sys.exit(1)

    # Get repo root (where this script is located)
    repo_root = Path(__file__).parent.resolve()

    # Calculate total potential fetches and show confirmation if needed
    total_potential_fetches = len(years) * len(days)
    needs_confirmation = (
        not args.yes 
        and not args.dry_run 
        and (len(days) > CONFIRM_THRESHOLD_DAYS or len(years) > 1)
    )
    
    if needs_confirmation:
        if not show_disclaimer_and_confirm(total_potential_fetches, years, days):
            print("\n❌ Operation cancelled.")
            sys.exit(0)
        print()  # Add spacing after confirmation

    # Stats
    stats = {"fetched": 0, "skipped": 0, "failed": 0, "created_dirs": 0}

    print(f"🎄 Advent of Code Puzzle Fetcher")
    print(f"   Years: {min(years)}-{max(years)} ({len(years)} years)")
    print(f"   Days: {min(days)}-{max(days)} ({len(days)} days)")
    print(f"   Delay: {args.delay}s between requests")
    if args.dry_run:
        print(f"   Mode: DRY RUN (no files will be written)")
    if args.force:
        print(f"   Mode: FORCE (overwriting existing files)")
    print()

    for year in years:
        print(f"📅 Year {year}")

        # Check available days
        available = get_available_days(year)
        if available == 0:
            print(f"  ⏭ No puzzles available yet for {year}")
            continue

        year_dir = repo_root / f"aoc_{year}"
        puzzles_fetched: dict[int, str] = {}

        # Create year directory if needed
        if not year_dir.exists():
            if not args.dry_run:
                year_dir.mkdir(parents=True)
                stats["created_dirs"] += 1
            print(f"  📁 Created aoc_{year}/")

        for day in days:
            if day > available:
                print(f"  ⏭ Day {day:02d}: Not yet available")
                continue

            dd = f"{day:02d}"
            day_dir = year_dir / f"day_{dd}"
            md_file = day_dir / f"{dd}.md"

            # Check if file exists
            if md_file.exists() and not args.force:
                print(f"  ⏭ Day {dd}: Already exists (use --force to overwrite)")
                stats["skipped"] += 1
                # Try to extract title from existing file for README
                try:
                    existing = md_file.read_text()
                    title_match = re.search(r"^# Day \d+: (.+)$", existing, re.MULTILINE)
                    if title_match:
                        puzzles_fetched[day] = title_match.group(1)
                except Exception:
                    pass
                continue

            if args.dry_run:
                print(f"  🔍 Day {dd}: Would fetch")
                stats["fetched"] += 1
                continue

            print(f"  🔄 Day {dd}: Fetching...", end="", flush=True)
            result = fetch_puzzle(year, day, args.delay)

            if result is None:
                stats["failed"] += 1
                continue

            title, content = result
            puzzles_fetched[day] = title

            # Create day directory
            if not day_dir.exists():
                day_dir.mkdir(parents=True)
                stats["created_dirs"] += 1

            # Write puzzle file
            fetch_date = datetime.now().strftime("%Y-%m-%d")
            file_content = f"""<!-- Fetched from {AOC_BASE_URL}/{year}/day/{day} on {fetch_date} -->
<!-- Part 1 only — Part 2 requires authentication -->

# Day {day}: {title}

{content}
"""
            md_file.write_text(file_content)
            print(f"\r  ✓ Day {dd}: {title}")
            stats["fetched"] += 1

        # Create/update README for this year
        if puzzles_fetched and not args.dry_run:
            create_readme(year, year_dir, puzzles_fetched)

        print()

    # Summary
    print("=" * 50)
    print(f"📊 Summary:")
    print(f"   Fetched: {stats['fetched']}")
    print(f"   Skipped: {stats['skipped']}")
    print(f"   Failed:  {stats['failed']}")
    print(f"   Directories created: {stats['created_dirs']}")

    if args.dry_run:
        print("\n💡 This was a dry run. No files were written.")
        print("   Remove --dry-run to actually fetch and save puzzles.")


if __name__ == "__main__":
    main()
