package aoc2023

import (
	"bufio"
)

const (
	lavaEmptySpace   = '.'
	lavaMirrorFSlash = '/'
	lavaMirrorBSlash = '\\'
	lavaSplitterUD   = '|'
	lavaSplitterRL   = '-'
	lavaEnergized    = '#'
	lavaDirUp        = 'U'
	lavaDirDown      = 'D'
	lavaDirRight     = 'R'
	lavaDirLeft      = 'L'
)

// Day16Part1 ...
func Day16Part1(filepath string) any {
	var result int

	file, scanner := readFile(filepath)
	defer file.Close()

	lts := newLavaTiles(scanner)

	// energize lava tiles
	ets := lts.energize(coord{0, -1})

	// count energized tiles
	result = ets.energized()

	return result
}

// Day16Part2 ...
func Day16Part2(filepath string) any {
	var result int

	file, scanner := readFile(filepath)
	defer file.Close()

	lts := newLavaTiles(scanner)

	// energize lava tiles from all sides and calculate max energized
	rows := []int{-1, len(lts)}
	cols := []int{-1, len(lts)}

	for _, i := range rows {
		for j := range lts[0] {
			ets := lts.energize(coord{i, j})

			// count energized tiles
			count := ets.energized()
			if count > result {
				result = count
			}
		}
	}

	for i := range lts {
		for _, j := range cols {
			ets := lts.energize(coord{i, j})

			// count energized tiles
			count := ets.energized()
			if count > result {
				result = count
			}
		}
	}

	return result
}

type lavaTiles [][]byte

func (lts lavaTiles) maxX() int {
	return len(lts) - 1
}

func (lts lavaTiles) maxY() int {
	if len(lts) > 0 {
		return len(lts[0]) - 1
	}
	return 0
}

func (lts lavaTiles) String() string {
	var out string
	for _, lt := range lts {
		out += string(lt) + "\n"
	}
	return out
}

var beamCache map[string]map[byte]bool

func (lts lavaTiles) energize(startPos coord) lavaTiles {
	// initialize energized tiles
	energized := make([][]byte, len(lts))
	for i := range lts {
		energized[i] = make([]byte, len(lts[i]))
		for j := range lts[i] {
			energized[i][j] = lavaEmptySpace
		}
	}

	// beam starts at top-left going left
	// empty space - beam keeps going
	// mirrors / and \ - beam is reflected at 90 degrees
	// splitters - pass through, or split in to 2 perpendicular to direction

	// DFS to follow path through tiles and handle splits
	// dfs func needs cur pos, direction, lts, energized - return energized

	var dir byte
	switch {
	case startPos.x < 0:
		dir = 'D'
	case startPos.x >= len(lts):
		dir = 'U'
	case startPos.y < 0:
		dir = 'R'
	case startPos.y >= len(lts[0]):
		dir = 'R'
	}

	beamCache = map[string]map[byte]bool{}
	energized = lavaDFS(startPos, dir, lts, energized)

	return lavaTiles(energized)
}

func lavaDFS(curPos coord, dir byte, lts, energized lavaTiles) lavaTiles {
	// get next pos
	switch dir {
	case lavaDirUp:
		curPos.x--
	case lavaDirDown:
		curPos.x++
	case lavaDirRight:
		curPos.y++
	case lavaDirLeft:
		curPos.y--
	}

	// base case
	// stop if out of bounds or if a beam already passed through it in that direction
	if curPos.outOfBounds(lts.maxX(), lts.maxY()) {
		return energized
	}

	if beamCache[curPos.String()] == nil {
		beamCache[curPos.String()] = map[byte]bool{}
	}

	newTile := lts[curPos.x][curPos.y]

	// initialize direction symbol
	dirSym := newTile
	if dirSym == lavaEmptySpace {
		switch dir {
		case lavaDirUp:
			dirSym = '^'
		case lavaDirDown:
			dirSym = 'v'
		case lavaDirLeft:
			dirSym = '<'
		case lavaDirRight:
			dirSym = '>'
		}
		if _, ok := beamCache[curPos.String()][dirSym]; ok {
			return energized
		}
	}

	beamCache[curPos.String()][dirSym] = true

	energized[curPos.x][curPos.y] = dirSym

	// recursive case
	switch newTile {
	case lavaEmptySpace:
		energized = lavaDFS(curPos, dir, lts, energized)
	case lavaSplitterUD:
		switch dir { // "|" splitter
		case lavaDirRight, lavaDirLeft:
			energized = lavaDFS(curPos, 'U', lts, energized)
			energized = lavaDFS(curPos, 'D', lts, energized)
		case lavaDirUp, lavaDirDown:
			energized = lavaDFS(curPos, dir, lts, energized)
		}
	case lavaSplitterRL: // "-" splitter
		switch dir {
		case lavaDirUp, lavaDirDown:
			energized = lavaDFS(curPos, 'R', lts, energized)
			energized = lavaDFS(curPos, 'L', lts, energized)
		case lavaDirRight, lavaDirLeft:
			energized = lavaDFS(curPos, dir, lts, energized)
		}
	case lavaMirrorFSlash: // "/" mirror
		switch dir {
		case lavaDirUp:
			energized = lavaDFS(curPos, 'R', lts, energized)
		case lavaDirDown:
			energized = lavaDFS(curPos, 'L', lts, energized)
		case lavaDirRight:
			energized = lavaDFS(curPos, 'U', lts, energized)
		case lavaDirLeft:
			energized = lavaDFS(curPos, 'D', lts, energized)
		}
	case lavaMirrorBSlash: // "\" mirror
		switch dir {
		case lavaDirUp:
			energized = lavaDFS(curPos, 'L', lts, energized)
		case lavaDirDown:
			energized = lavaDFS(curPos, 'R', lts, energized)
		case lavaDirRight:
			energized = lavaDFS(curPos, 'D', lts, energized)
		case lavaDirLeft:
			energized = lavaDFS(curPos, 'U', lts, energized)
		}
	}

	return energized
}

func (lts lavaTiles) energized() int {
	var result int
	for i := range lts {
		for j := range lts[i] {
			if lts[i][j] != lavaEmptySpace {
				result++
			}
		}
	}
	return result
}

func newLavaTiles(scanner *bufio.Scanner) lavaTiles {
	var lts lavaTiles
	for scanner.Scan() {
		line := scanner.Text()
		lts = append(lts, []byte(line))
	}
	return lts

}
