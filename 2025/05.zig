const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();

    const alloc = gpa.allocator();

    const cwd = std.fs.cwd();
    const data = try cwd.readFileAlloc(alloc, "05_test.txt", 32 * 1024);
    defer alloc.free(data);

    var fresh_ranges: [200][2]u64 = undefined;
    var all_ids: [1200]u64 = undefined;
    var total: u64 = 0;
    var range_mode: bool = true;

    var lines = std.mem.tokenizeAny(u8, data, "\n");

    var range_count: usize = 0;
    var id_count: usize = 0;

    while (lines.next()) |line| {
        if (range_mode and std.mem.indexOf(u8, line, "-") == null) {
            range_mode = false;
        }

        if (range_mode) {
            var parts = std.mem.splitAny(u8, line, "-");
            const start = try std.fmt.parseInt(u64, parts.next().?, 10);
            const end = try std.fmt.parseInt(u64, parts.next().?, 10);
            fresh_ranges[range_count] = .{ start, end };
            range_count += 1;
        } else {
            const id = try std.fmt.parseInt(u64, line, 10);
            all_ids[id_count] = id;
            id_count += 1;
        }
    }

    for (all_ids[0..id_count]) |id| {
        for (fresh_ranges[0..range_count]) |range| {
            if (id >= range[0] and id <= range[1]) {
                total += 1;
                break;
            }
        }
    }

    std.debug.print("Total fresh: {d}\n", .{total});
}
