const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();

    const alloc = gpa.allocator();

    const cwd = std.fs.cwd();
    const data = try cwd.readFileAlloc(alloc, "02.txt", 32 * 1024);
    defer alloc.free(data);

    var invalid: u64 = 0;

    var ranges = std.mem.tokenizeAny(u8, data, ",");
    while (ranges.next()) |range| {
        if (range.len == 0) {
            continue;
        }

        var parts = std.mem.splitAny(u8, range, "-");
        const leftStr = parts.next() orelse "";
        const rightStr = parts.next() orelse "";

        const left = try std.fmt.parseInt(u64, leftStr, 10);
        const right = try std.fmt.parseInt(u64, rightStr, 10);

        std.debug.print("Range: {s} - {s}\n", .{ leftStr, rightStr });

        var i: u64 = left;
        while (i <= right) : (i += 1) {
            const iStr = try std.fmt.allocPrint(alloc, "{d}", .{i});

            for (2..iStr.len + 1) |n| {
                if (@rem(iStr.len, n) != 0) {
                    continue;
                }

                const chunkSize = @divTrunc(iStr.len, n);
                var chunks = try alloc.alloc([]const u8, n);
                defer alloc.free(chunks);

                for (0..n) |j| {
                    const start = j * chunkSize;
                    const end = start + chunkSize;
                    chunks[j] = iStr[start..end];
                }

                var allSame = true;
                for (1..chunks.len) |j| {
                    if (!std.mem.eql(u8, chunks[0], chunks[j])) {
                        allSame = false;
                        break;
                    }
                }

                if (allSame) {
                    std.debug.print(" - Invalid number: {s}\n", .{iStr});
                    invalid += i;
                    std.debug.print("   Invalid count so far: {d}\n", .{invalid});
                    break;
                }

                // You can use 'n' here for further processing as needed
            }

            alloc.free(iStr);
        }
    }

    std.debug.print("Invalid count: {d}\n", .{invalid});
}

const Direction = enum { Left, Right };
