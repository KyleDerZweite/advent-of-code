# Advent of Code 2021 - Day 05: Hydrothermal Venture
# https://adventofcode.com/2021/day/5

defmodule Solution do
  def parse_input(filename) do
    File.read!(filename) |> String.trim()
  end

  def part1(_data) do
    # TODO: Implement solution
    0
  end

  def part2(_data) do
    # TODO: Implement solution
    0
  end

  def test do
    example_data = ""

    # Part 1: TODO - add expected value
    # assert part1(example_data) == x, "Part 1 failed"
    
    # Part 2: TODO - add expected value
    # assert part2(example_data) == x, "Part 2 failed"

    IO.puts("All tests passed!")
  end

  def main do
    test()

    data = parse_input("input.txt")

    result1 = part1(data)
    IO.puts("Part 1: #{result1}")

    result2 = part2(data)
    IO.puts("Part 2: #{result2}")
  end
end

Solution.main()
