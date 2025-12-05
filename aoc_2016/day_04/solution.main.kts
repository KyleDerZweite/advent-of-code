/**
 * Advent of Code 2016 - Day 04: Security Through Obscurity
 * https://adventofcode.com/2016/day/4
 */

import java.io.File

fun parseInput(filename: String): String {
    return File(filename).readText().trim()
}

fun part1(data: String): Int {
    // TODO: Implement solution
    return 0
}

fun part2(data: String): Int {
    // TODO: Implement solution
    return 0
}

fun test() {
    val exampleData = """"""

    // Part 1: TODO - add expected value
    // check(part1(exampleData) == X) { "Part 1 failed" }
    
    // Part 2: TODO - add expected value
    // check(part2(exampleData) == X) { "Part 2 failed" }
    
    println("All tests passed!")
}

fun main() {
    test()
    
    val data = parseInput("input.txt")
    
    val result1 = part1(data)
    println("Part 1: $result1")
    
    val result2 = part2(data)
    println("Part 2: $result2")
}

main()
