// Advent of Code 2023 - Day 20: Pulse Propagation
// https://adventofcode.com/2023/day/20

use std::fs;

fn parse_input(filename: &str) -> String {
    fs::read_to_string(filename)
        .expect("Failed to read input file")
        .trim()
        .to_string()
}

fn part1(_data: &str) -> i64 {
    // TODO: Implement solution
    0
}

fn part2(_data: &str) -> i64 {
    // TODO: Implement solution
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let _example_data = "";
        // TODO: Add expected value
        // assert_eq!(part1(example_data), expected);
    }

    #[test]
    fn test_part2() {
        let _example_data = "";
        // TODO: Add expected value
        // assert_eq!(part2(example_data), expected);
    }
}

fn main() {
    let data = parse_input("input.txt");

    let result1 = part1(&data);
    println!("Part 1: {}", result1);

    let result2 = part2(&data);
    println!("Part 2: {}", result2);
}
