package aoc2022_test

import (
	"testing"

	"github.com/johnmarkli/advent-of-code/aoc2022"
	"github.com/stretchr/testify/require"
)

func TestDay1(t *testing.T) {
	require.Equal(t, 24000, aoc2022.Day1Part1("testdata/day1_example"))
	require.Equal(t, 45000, aoc2022.Day1Part2("testdata/day1_example"))
}

func TestDay2(t *testing.T) {
	require.Equal(t, 15, aoc2022.Day2Part1("testdata/day2_example"))
	require.Equal(t, 12, aoc2022.Day2Part2("testdata/day2_example"))
}

func TestDay3(t *testing.T) {
	require.Equal(t, 157, aoc2022.Day3Part1("testdata/day3_example"))
	require.Equal(t, 70, aoc2022.Day3Part2("testdata/day3_example"))
}

func TestDay4(t *testing.T) {
	require.Equal(t, 2, aoc2022.Day4Part1("testdata/day4_example"))
	require.Equal(t, 4, aoc2022.Day4Part2("testdata/day4_example"))
}

func TestDay5(t *testing.T) {
	require.Equal(t, "CMZ", aoc2022.Day5Part1("testdata/day5_example"))
	require.Equal(t, "MCD", aoc2022.Day5Part2("testdata/day5_example"))
}

func TestDay6(t *testing.T) {
	require.Equal(t, 7, aoc2022.Day6Part1("testdata/day6_ex1"))
	require.Equal(t, 5, aoc2022.Day6Part1("testdata/day6_ex2"))
	require.Equal(t, 6, aoc2022.Day6Part1("testdata/day6_ex3"))
	require.Equal(t, 10, aoc2022.Day6Part1("testdata/day6_ex4"))
	require.Equal(t, 11, aoc2022.Day6Part1("testdata/day6_ex5"))

	require.Equal(t, 19, aoc2022.Day6Part2("testdata/day6_ex1"))
	require.Equal(t, 23, aoc2022.Day6Part2("testdata/day6_ex2"))
	require.Equal(t, 23, aoc2022.Day6Part2("testdata/day6_ex3"))
	require.Equal(t, 29, aoc2022.Day6Part2("testdata/day6_ex4"))
	require.Equal(t, 26, aoc2022.Day6Part2("testdata/day6_ex5"))
}

func TestDay7(t *testing.T) {
	require.Equal(t, 95437, aoc2022.Day7Part1("testdata/day7_ex"))
	require.Equal(t, 24933642, aoc2022.Day7Part2("testdata/day7_ex"))
}

func TestDay8(t *testing.T) {
	require.Equal(t, 21, aoc2022.Day8Part1("testdata/day8_ex"))
	require.Equal(t, 8, aoc2022.Day8Part2("testdata/day8_ex"))
}

func TestDay9(t *testing.T) {
	// require.Equal(t, 13, aoc2022.Day9Part1("testdata/day9_ex"))
	// require.Equal(t, 1, aoc2022.Day9Part2("testdata/day9_ex"))
	require.Equal(t, 36, aoc2022.Day9Part2("testdata/day9_ex2"))
}
