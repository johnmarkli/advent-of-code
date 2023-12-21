package aoc2023

import (
	"bufio"
)

// Day13Part1 ...
func Day13Part1(filepath string) any {
	var result int

	file, scanner := readFile(filepath)
	defer file.Close()

	rms := newReflMaps(scanner)
	// fmt.Println(rms)

	var res int
	for _, rm := range rms {
		res = rm.verticalReflCols()
		result += res
		res = rm.horizontalReflRows() * 100
		result += res
	}

	return result
}

// Day13Part2 ...
func Day13Part2(filepath string) any {
	var result int

	file, scanner := readFile(filepath)
	defer file.Close()

	rms := newReflMaps(scanner)

	// for each rm
	for _, rm := range rms {
		// come up with altered ReflMap first
		// then check vert and horiz, if only one has relfection, that is the result
		// if a reflection was present without fixing the smudge, it should be ignored. Following this, there is a unique possible line of reflection with other rules specified.
		i, j := 0, 0
		index := 0
		originalVert := rm.verticalReflCols()
		originalHoriz := rm.horizontalReflRows()
		// loop through each char in map, flip and create new ReflMap, then check index
		// originalIndex := originalVert + (originalHoriz * 100)
		for index == 0 && i < len(rm.rows) {
			for index == 0 && j < len(rm.rows[i]) {
				newRm := rm.flipCharAt(i, j)
				vert := newRm.verticalReflCols(originalVert)
				horiz := newRm.horizontalReflRows(originalHoriz)
				if horiz != 0 {
					index = horiz * 100
				} else if vert != 0 {
					index = vert
				}
				j++
			}
			j = 0
			i++
		}
		// fmt.Println("adding index", num, index)
		result += index
	}

	return result
}

type reflMaps []*ReflMap

func (rms reflMaps) String() string {
	var out string
	for _, rm := range rms {
		out += rm.String() + "\n"
	}
	return out
}

// ReflMap ...
type ReflMap struct {
	rows []string
	cols []string
}

func (rm *ReflMap) String() string {
	var out string
	out += "rows\n"
	for _, r := range rm.rows {
		out += r + "\n"
	}
	out += "\n"
	out += "cols\n"
	for _, r := range rm.cols {
		out += r + "\n"
	}
	return out
}

func (rm *ReflMap) flipCharAt(i, j int) *ReflMap {
	rowsCopy := make([]string, len(rm.rows))
	copy(rowsCopy, rm.rows)
	// alter char in row copy
	var newStr []byte
	var newChar byte
	if rowsCopy[i][j] == '#' {
		newChar = '.'
	} else {
		newChar = '#'
	}
	newStr = []byte(rowsCopy[i])
	newStr[j] = newChar
	rowsCopy[i] = string(newStr)

	colsCopy := make([]string, len(rowsCopy[0]))
	for _, str := range rowsCopy {
		for i, c := range str {
			colsCopy[i] += string(c)
		}
	}

	newRm := &ReflMap{
		rows: rowsCopy,
		cols: colsCopy,
	}
	return newRm
}

func (rm *ReflMap) horizontalReflRows(ignore ...int) int {
	result := rm.findReflectionIndex(rm.rows, ignore...)
	return result
}

func (rm *ReflMap) verticalReflCols(ignore ...int) int {
	result := rm.findReflectionIndex(rm.cols, ignore...)
	return result
}

func (rm *ReflMap) findReflectionIndex(patterns []string, ignore ...int) int {
	var result int
	ig := -1
	if len(ignore) > 0 {
		ig = ignore[0]
	}
	// say line exists before the row you're looking at
	// for each row, compare row above - if same, incr below and  decr above and
	// compare until reach end of map

	for i := 1; i < len(patterns); i++ {
		current := patterns[i]
		above := patterns[i-1]
		if current == above {
			if (i == len(patterns)-1 || i == 1) && i != ig {
				return i
			}
			j := i - 2
			k := i + 1
			match := true
			for match && j >= 0 && k < len(patterns) {
				if patterns[j] == patterns[k] {
					if (j == 0 || k == len(patterns)-1) && i != ig {
						return i
					}
					j--
					k++
				} else {
					match = false
				}
			}
		}
	}

	return result
}

func newReflMaps(scanner *bufio.Scanner) reflMaps {
	var rms reflMaps
	var rows []string
	var cols []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // if empty line, reset for new refl map
			rm := &ReflMap{
				rows: rows,
				cols: cols,
			}
			rms = append(rms, rm)
			rows = []string{}
			cols = []string{}
		} else {
			rows = append(rows, line)
			if len(cols) == 0 {
				cols = make([]string, len(line))
			}
			// add each character in line to its col
			for i, c := range line {
				cols[i] += string(c)
			}
		}

	}
	if len(rows) > 0 && len(cols) > 0 {
		rm := &ReflMap{
			rows: rows,
			cols: cols,
		}
		rms = append(rms, rm)
	}
	return rms
}
