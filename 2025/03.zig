const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();

    const alloc = gpa.allocator();

    const cwd = std.fs.cwd();
    const data = try cwd.readFileAlloc(alloc, "03.txt", 32 * 1024);
    defer alloc.free(data);

    var power: u64 = 0;

    var lines = std.mem.tokenizeAny(u8, data, "\n");
    while (lines.next()) |line| {
        if (line.len == 0) {
            continue;
        }

        var nums = try alloc.alloc(u8, line.len);
        defer alloc.free(nums);

        for (line, 0..) |char, i| {
            nums[i] = try std.fmt.parseInt(u8, &[_]u8{char}, 10);
        }

        var joltage: [12]u8 = .{0} ** 12;
        var res: Result = undefined;

        for (0..12) |i| {
            const from = nums.len - (12 - i);
            const to = if (i == 0) 0 else res.index + 1;
            res = findHighest(nums, from, to);
            joltage[i] = res.value;
        }

        var joltageStr: [12]u8 = undefined;
        for (joltage, 0..) |value, i| {
            joltageStr[i] = '0' + value;
        }

        std.debug.print("{s}\n", .{joltageStr});

        const joltageNum = try std.fmt.parseInt(u64, &joltageStr, 10);

        power += joltageNum;
    }

    std.debug.print("Total power: {d}\n", .{power});
}

fn findHighest(bank: []const u8, from: usize, to: usize) Result {
    var highest: u8 = 0;
    var at: usize = from;

    var i = from;
    while (i >= to) : (i -= 1) {
        const n = bank[i];
        if (n >= highest) {
            highest = n;
            at = i;
        }
        if (i == 0) break;
    }

    return .{ .value = highest, .index = at };
}

const Result = struct {
    value: u8,
    index: usize,
};
