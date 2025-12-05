#=
Advent of Code 2015 - Day 22: Wizard Simulator 20XX
https://adventofcode.com/2015/day/22
=#

function parse_input(filename::String)::String
    return read(filename, String)
end

function part1(data::String)::Int
    # TODO: Implement solution
    return 0
end

function part2(data::String)::Int
    # TODO: Implement solution
    return 0
end

function test()
    """Test with example data from the puzzle."""
    example_data = """"""

    # Part 1: TODO - add expected value
    # @assert part1(example_data) == X "Part 1 failed"
    
    # Part 2: TODO - add expected value
    # @assert part2(example_data) == X "Part 2 failed"
    
    println("All tests passed!")
end

function main()
    test()
    
    data = parse_input("input.txt")
    
    result1 = part1(data)
    println("Part 1: $result1")
    
    result2 = part2(data)
    println("Part 2: $result2")
end

main()
