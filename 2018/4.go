package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("December 4th, 2018")

	input := `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-01 00:25] wakes up
[1518-11-02 00:50] wakes up
[1518-11-05 00:45] falls asleep
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-03 00:29] wakes up
[1518-11-05 00:55] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-04 00:36] falls asleep
[1518-11-05 00:03] Guard #99 begins shift`

	log := strings.Split(input, "\n")
	sort.Strings(log)
	guard := ""
	sleeping := false
	schedule := map[string]map[int]int{}
	for i := range log {
		minute, _ := strconv.ParseInt(log[i][15:17], 10, 64)
		if log[i][19] == byte('G') {
			guard = strings.Split(strings.Split(log[i], "#")[1], " ")[0]
			if _, ok := schedule[guard]; !ok {
				schedule[guard] = map[int]int{}
			}
		} else if log[i][19] == byte('f') {
			sleeping = true
		} else if log[i][19] == byte('w') {
			if sleeping {
				firstMinute, _ := strconv.ParseInt(log[i-1][15:17], 10, 64)
				lastMinute := minute - 1
				for i := firstMinute; i <= lastMinute; i++ {
					schedule[guard][int(i)]++
				}
			}
			sleeping = false
		}
	}
	// fmt.Printf("Schedule: %+v\n", schedule)
	mostGuard := ""
	mostMinutes := 0
	for guard, minutes := range schedule {
		numminutes := 0
		for _, num := range minutes {
			numminutes += num
		}
		if numminutes > mostMinutes {
			mostMinutes = numminutes
			mostGuard = guard
		}
	}
	// fmt.Printf("Most sleepy guard: %s (%d minutes)\n", mostGuard, mostMinutes)
	mostMinute := 0
	mostMinutes = 0
	for minute, num := range schedule[mostGuard] {
		if num > mostMinutes {
			mostMinute = minute
			mostMinutes = num
		}
	}
	// fmt.Printf("Most popular minute: %d\n", mostMinute)
	guardID, _ := strconv.ParseInt(mostGuard, 10, 64)
	fmt.Printf("Result (part 1): %v\n", mostMinute*int(guardID))

	guardOnMinute := ""
	minuteOnMinute := 0
	numMinutes := 0
	for guard, minutes := range schedule {
		for minute, num := range minutes {
			if num > numMinutes {
				numMinutes = num
				minuteOnMinute = minute
				guardOnMinute = guard
			}
		}
	}
	guardID, _ = strconv.ParseInt(guardOnMinute, 10, 64)
	fmt.Printf("Result (part 2): %v\n", minuteOnMinute*int(guardID))
}
