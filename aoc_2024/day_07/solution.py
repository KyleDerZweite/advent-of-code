"""
Advent of Code 2024 - Day 07: Bridge Repair
https://adventofcode.com/2024/day/7
"""


def parse_input(filename: str) -> str:
    with open(filename, 'r') as f:
        data = f.read()
    
    return data


def part1(data: str) -> int:
    def can_make(target: int, nums: list[int]) -> bool:
        possible = {nums[0]}

        for n in nums[1:]:
            next_vals = set()
            for val in possible:
                add_val = val + n
                mul_val = val * n

                if add_val <= target:
                    next_vals.add(add_val)
                if mul_val <= target:
                    next_vals.add(mul_val)

            if not next_vals:
                return False
            possible = next_vals

        return target in possible

    total = 0

    for line in data.strip().splitlines():
        if not line.strip():
            continue

        target_str, nums_str = line.split(":")
        target = int(target_str)
        nums = list(map(int, nums_str.split()))

        if can_make(target, nums):
            total += target

    return total


def part2(data: str) -> int:
    def concat(a: int, b: int) -> int:
        # Concatenate digits of a and b
        pow_10 = 10 ** len(str(b))
        return a * pow_10 + b

    def can_make(target: int, nums: list[int]) -> bool:
        possible = {nums[0]}

        for n in nums[1:]:
            next_vals = set()
            for val in possible:
                add_val = val + n
                mul_val = val * n
                cat_val = concat(val, n)

                if add_val <= target:
                    next_vals.add(add_val)
                if mul_val <= target:
                    next_vals.add(mul_val)
                if cat_val <= target:
                    next_vals.add(cat_val)

            if not next_vals:
                return False
            possible = next_vals

        return target in possible

    total = 0

    for line in data.strip().splitlines():
        if not line.strip():
            continue

        target_str, nums_str = line.split(":")
        target = int(target_str)
        nums = list(map(int, nums_str.split()))

        if can_make(target, nums):
            total += target

    return total


def test():
    """Test with example data from the puzzle."""
    example_data = """190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20"""

    # Part 1: Total calibration result is 3749
    assert part1(example_data) == 3749, "Part 1 failed"
    
    # Part 2: Total calibration result is 11387
    assert part2(example_data) == 11387, "Part 2 failed"
    
    print("All tests passed!")


if __name__ == "__main__":
    test()
    
    data = parse_input("input.txt")

    result1 = part1(data)
    print(f"Part 1: {result1}")
    
    result2 = part2(data)
    print(f"Part 2: {result2}")
