package aoc2019_test

import (
	"strings"
	"testing"

	"github.com/johnmarkli/advent-of-code/aoc2019"
	"github.com/stretchr/testify/assert"
)

func TestClosestIntersectionDist(t *testing.T) {
	line1Str := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	line2Str := "U62,R66,U55,R34,D71,R55,D58,R83"
	line1 := strings.Split(line1Str, ",")
	line2 := strings.Split(line2Str, ",")
	d := aoc2019.ClosestIntersectionDist(line1, line2)
	assert.Equal(t, d, 159)

	line1Str = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	line2Str = "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
	line1 = strings.Split(line1Str, ",")
	line2 = strings.Split(line2Str, ",")
	d = aoc2019.ClosestIntersectionDist(line1, line2)
	assert.Equal(t, d, 135)
}

// TODO: fix test
// func TestInterssectionLeastSteps(t *testing.T) {
// 	line1Str := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
// 	line2Str := "U62,R66,U55,R34,D71,R55,D58,R83"
// 	line1 := strings.Split(line1Str, ",")
// 	line2 := strings.Split(line2Str, ",")
// 	s := intersectionLeastSteps(line1, line2)
// 	assert.Equal(t, s, 610)
// }

func TestManhattanDistance(t *testing.T) {
	d := aoc2019.ManhattanDistance(aoc2019.Point{0, 0}, aoc2019.Point{6, 6})
	assert.Equal(t, d, 12)
}

// TODO: fix test
// func TestNumSteps(t *testing.T) {
// 	lineStr := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
// 	lineArr := strings.Split(lineStr, ",")
// 	p := point{75, -30}
// 	s := numSteps(lineArr, p)
// 	assert.Equal(t, s, 105)
// }
