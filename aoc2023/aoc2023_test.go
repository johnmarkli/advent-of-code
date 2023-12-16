// Package aoc2023_test
package aoc2023_test

import (
	"testing"

	"github.com/johnmarkli/advent-of-code/aoc2023"
	"github.com/stretchr/testify/require"
)

func TestDay1(t *testing.T) {
	require.Equal(t, 142, aoc2023.Day1Part1("testdata/day1_example"))

	require.Equal(t, 281, aoc2023.Day1Part2("testdata/day1_example2"))
}

func TestDay2(t *testing.T) {
	require.Equal(t, 8, aoc2023.Day2Part1("testdata/day2_example"))

	require.Equal(t, 2286, aoc2023.Day2Part2("testdata/day2_example"))
}

func TestDay3(t *testing.T) {
	require.Equal(t, 4361, aoc2023.Day3Part1("testdata/day3_example"))
	require.Equal(t, 4361, aoc2023.Day3Part1("testdata/day3_example2"))

	require.Equal(t, 467835, aoc2023.Day3Part2("testdata/day3_example"))
	require.Equal(t, 467835, aoc2023.Day3Part2("testdata/day3_example2"))
	require.Equal(t, 467835, aoc2023.Day3Part2("testdata/day3_example3"))
}

func TestDay4(t *testing.T) {
	require.Equal(t, 13, aoc2023.Day4Part1("testdata/day4_example"))

	require.Equal(t, 30, aoc2023.Day4Part2("testdata/day4_example"))
}

func TestDay5(t *testing.T) {
	require.Equal(t, 35, aoc2023.Day5Part1("testdata/day5_example"))

	require.Equal(t, 46, aoc2023.Day5Part2("testdata/day5_example"))
}

func TestDay6(t *testing.T) {
	require.Equal(t, 288, aoc2023.Day6Part1("testdata/day6_example"))

	require.Equal(t, 71503, aoc2023.Day6Part2("testdata/day6_example"))
}

func TestDay7(t *testing.T) {
	require.Equal(t, 6440, aoc2023.Day7Part1("testdata/day7_example"))

	require.Equal(t, 5905, aoc2023.Day7Part2("testdata/day7_example"))
}

func TestDay8(t *testing.T) {
	require.Equal(t, 2, aoc2023.Day8Part1("testdata/day8_example"))
	require.Equal(t, 6, aoc2023.Day8Part1("testdata/day8_example2"))

	require.Equal(t, 6, aoc2023.Day8Part2("testdata/day8_example3"))
}

func TestDay9(t *testing.T) {
	require.Equal(t, 114, aoc2023.Day9Part1("testdata/day9_example"))
	require.Equal(t, 2, aoc2023.Day9Part2("testdata/day9_example"))
}

func TestDay10(t *testing.T) {
	// require.Equal(t, 4, aoc2023.Day10Part1("testdata/day10_example"))
	// require.Equal(t, 4, aoc2023.Day10Part1("testdata/day10_example2"))
	// require.Equal(t, 8, aoc2023.Day10Part1("testdata/day10_example3"))
	// require.Equal(t, 8, aoc2023.Day10Part1("testdata/day10_example4"))
	//
	require.Equal(t, 4, aoc2023.Day10Part2("testdata/day10_example5"))
	require.Equal(t, 8, aoc2023.Day10Part2("testdata/day10_example6"))
	require.Equal(t, 10, aoc2023.Day10Part2("testdata/day10_example7"))
}
