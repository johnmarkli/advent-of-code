package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay1(t *testing.T) {
	res := Day1("testdata/day1_example")
	require.Equal(t, 24000, res)
}
