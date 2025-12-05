//// Advent of Code 2017 - Day 24: Electromagnetic Moat
//// https://adventofcode.com/2017/day/24

import gleam/io
import gleam/string
import simplifile

pub fn parse_input(filename: String) -> String {
  case simplifile.read(filename) {
    Ok(content) -> string.trim(content)
    Error(_) -> ""
  }
}

pub fn part1(data: String) -> Int {
  // TODO: Implement solution
  0
}

pub fn part2(data: String) -> Int {
  // TODO: Implement solution
  0
}

pub fn test() {
  let example_data = ""

  // Part 1: TODO - add expected value
  // let assert True = part1(example_data) == X
  
  // Part 2: TODO - add expected value
  // let assert True = part2(example_data) == X
  
  io.println("All tests passed!")
}

pub fn main() {
  test()
  
  let data = parse_input("input.txt")
  
  let result1 = part1(data)
  io.println("Part 1: " <> string.inspect(result1))
  
  let result2 = part2(data)
  io.println("Part 2: " <> string.inspect(result2))
}
