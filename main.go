package main

import (
	"fmt"
	"os"
	"ruan.moolman/aoc2023/day1"
	"ruan.moolman/aoc2023/day10"
	"ruan.moolman/aoc2023/day2"
	"ruan.moolman/aoc2023/day3"
	"ruan.moolman/aoc2023/day4"
	"ruan.moolman/aoc2023/day5"
	"ruan.moolman/aoc2023/day6"
	"ruan.moolman/aoc2023/day7"
	"ruan.moolman/aoc2023/day8"
	"ruan.moolman/aoc2023/day9"
	"time"
)

var days = map[string]map[string]func() string{
	"1": {
		"1": day1.Part1,
		"2": day1.Part2,
	},
	"2": {
		"1": day2.Part1,
		"2": day2.Part2,
	},
	"3": {
		"1": day3.Part1,
		"2": day3.Part2,
	},
	"4": {
		"1": day4.Part1,
		"2": day4.Part2,
	},
	"5": {
		"1": day5.Part1,
		"2": day5.Part2,
	},
	"6": {
		"1": day6.Part1,
		"2": day6.Part2,
	},
	"7": {
		"1": day7.Part1,
		"2": day7.Part2,
	},
	"8": {
		"1": day8.Part1,
		"2": day8.Part2,
	},
	"9": {
		"1": day9.Part1,
		"2": day9.Part2,
	},
	"10": {
		"1": day10.Part1,
		"2": day10.Part2,
	},
}

func main() {
	dayArg := os.Args[1]
	day, ok := days[dayArg]
	if !ok {
		fmt.Printf("Invalid day: %s\n", dayArg)
		os.Exit(1)
	}

	partArg := os.Args[2]
	fn, ok := day[partArg]
	if !ok {
		fmt.Printf("Invalid part: %s\n", partArg)
		os.Exit(1)
	}

	t := time.Now()
	answer := fn()
	fmt.Printf("%d μs\n", time.Now().Sub(t).Microseconds())

	fmt.Printf("Day %s Part %s: %s\n", dayArg, partArg, answer)
}
