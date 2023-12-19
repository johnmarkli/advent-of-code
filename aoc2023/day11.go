package aoc2023

import (
	"bufio"
	"strconv"
)

const (
	defaultExpandFactor = 1_000_000
)

// Day11Part1 ...
func Day11Part1(filepath string) any {
	var result int

	file, fileScanner := readFile(filepath)
	defer file.Close()

	gm := newGalaxyMap(fileScanner)
	gm.expand(2)
	result = gm.shortestPaths()

	// fmt.Println(gm)

	return result
}

// Day11Part2 ...
func Day11Part2(filepath string) any {
	return Day11Part2Solve(filepath, defaultExpandFactor)
}

// Day11Part2Solve ...
func Day11Part2Solve(filepath string, factor ...int) any {
	var result int

	file, fileScanner := readFile(filepath)
	defer file.Close()

	gm := newGalaxyMap(fileScanner)
	gm.expand(factor...)
	result = gm.shortestPaths()

	// fmt.Println(gm)

	return result
}

type galaxyMap struct {
	originalMap  [][]byte
	expandedMap  [][]byte
	galaxyCoords []*coord
}

func (gm *galaxyMap) String() string {
	out := "originalMap:\n"
	for i := range gm.originalMap {
		var rowOut string
		for j := range gm.originalMap[i] {
			rowOut += string(gm.originalMap[i][j])
		}
		out += rowOut + "\n"
	}

	out += "\n"

	galaxyCount := 1
	out += "expandedMap:\n"
	for i := range gm.expandedMap {
		var rowOut string
		for j := range gm.expandedMap[i] {
			c := string(gm.expandedMap[i][j])
			if gm.expandedMap[i][j] == '#' {
				c = string(strconv.Itoa(galaxyCount)[0])
				galaxyCount++
			}
			rowOut += c
		}
		out += rowOut + "\n"
	}

	return out
}

func (gm *galaxyMap) expand(factor ...int) {
	fac := defaultExpandFactor - 1
	if len(factor) > 0 {
		fac = factor[0] - 1
	}
	// loop through originalMap and keep track of num galaxies found in each row and each column
	rowGalaxies := map[int]int{}
	for r := range gm.originalMap {
		rowGalaxies[r] = 0
	}
	colGalaxies := map[int]int{}
	for c := range gm.originalMap[0] {
		colGalaxies[c] = 0
	}

	for r := range gm.originalMap {
		for c := range gm.originalMap[r] {
			el := gm.originalMap[r][c]
			if el == '#' {
				rowGalaxies[r]++
				colGalaxies[c]++
			}
		}
	}

	// Builds actual expanded map - commented out for part 2 since it's too big
	// // get num of extra rows and cols
	// extraRows := 0
	// for _, count := range rowGalaxies {
	// 	if count == 0 {
	// 		// extraRows++
	// 		extraRows += fac
	// 	}
	// }
	// extraCols := 0
	// for _, count := range colGalaxies {
	// 	if count == 0 {
	// 		// extraCols++
	// 		extraCols += fac
	// 	}
	// }
	//
	// // initialize expanded map with extra rows and cols
	// expandedMap := make([][]byte, len(gm.originalMap)+extraRows)
	// for i := range expandedMap {
	// 	expandedMap[i] = make([]byte, len(gm.originalMap[0])+extraCols)
	// 	for j := range expandedMap[0] {
	// 		expandedMap[i][j] = '*'
	// 	}
	// }

	// loop originalMap
	// if col found in colGalaxies, add another empty col
	// if row found in rowGalaxies, add another empty row

	galaxyCoords := []*coord{}
	i, j, rowOffset, colOffset := 0, 0, 0, 0
	for i < len(gm.originalMap) {
		if val, ok := rowGalaxies[i]; ok && val == 0 {
			// for k := range expandedMap[0] {
			// 	expandedMap[i+rowOffset][k] = '.'
			// }
			l := 1
			for l <= fac {
				rowOffset++
				// for k := range expandedMap[0] {
				// 	expandedMap[i+rowOffset][k] = '.'
				// }
				l++
			}
		} else {
			for j < len(gm.originalMap[0]) {
				if val, ok := colGalaxies[j]; ok && val == 0 {
					// expandedMap[i+rowOffset][j+colOffset] = '.'
					l := 1
					for l <= fac {
						colOffset++
						// expandedMap[i+rowOffset][j+colOffset] = '.'
						l++
					}
				} else {
					el := gm.originalMap[i][j]
					// expandedMap[i+rowOffset][j+colOffset] = el
					if el == '#' {
						c := &coord{x: i + rowOffset, y: j + colOffset}
						galaxyCoords = append(galaxyCoords, c)
					}
				}
				j++
			}
		}
		j, colOffset = 0, 0
		i++
	}
	// gm.expandedMap = expandedMap
	gm.galaxyCoords = galaxyCoords
}

func (gm *galaxyMap) shortestPaths() int {
	var result int
	// dist b/w 0,4 and 1,9 is 6
	// rise + run => 1 + 5 = 6
	// dist b/w 6,1 and 11,5 is 9
	// rise + run => 11-6 + 5-1 = 5 + 4 = 9

	for i, g1 := range gm.galaxyCoords {
		for _, g2 := range gm.galaxyCoords[i+1:] {
			dist := g1.dist(g2)
			result += dist
		}
	}
	return result
}

func newGalaxyMap(fileScanner *bufio.Scanner) *galaxyMap {
	originalMap := [][]byte{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		originalMap = append(originalMap, []byte(line))
	}
	return &galaxyMap{
		originalMap: originalMap,
	}
}
