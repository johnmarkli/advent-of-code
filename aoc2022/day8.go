package aoc2022

import (
	"bufio"
	"os"
	"strconv"
)

/*
The expedition comes across a peculiar patch of tall trees all planted
carefully in a grid. The Elves explain that a previous expedition planted these
trees as a reforestation effort. Now, they're curious if this would be a good
location for a tree house.

First, determine whether there is enough tree cover here to keep a tree house
hidden. To do this, you need to count the number of trees that are visible from
outside the grid when looking directly along a row or column.

The Elves have already launched a quadcopter to generate a map with the height
of each tree (your puzzle input). For example:

30373
25512
65332
33549
35390

Each tree is represented as a single digit whose value is its height, where 0
is the shortest and 9 is the tallest.

A tree is visible if all of the other trees between it and an edge of the grid
are shorter than it. Only consider trees in the same row or column; that is,
only look up, down, left, or right from any given tree.

All of the trees around the edge of the grid are visible - since they are
already on the edge, there are no trees to block the view. In this example,
that only leaves the interior nine trees to consider:

The top-left 5 is visible from the left and top. (It isn't visible from the
right or bottom since other trees of height 5 are in the way.)

The top-middle 5 is visible from the top and right.

The top-right 1 is not visible from any direction; for it to be visible, there
would need to only be trees of height 0 between it and an edge.

The left-middle 5 is visible, but only from the right.

The center 3 is not visible from any direction; for it to be visible, there
would need to be only trees of at most height 2 between it and an edge.

The right-middle 3 is visible from the right.

In the bottom row, the middle 5 is visible, but the 3 and 4 are not.

With 16 trees visible on the edge and another 5 visible in the interior, a
total of 21 trees are visible in this arrangement.

Consider your map; how many trees are visible from outside the grid?
*/

func parseMap(fileScanner *bufio.Scanner) [][]int {
	treeMap := [][]int{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		row := []int{}
		for _, char := range line {
			height, _ := strconv.Atoi(string(char))
			row = append(row, height)
		}
		treeMap = append(treeMap, row)
	}
	return treeMap
}

/*
30373
25512
65332
33549
35390
*/

// Day8Part1 ...
func Day8Part1(filepath string) any {
	var res int

	// open file
	readFile, _ := os.Open(filepath)

	// read line
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// parse map in to [][]int
	treeMap := parseMap(fileScanner)
	// fmt.Println(treeMap)

	// init visible trees 2D array
	visible := make([][]bool, len(treeMap))
	for i := range visible {
		visible[i] = make([]bool, len(treeMap[0]))
	}

	// look from left to r
	for i := range treeMap {
		max := 0
		for j := range treeMap[0] {
			if i == 0 || i == len(treeMap)-1 || j == 0 || j == len(treeMap[0])-1 || treeMap[i][j] > max {
				visible[i][j] = true
				max = treeMap[i][j]
			}
		}
	}

	// look from right to l
	for i := range treeMap {
		max := 0
		for j := len(treeMap[0]) - 1; j >= 0; j-- {
			if i == 0 || i == len(treeMap)-1 || j == 0 || j == len(treeMap[0])-1 || treeMap[i][j] > max {
				visible[i][j] = true
				max = treeMap[i][j]
			}
		}
	}

	// look from up to down
	for j := 0; j <= len(treeMap[0])-1; j++ {
		max := 0
		for i := 0; i <= len(treeMap)-1; i++ {
			if i == 0 || i == len(treeMap)-1 || j == 0 || j == len(treeMap[0])-1 || treeMap[i][j] > max {
				visible[i][j] = true
				max = treeMap[i][j]
			}
		}
	}

	// look from down to up
	for j := 0; j <= len(treeMap[0])-1; j++ {
		max := 0
		for i := len(treeMap) - 1; i >= 0; i-- {
			if i == 0 || i == len(treeMap)-1 || j == 0 || j == len(treeMap[0])-1 || treeMap[i][j] > max {
				visible[i][j] = true
				max = treeMap[i][j]
			}
		}
	}

	// traverse visible trees 2D array and count visibles
	for i := range visible {
		for j := range visible[i] {
			if visible[i][j] {
				res++
			}
		}
	}

	return res
}

/*
Content with the amount of tree cover available, the Elves just need to know
the best spot to build their tree house: they would like to be able to see a
lot of trees.

To measure the viewing distance from a given tree, look up, down, left, and
right from that tree; stop if you reach an edge or at the first tree that is
the same height or taller than the tree under consideration. (If a tree is
right on the edge, at least one of its viewing distances will be zero.)

The Elves don't care about distant trees taller than those found by the rules
above; the proposed tree house has large eaves to keep it dry, so they wouldn't
be able to see higher than the tree house anyway.

In the example above, consider the middle 5 in the second row:

30373
25512
65332
33549
35390

Looking up, its view is not blocked; it can see 1 tree (of height 3).

Looking left, its view is blocked immediately; it can see only 1 tree (of
height 5, right next to it).

Looking right, its view is not blocked; it can see 2 trees.

Looking down, its view is blocked eventually; it can see 2 trees (one of height
3, then the tree of height 5 that blocks its view).

A tree's scenic score is found by multiplying together its viewing distance in
each of the four directions. For this tree, this is 4 (found by multiplying 1 *
1 * 2 * 2).

However, you can do even better: consider the tree of height 5 in the middle of
the fourth row:

30373
25512
65332
33549
35390

Looking up, its view is blocked at 2 trees (by another tree with a height of
5).

Looking left, its view is not blocked; it can see 2 trees.

Looking down, its view is also not blocked; it can see 1 tree.

Looking right, its view is blocked at 2 trees (by a massive tree of height 9).

This tree's scenic score is 8 (2 * 2 * 1 * 2); this is the ideal spot for the
tree house.

Consider each tree on your map. What is the highest scenic score possible for
any tree?
*/

// Day8Part2 ...
func Day8Part2(filepath string) any {
	var res int

	// parse out treeMap from file
	// open file
	readFile, _ := os.Open(filepath)

	// read line
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// parse map in to [][]int
	treeMap := parseMap(fileScanner)

	// iterate every tree
	for i := range treeMap {
		for j := range treeMap[0] {
			// for every tree, calculate scenicScore
			s := getScenicScore(treeMap, i, j)
			// update max scenic score
			if s > res {
				res = s
			}
		}
	}
	return res
}

func getScenicScore(treeMap [][]int, r, c int) int {
	score := 1
	i, j := 0, 0
	tree := treeMap[r][c]
	// for each direction, count how many trees have height <= to current

	// iterate up
	j = c
	upScore := 0
	for i := r - 1; i >= 0; i-- {
		cur := treeMap[i][j]
		upScore++
		if cur >= tree {
			break
		}
	}
	score *= upScore

	// iterate down
	downScore := 0
	for i := r + 1; i <= len(treeMap)-1; i++ {
		cur := treeMap[i][j]
		downScore++
		if cur >= tree {
			break
		}
	}
	score *= downScore

	// iterate left
	i = r
	leftScore := 0
	for j := c - 1; j >= 0; j-- {
		cur := treeMap[i][j]
		leftScore++
		if cur >= tree {
			break
		}
	}
	score *= leftScore

	// iterate right
	rightScore := 0
	for j := c + 1; j <= len(treeMap[0])-1; j++ {
		cur := treeMap[i][j]
		rightScore++
		if cur >= tree {
			break
		}
	}
	score *= rightScore

	return score
}
