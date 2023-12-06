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
