"""
Advent of Code 2024 - Day 1: Historian Hysteria
https://adventofcode.com/2024/day/1
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


if __name__ == "__main__":
    left, right = parse_input("input.txt")
    
    result1 = part1(left, right)
    print(f"Part 1 - Total distance: {result1}")
    
    result2 = part2(left, right)
    print(f"Part 2 - Similarity score: {result2}")
