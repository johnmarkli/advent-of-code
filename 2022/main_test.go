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
