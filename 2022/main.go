package main

import "fmt"

var parts = [][]func(string) any{
	{Day1Part1, Day1Part2},
	{Day2Part1, Day2Part2},
	{Day3Part1, Day3Part2},
	{Day4Part1, Day4Part2},
	{Day5Part1, Day5Part2},
	{Day6Part1, Day6Part2},
	{Day7Part1, Day7Part2},
	{Day8Part1, Day8Part2},
}

func main() {
	for i, fs := range parts {
		for j, f := range fs {
			day := i + 1
			part := j + 1
			testData := fmt.Sprintf("testdata/day%d", day)
			fmt.Println(fmt.Sprintf("Day %d - Part %d:", day, part), f(testData))
		}
	}
}
