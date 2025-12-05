(*
  Advent of Code 2018 - Day 23: Experimental Emergency Teleportation
  https://adventofcode.com/2018/day/23
*)

let parse_input filename =
  let ic = open_in filename in
  let n = in_channel_length ic in
  let s = really_input_string ic n in
  close_in ic;
  String.trim s

let part1 data =
  (* TODO: Implement solution *)
  0

let part2 data =
  (* TODO: Implement solution *)
  0

let test () =
  let example_data = "" in
  
  (* Part 1: TODO - add expected value *)
  (* assert (part1 example_data = X); *)
  
  (* Part 2: TODO - add expected value *)
  (* assert (part2 example_data = X); *)
  
  print_endline "All tests passed!"

let () =
  test ();
  
  let data = parse_input "input.txt" in
  
  let result1 = part1 data in
  Printf.printf "Part 1: %d\n" result1;
  
  let result2 = part2 data in
  Printf.printf "Part 2: %d\n" result2
