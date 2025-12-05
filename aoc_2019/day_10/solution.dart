/// Advent of Code 2019 - Day 10: Monitoring Station
/// https://adventofcode.com/2019/day/10

import 'dart:io';

String parseInput(String filename) {
  return File(filename).readAsStringSync().trim();
}

int part1(String data) {
  // TODO: Implement solution
  return 0;
}

int part2(String data) {
  // TODO: Implement solution
  return 0;
}

void test() {
  var exampleData = '';

  // Part 1: TODO - add expected value
  // assert(part1(exampleData) == X, 'Part 1 failed');
  
  // Part 2: TODO - add expected value
  // assert(part2(exampleData) == X, 'Part 2 failed');
  
  print('All tests passed!');
}

void main() {
  test();
  
  var data = parseInput('input.txt');
  
  var result1 = part1(data);
  print('Part 1: $result1');
  
  var result2 = part2(data);
  print('Part 2: $result2');
}
