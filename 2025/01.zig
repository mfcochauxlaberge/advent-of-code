const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();

    const alloc = gpa.allocator();

    const cwd = std.fs.cwd();
    const data = try cwd.readFileAlloc(alloc, "01.txt", 32 * 1024);
    defer alloc.free(data);

    var dial: i16 = 50;
    var pw: u14 = 0;
    var extra: u14 = 0;

    var lines = std.mem.tokenizeAny(u8, data, "\n");
    while (lines.next()) |line| {
        const startedAt0 = dial == 0;

        if (line.len == 0) {
            continue;
        }

        var dir: ?Direction = null;

        if (line[0] == 'L') {
            dir = Direction.Left;
        } else if (line[0] == 'R') {
            dir = Direction.Right;
        } else {
            std.debug.print("Invalid direction: {s}\n", .{line});
            return error.InvalidDirection;
        }

        var v: u11 = 0;

        const rest = line[1..];
        v = try std.fmt.parseInt(u11, rest, 10);

        extra += @divTrunc(v, 100);
        std.debug.print("Extra 0s: {d}\n", .{@divTrunc(v, 100)});
        v = @rem(v, 100);

        std.debug.print("Dial: {d} ({s} {d}, {d})\n", .{ dial, @tagName(dir.?), v, extra });

        switch (dir.?) {
            Direction.Left => dial -= @intCast(v),
            Direction.Right => dial += @intCast(v),
        }

        std.debug.print("Dial position after: {d} ({d}, {d})\n", .{ dial, @divTrunc(dial, 100), 100 * (@divTrunc(dial, 100) + 1) });

        if (dial < 0) {
            if (!startedAt0) {
                extra += 1;
                std.debug.print("One extra added for passing 0\n", .{});
            }
            dial += 100 * (@divTrunc(dial * -1, 100) + 1);
        } else if (dial >= 100) {
            if (!startedAt0) {
                extra += 1;
                std.debug.print("One extra added for passing 0\n", .{});
            }
            dial -= 100 * @divTrunc(dial, 100);
        } else if (dial == 0) {
            extra += 1;
            std.debug.print("One extra added for stopping on 0\n", .{});
        }

        std.debug.print("Dial position after correction: {d}\n", .{dial});
        std.debug.print("Extra 0s now at: {d}\n", .{extra});

        if (dial == 0) {
            pw += 1;
        }

        std.debug.print("\n", .{});
    }

    std.debug.print("Final dial position: {d}\n", .{dial});
    std.debug.print("Password: {d}\n", .{extra});

    // std.debug.print("{s}", .{data});
}

const Direction = enum { Left, Right };
