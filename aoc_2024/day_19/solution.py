"""
Advent of Code 2024 - Day 19: Historian Hysteria
https://adventofcode.com/2024/day/19
"""

from collections import Counter


def parse_input(filename: str) -> tuple[list[int], list[int]]:
    """Parse input file into two separate lists."""
    left_list = []
    right_list = []
    
    with open(filename, 'r') as f:
        for line in f:
            if line.strip():
                left, right = line.split()
                left_list.append(int(left))
                right_list.append(int(right))
    
    return left_list, right_list


def part1(left_list: list[int], right_list: list[int]) -> int:
    """
    Calculate the total distance between two lists.
    
    Pairs up numbers from smallest to largest and sums the absolute differences.
    """
    left_sorted = sorted(left_list)
    right_sorted = sorted(right_list)
    
    total_distance = 0
    for left, right in zip(left_sorted, right_sorted):
        total_distance += abs(left - right)
    
    return total_distance


def part2(left_list: list[int], right_list: list[int]) -> int:
    """
    Calculate the similarity score between two lists.
    
    For each number in the left list, multiply it by how many times
    it appears in the right list, then sum all results.
    """
    right_counts = Counter(right_list)
    
    similarity_score = 0
    for num in left_list:
        similarity_score += num * right_counts[num]
    
    return similarity_score


def test():
    """Test with example data from the puzzle."""
    example_left = [3, 4, 2, 1, 3, 3]
    example_right = [4, 3, 5, 3, 9, 3]
    
    # Part 1: distances sum to 11
    assert part1(example_left, example_right) == 11, "Part 1 failed"
    
    # Part 2: similarity score is 31
    assert part2(example_left, example_right) == 31, "Part 2 failed"
    
    print("All tests passed!")


if __name__ == "__main__":
    # test()
    
    left, right = parse_input("input.txt")
    
    result1 = part1(left, right)
    print(f"Part 1 - Total distance: {result1}")
    
    result2 = part2(left, right)
    print(f"Part 2 - Similarity score: {result2}")
