const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();

    const alloc = gpa.allocator();

    const cwd = std.fs.cwd();
    const data = try cwd.readFileAlloc(alloc, "04.txt", 32 * 1024);
    defer alloc.free(data);

    var totalRolls: u64 = 0;

    var grid: [200][200]u8 = undefined;

    var lines = std.mem.tokenizeAny(u8, data, "\n");
    var r: usize = 0;
    while (lines.next()) |line| : (r += 1) {
        for (line, 0..) |char, c| {
            grid[r][c] = char;
        }
    }

    while (true) {
        var rolls: u64 = 0;

        for (grid, 0..) |line, row| {
            for (line, 0..) |char, col| {
                if (char != '@') continue;

                var count: u8 = 0;

                const directions = [_][2]i32{
                    .{ -1, -1 }, .{ -1, 0 }, .{ -1, 1 },
                    .{ 0, -1 },  .{ 0, 1 },  .{ 1, -1 },
                    .{ 1, 0 },   .{ 1, 1 },
                };

                for (directions) |dir| {
                    const new_row = @as(i32, @intCast(row)) + dir[0];
                    const new_col = @as(i32, @intCast(col)) + dir[1];

                    if (new_row < 0 or new_row >= grid.len) continue;
                    if (new_col < 0 or new_col >= line.len) continue;

                    const c = grid[@intCast(new_row)][@intCast(new_col)];
                    if (c != '@') continue;

                    count += 1;
                }

                if (count <= 3) {
                    rolls += 1;
                    grid[row][col] = '.';
                }
            }
        }

        if (rolls == 0) break;

        totalRolls += rolls;
    }

    std.debug.print("Total rolls: {d}\n", .{totalRolls});
}
