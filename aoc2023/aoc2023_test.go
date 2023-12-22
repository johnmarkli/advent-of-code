// Package aoc2023_test
package aoc2023_test

import (
	"testing"

	"github.com/johnmarkli/advent-of-code/aoc2023"
	"github.com/stretchr/testify/require"
)

func TestDay1(t *testing.T) {
	require.Equal(t, 142, aoc2023.Day1Part1("testdata/day1"))

	require.Equal(t, 281, aoc2023.Day1Part2("testdata/day1-2"))
}

func TestDay2(t *testing.T) {
	require.Equal(t, 8, aoc2023.Day2Part1("testdata/day2"))

	require.Equal(t, 2286, aoc2023.Day2Part2("testdata/day2"))
}

func TestDay3(t *testing.T) {
	require.Equal(t, 4361, aoc2023.Day3Part1("testdata/day3"))
	require.Equal(t, 4361, aoc2023.Day3Part1("testdata/day3-2"))

	require.Equal(t, 467835, aoc2023.Day3Part2("testdata/day3"))
	require.Equal(t, 467835, aoc2023.Day3Part2("testdata/day3-2"))
	require.Equal(t, 467835, aoc2023.Day3Part2("testdata/day3-3"))
}

func TestDay4(t *testing.T) {
	require.Equal(t, 13, aoc2023.Day4Part1("testdata/day4"))

	require.Equal(t, 30, aoc2023.Day4Part2("testdata/day4"))
}

func TestDay5(t *testing.T) {
	require.Equal(t, 35, aoc2023.Day5Part1("testdata/day5"))

	require.Equal(t, 46, aoc2023.Day5Part2("testdata/day5"))
}

func TestDay6(t *testing.T) {
	require.Equal(t, 288, aoc2023.Day6Part1("testdata/day6"))

	require.Equal(t, 71503, aoc2023.Day6Part2("testdata/day6"))
}

func TestDay7(t *testing.T) {
	require.Equal(t, 6440, aoc2023.Day7Part1("testdata/day7"))

	require.Equal(t, 5905, aoc2023.Day7Part2("testdata/day7"))
}

func TestDay8(t *testing.T) {
	require.Equal(t, 2, aoc2023.Day8Part1("testdata/day8"))
	require.Equal(t, 6, aoc2023.Day8Part1("testdata/day8-2"))

	require.Equal(t, 6, aoc2023.Day8Part2("testdata/day8-3"))
}

func TestDay9(t *testing.T) {
	require.Equal(t, 114, aoc2023.Day9Part1("testdata/day9"))
	require.Equal(t, 2, aoc2023.Day9Part2("testdata/day9"))
}

func TestDay10(t *testing.T) {
	require.Equal(t, 4, aoc2023.Day10Part1("testdata/day10"))
	require.Equal(t, 4, aoc2023.Day10Part1("testdata/day10-2"))
	require.Equal(t, 8, aoc2023.Day10Part1("testdata/day10-3"))
	require.Equal(t, 8, aoc2023.Day10Part1("testdata/day10-4"))

	require.Equal(t, 4, aoc2023.Day10Part2("testdata/day10-5"))
	require.Equal(t, 8, aoc2023.Day10Part2("testdata/day10-6"))
	require.Equal(t, 10, aoc2023.Day10Part2("testdata/day10-7"))
}

func TestDay11(t *testing.T) {
	require.Equal(t, 374, aoc2023.Day11Part1("testdata/day11"))
	require.Equal(t, 7, aoc2023.Day11Part1("testdata/day11-2"))

	require.Equal(t, 1030, aoc2023.Day11Part2Solve("testdata/day11", 10))
	require.Equal(t, 8410, aoc2023.Day11Part2Solve("testdata/day11", 100))
}

func TestDay12(t *testing.T) {
	tcs := []struct {
		sr     *aoc2023.SpringRow
		expect int
	}{
		{
			sr: &aoc2023.SpringRow{
				Springs: []byte("???.###"),
				Groups:  []int{1, 1, 3},
			},
			expect: 1,
		},
		{
			sr: &aoc2023.SpringRow{
				Springs: []byte(".??..??...?##."),
				Groups:  []int{1, 1, 3},
			},
			expect: 4,
		},
		{
			sr: &aoc2023.SpringRow{
				Springs: []byte("?#?#?#?#?#?#?#?"),
				Groups:  []int{1, 3, 1, 6},
			},
			expect: 1,
		},
		{
			sr: &aoc2023.SpringRow{
				Springs: []byte("????.#...#..."),
				Groups:  []int{4, 1, 1},
			},
			expect: 1,
		},
		{
			sr: &aoc2023.SpringRow{
				Springs: []byte("????.######..#####."),
				Groups:  []int{1, 6, 5},
			},
			expect: 4,
		},
		{
			sr: &aoc2023.SpringRow{
				Springs: []byte("?###????????"),
				Groups:  []int{3, 2, 1},
			},
			expect: 10,
		},
	}

	for _, tc := range tcs {
		require.Equal(t, tc.expect, tc.sr.Arrangements())
	}

	require.Equal(t, 21, aoc2023.Day12Part1("testdata/day12"))

	require.Equal(t, 525152, aoc2023.Day12Part2("testdata/day12"))
}

func TestDay13(t *testing.T) {
	require.Equal(t, 405, aoc2023.Day13Part1("testdata/day13"))
	require.Equal(t, 10, aoc2023.Day13Part1("testdata/day13-2"))
	require.Equal(t, 600, aoc2023.Day13Part1("testdata/day13-3"))
	require.Equal(t, 2101, aoc2023.Day13Part1("testdata/day13-4"))
	require.Equal(t, 400, aoc2023.Day13Part1("testdata/day13-9"))

	require.Equal(t, 400, aoc2023.Day13Part2("testdata/day13"))
	require.Equal(t, 5, aoc2023.Day13Part2("testdata/day13-5"))
	require.Equal(t, 200, aoc2023.Day13Part2("testdata/day13-6"))
	require.Equal(t, 5, aoc2023.Day13Part2("testdata/day13-7"))
	require.Equal(t, 100, aoc2023.Day13Part2("testdata/day13-8"))
	require.Equal(t, 1100, aoc2023.Day13Part2("testdata/day13-9"))
	require.Equal(t, 800, aoc2023.Day13Part2("testdata/day13-10"))
}

func TestDay14(t *testing.T) {
	require.Equal(t, 136, aoc2023.Day14Part1("testdata/day14"))

	require.Equal(t, 64, aoc2023.Day14Part2("testdata/day14"))
}
