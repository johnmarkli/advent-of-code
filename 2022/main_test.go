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
