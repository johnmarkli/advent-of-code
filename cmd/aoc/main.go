// Package main
package main

import (
	"fmt"
	"os"

	"github.com/johnmarkli/advent-of-code/aoc2022"
	"github.com/johnmarkli/advent-of-code/aoc2023"
)

var parts = map[string][][]func(string) any{
	"2022": {
		{aoc2022.Day1Part1, aoc2022.Day1Part2},
		{aoc2022.Day2Part1, aoc2022.Day2Part2},
		{aoc2022.Day3Part1, aoc2022.Day3Part2},
		{aoc2022.Day4Part1, aoc2022.Day4Part2},
		{aoc2022.Day5Part1, aoc2022.Day5Part2},
		{aoc2022.Day6Part1, aoc2022.Day6Part2},
		{aoc2022.Day7Part1, aoc2022.Day7Part2},
		{aoc2022.Day8Part1, aoc2022.Day8Part2},
		{aoc2022.Day9Part1, aoc2022.Day9Part2},
	},
	"2023": {
		{aoc2023.Day1Part1, aoc2023.Day1Part2},
		{aoc2023.Day2Part1, aoc2023.Day2Part2},
		{aoc2023.Day3Part1, aoc2023.Day3Part2},
	},
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Year required\n")
		os.Exit(1)
	}
	year := os.Args[1]
	for i, fs := range parts[year] {
		for j, f := range fs {
			day := i + 1
			part := j + 1
			testData := fmt.Sprintf("testdata/day%d", day)
			fmt.Println(fmt.Sprintf("Day %d - Part %d:", day, part), f(testData))
		}
	}
}
