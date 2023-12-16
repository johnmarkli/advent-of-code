// Package main
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/johnmarkli/advent-of-code/aoc2019"
	"github.com/johnmarkli/advent-of-code/aoc2022"
	"github.com/johnmarkli/advent-of-code/aoc2023"
)

var parts = map[string][][]func(string) any{
	"2019": {
		{aoc2019.Day1Part1, aoc2019.Day1Part2},
		{aoc2019.Day2Part1, aoc2019.Day2Part2},
		{aoc2019.Day3Part1},
		{aoc2019.Day4Part1},
	},
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
		{aoc2023.Day4Part1, aoc2023.Day4Part2},
		{aoc2023.Day5Part1, aoc2023.Day5Part2},
		{aoc2023.Day6Part1, aoc2023.Day6Part2},
		{aoc2023.Day7Part1, aoc2023.Day7Part2},
		{aoc2023.Day8Part1, aoc2023.Day8Part2},
		{aoc2023.Day9Part1, aoc2023.Day9Part2},
		{aoc2023.Day10Part1, aoc2023.Day10Part2},
	},
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Year required\n")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Printf("input directory required\n")
		os.Exit(1)
	}
	year := os.Args[1]
	day, _ := strconv.Atoi(os.Args[2])
	part, _ := strconv.Atoi(os.Args[3])
	inputDir := os.Args[4]
	inputData := fmt.Sprintf("%s/day%d", inputDir, day)
	f := parts[year][day-1][part-1]
	fmt.Println(fmt.Sprintf("Day %d - Part %d:", day, part), f(inputData))
}
