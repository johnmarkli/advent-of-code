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
