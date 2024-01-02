package aoc2023

import (
	"fmt"
	"math"
)

type coord struct {
	x int
	y int
}

func (c *coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c *coord) outOfBounds(maxX, maxY int) bool {
	return c.x < 0 || c.x > maxX || c.y < 0 || c.y > maxY
}

func (c *coord) sameAs(c2 *coord) bool {
	return c.x == c2.x && c.y == c2.y
}

// manhattan distance |x1-x2| + |y1-y2|
func (c *coord) dist(c2 *coord) int {
	return int(math.Abs(float64(c.x)-float64(c2.x)) + math.Abs(float64(c.y)-float64(c2.y)))
}

func (c *coord) up() *coord {
	return &coord{
		x: c.x - 1,
		y: c.y,
	}
}

func (c *coord) down() *coord {
	return &coord{
		x: c.x + 1,
		y: c.y,
	}
}

func (c *coord) left() *coord {
	return &coord{
		x: c.x,
		y: c.y - 1,
	}
}

func (c *coord) right() *coord {
	return &coord{
		x: c.x,
		y: c.y + 1,
	}
}
