-- Advent of Code 2020 - Day 23: Crab Cups
-- https://adventofcode.com/2020/day/23

import System.IO

parseInput :: String -> String
parseInput = id

part1 :: String -> Int
part1 _ = 0  -- TODO: Implement solution

part2 :: String -> Int
part2 _ = 0  -- TODO: Implement solution

test :: IO ()
test = do
    let exampleData = ""
    
    -- Part 1: TODO - add expected value
    -- assert (part1 exampleData == x) "Part 1 failed"
    
    -- Part 2: TODO - add expected value
    -- assert (part2 exampleData == x) "Part 2 failed"
    
    putStrLn "All tests passed!"

main :: IO ()
main = do
    test
    
    contents <- readFile "input.txt"
    let input = parseInput contents
    
    putStrLn $ "Part 1: " ++ show (part1 input)
    putStrLn $ "Part 2: " ++ show (part2 input)
