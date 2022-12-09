package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay1Part1(t *testing.T) {
	res := Day1Part1("testdata/day1_example")
	require.Equal(t, 24000, res)
}

func TestDay1Part2(t *testing.T) {
	res := Day1Part2("testdata/day1_example")
	require.Equal(t, 45000, res)
}
