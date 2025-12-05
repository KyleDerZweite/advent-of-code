// Advent of Code 2022 - Day 08: Treetop Tree House
// https://adventofcode.com/2022/day/8

const std = @import("std");

fn parseInput(allocator: std.mem.Allocator, filename: []const u8) ![]const u8 {
    const file = try std.fs.cwd().openFile(filename, .{});
    defer file.close();
    return try file.readToEndAlloc(allocator, std.math.maxInt(usize));
}

fn part1(data: []const u8) i64 {
    // TODO: Implement solution
    _ = data;
    return 0;
}

fn part2(data: []const u8) i64 {
    // TODO: Implement solution
    _ = data;
    return 0;
}

fn test_solutions() !void {
    const example_data = "";

    // Part 1: TODO - add expected value
    // try std.testing.expectEqual(@as(i64, x), part1(example_data));

    // Part 2: TODO - add expected value
    // try std.testing.expectEqual(@as(i64, x), part2(example_data));

    _ = example_data;
    std.debug.print("All tests passed!\n", .{});
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    try test_solutions();

    const data = try parseInput(allocator, "input.txt");
    defer allocator.free(data);

    const result1 = part1(data);
    std.debug.print("Part 1: {}\n", .{result1});

    const result2 = part2(data);
    std.debug.print("Part 2: {}\n", .{result2});
}
