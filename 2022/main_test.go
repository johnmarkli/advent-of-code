package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay1(t *testing.T) {
	require.Equal(t, 24000, Day1Part1("testdata/day1_example"))
	require.Equal(t, 45000, Day1Part2("testdata/day1_example"))
}

func TestDay2(t *testing.T) {
	require.Equal(t, 15, Day2Part1("testdata/day2_example"))
	require.Equal(t, 12, Day2Part2("testdata/day2_example"))
}

func TestDay3(t *testing.T) {
	require.Equal(t, 157, Day3Part1("testdata/day3_example"))
	require.Equal(t, 70, Day3Part2("testdata/day3_example"))
}

func TestDay4(t *testing.T) {
	require.Equal(t, 2, Day4Part1("testdata/day4_example"))
	require.Equal(t, 4, Day4Part2("testdata/day4_example"))
}

func TestDay5(t *testing.T) {
	require.Equal(t, "CMZ", Day5Part1("testdata/day5_example"))
	require.Equal(t, "MCD", Day5Part2("testdata/day5_example"))
}

func TestDay6(t *testing.T) {
	require.Equal(t, 7, Day6Part1("testdata/day6_ex1"))
	require.Equal(t, 5, Day6Part1("testdata/day6_ex2"))
	require.Equal(t, 6, Day6Part1("testdata/day6_ex3"))
	require.Equal(t, 10, Day6Part1("testdata/day6_ex4"))
	require.Equal(t, 11, Day6Part1("testdata/day6_ex5"))

	require.Equal(t, 19, Day6Part2("testdata/day6_ex1"))
	require.Equal(t, 23, Day6Part2("testdata/day6_ex2"))
	require.Equal(t, 23, Day6Part2("testdata/day6_ex3"))
	require.Equal(t, 29, Day6Part2("testdata/day6_ex4"))
	require.Equal(t, 26, Day6Part2("testdata/day6_ex5"))
}

func TestDay7(t *testing.T) {
	require.Equal(t, 95437, Day7Part1("testdata/day7_ex"))
	require.Equal(t, 24933642, Day7Part2("testdata/day7_ex"))
}
