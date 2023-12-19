package aoc2023

import (
	"bufio"
	"fmt"
	"strconv"
)

// Day10Part1 ...
func Day10Part1(filepath string) any {
	var result int

	file, fileScanner := readFile(filepath)
	defer file.Close()

	pipeMap := NewPipeMap(fileScanner)

	// loop through pipemap from start following a connected path
	// count steps until reach start again
	steps := pipeMap.traverse()
	pipeMap.showPipes()
	// result is count / 2
	result = steps / 2

	return result
}

// Day10Part2 ...
func Day10Part2(filepath string) any {
	var result int

	file, fileScanner := readFile(filepath)
	defer file.Close()

	pipeMap := NewPipeMap(fileScanner)
	_ = pipeMap.traverse()
	// pipeMap.showPipes()

	result = pipeMap.getTilesInLoop()
	//
	pipeMap.showPipeLoop()
	// fmt.Println()
	// pipeMap.showEnclosed()

	return result
}

// PipeMap ...
type PipeMap struct {
	tiles    [][]byte
	start    *coord
	pipes    [][]int
	enclosed [][]int
}

const (
	startPoint = 'S'
	vertical   = '|'
	horizontal = '-'
	upRight    = 'L'
	upLeft     = 'J'
	downLeft   = '7'
	downRight  = 'F'
)

func (pm *PipeMap) showPipeLoop() {
	outRows := []string{}
	for i := range pm.tiles {
		var outRow string
		for j := range pm.tiles[i] {
			if pm.pipes[i][j] == 1 {
				pipe := pm.tiles[i][j]
				var out string
				switch pipe {
				case startPoint:
					out = "★"
				case vertical:
					out = "┃"
				case horizontal:
					out = "━"
				case upRight:
					out = "┗"
				case upLeft:
					out = "┛"
				case downLeft:
					out = "┓"
				case downRight:
					out = "┏"
				default:
					out = "#"
				}
				outRow += out
			} else {
				// outRow += "·"
				if pm.enclosed[i][j] == 1 {
					outRow += "I"
				} else {
					outRow += "O"
				}
			}
		}
		outRows = append(outRows, outRow)
	}

	for _, row := range outRows {
		fmt.Println(row)
	}
}

func (pm *PipeMap) showPipes() {
	outRows := []string{}
	for i := range pm.pipes {
		var outRow string
		for j := range pm.pipes[i] {
			outRow += strconv.Itoa(pm.pipes[i][j])
		}
		outRows = append(outRows, outRow)
	}

	for _, row := range outRows {
		fmt.Println(row)
	}
}

func (pm *PipeMap) showEnclosed() {
	outRows := []string{}
	for i := range pm.enclosed {
		var outRow string
		for j := range pm.enclosed[i] {
			outRow += strconv.Itoa(pm.enclosed[i][j])
		}
		outRows = append(outRows, outRow)
	}

	for _, row := range outRows {
		fmt.Println(row)
	}
}

func (pm *PipeMap) traverse() int {
	steps := 0
	curPos := pm.start
	pm.pipes[curPos.x][curPos.y] = 1
	var lastPos *coord

	// find connections
	conns, startPipe := pm.findConnectionsAt(curPos)
	// pick connection
	conn := conns[0]
	// move
	lastPos = curPos
	curPos = conn
	steps++

	// find connections
	for !curPos.sameAs(pm.start) {
		conns, _ := pm.findConnectionsAt(curPos)

		if len(conns) > 1 {
			// pick connection
			var conn *coord
			for _, c := range conns {
				if !c.sameAs(lastPos) {
					conn = c
				}
			}
			// move
			pm.pipes[curPos.x][curPos.y] = 1
			lastPos = curPos
			curPos = conn
			steps++
		} else {
			break
		}
	}

	// replace start with pipe
	pm.tiles[pm.start.x][pm.start.y] = startPipe

	return steps
}

// use point in polygon / even-odd algorithm
// scan each line and keep track of intersections
// if point is found when intersections is odd, then it is inside
// if point is found when intersections is even, then it is outside
func (pm *PipeMap) getTilesInLoop() int {
	var result int
	for r := range pm.tiles {
		enclosed := false
		var prevCorner byte
		for c := range pm.tiles[r] {
			// if corner to corner makes a U shape, region stays the same
			// if corner to corner doesn't make a U shape, it is a wall and regions switch
			tile := pm.tiles[r][c]
			curTileIsPipe := pm.pipes[r][c] == 1
			curTileIsCorner := (tile == upRight || tile == downRight || tile == downLeft || tile == upLeft)
			isWall := false
			if curTileIsPipe {
				isUShape := curTileIsCorner && ((tile == downLeft && prevCorner == downRight) || (tile == upLeft && prevCorner == upRight))
				isWall = (tile == vertical) || (curTileIsCorner && prevCorner != byte(0) && !isUShape)
				if curTileIsCorner && !isWall {
					if isUShape {
						prevCorner = byte(0)
					} else {
						prevCorner = tile
					}
				}
			}

			// if isWall, switch
			if isWall {
				enclosed = !enclosed
				pm.enclosed[r][c] = 2 // a switch occured
				prevCorner = byte(0)  // reset corner
			}

			if !curTileIsPipe && enclosed {
				pm.enclosed[r][c] = 1 // set enclosed
				result++
			}
		}
	}
	return result
}

func (pm *PipeMap) areConnected(pos1 *coord, pos2 *coord) bool {
	if pos1 == nil || pos2 == nil {
		return false
	}
	conns, _ := pm.findConnectionsAt(pos1)
	for _, conn := range conns {
		if conn.sameAs(pos2) {
			return true
		}
	}
	return false
}

// returns 2 coordinates that a position is connected to
func (pm *PipeMap) findConnectionsAt(pos *coord) ([]*coord, byte) {
	var up, right, left, down *coord

	result := []*coord{}

	switch pm.tiles[pos.x][pos.y] {
	case startPoint:
		up = pos.up()
		down = pos.down()
		right = pos.right()
		left = pos.left()
	case vertical:
		up = pos.up()
		down = pos.down()
	case horizontal:
		right = pos.right()
		left = pos.left()
	case upRight:
		up = pos.up()
		right = pos.right()
	case upLeft:
		up = pos.up()
		left = pos.left()
	case downLeft:
		down = pos.down()
		left = pos.left()
	case downRight:
		down = pos.down()
		right = pos.right()
	}

	var hasUp, hasDown, hasRight, hasLeft bool
	// check up for vertical or downLeft or downRight
	if up != nil && up.x >= 0 {
		el := pm.tiles[up.x][up.y]
		if el == vertical || el == downLeft || el == downRight || el == startPoint {
			result = append(result, up)
			hasUp = true
		}
	}
	// check right for horizontal or downLeft or upLeft
	if right != nil && right.y < len(pm.tiles[0]) {
		el := pm.tiles[right.x][right.y]
		if el == horizontal || el == downLeft || el == upLeft || el == startPoint {
			result = append(result, right)
			hasRight = true
		}
	}
	// check down for vertical or upRight or upLeft
	if down != nil && down.x < len(pm.tiles) {
		el := pm.tiles[down.x][down.y]
		if el == vertical || el == upRight || el == upLeft || el == startPoint {
			result = append(result, down)
			hasDown = true
		}
	}
	// check left for horizontal or downRight or upRight
	if left != nil && left.y >= 0 {
		el := pm.tiles[left.x][left.y]
		if el == horizontal || el == downRight || el == upRight || el == startPoint {
			result = append(result, left)
			hasLeft = true
		}
	}
	var pipe byte
	switch {
	case hasUp && hasLeft:
		pipe = upLeft
	case hasUp && hasRight:
		pipe = upRight
	case hasDown && hasRight:
		pipe = downRight
	case hasDown && hasLeft:
		pipe = downLeft
	case hasUp && hasDown:
		pipe = vertical
	case hasRight && hasLeft:
		pipe = horizontal
	}
	return result, pipe
}

// NewPipeMap ...
func NewPipeMap(fileScanner *bufio.Scanner) *PipeMap {
	tiles := [][]byte{}
	var start *coord
	row := 0
	for fileScanner.Scan() {
		tileRow := []byte{}
		line := fileScanner.Bytes()
		for i, b := range line {
			if b == 'S' {
				start = &coord{
					x: row,
					y: i,
				}
			}
			tileRow = append(tileRow, b)
		}
		tiles = append(tiles, tileRow)
		row++
	}

	pipes := make([][]int, len(tiles))
	for i := range pipes {
		pipes[i] = make([]int, len(tiles[0]))
	}

	enclosed := make([][]int, len(tiles))
	for i := range enclosed {
		enclosed[i] = make([]int, len(tiles[0]))
	}

	return &PipeMap{
		tiles:    tiles,
		start:    start,
		pipes:    pipes,
		enclosed: enclosed,
	}
}
