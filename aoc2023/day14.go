package aoc2023

import (
	"bufio"
)

const (
	roundRock = 'O'
	cubeRock  = '#'
	noRock    = '.'
)

// Day14Part1 ...
func Day14Part1(filepath string) any {
	var result int

	file, scanner := readFile(filepath)
	defer file.Close()

	rd := newReflectorDish(scanner)
	rd.tiltNorth()
	result = rd.loadNorth()

	return result
}

// Day14Part2 ...
func Day14Part2(filepath string) any {
	var result int

	file, scanner := readFile(filepath)
	defer file.Close()

	rd := newReflectorDish(scanner)
	rd.cycle(1_000_000_000)
	result = rd.loadNorth()

	return result
}

// ReflectorDish ...
type ReflectorDish [][]byte

func (rd ReflectorDish) String() string {
	var out string
	for _, row := range rd {
		out += string(row) + "\n"
	}
	return out
}

func (rd ReflectorDish) cycle(num int) {
	seen := map[string]int{}
	i := 1
	var matchingCycle int
	for i <= num {
		// tilt north
		rd.tiltNorth()
		rd.tiltWest()
		rd.tiltSouth()
		rd.tiltEast()
		if j, ok := seen[rd.String()]; ok {
			matchingCycle = j
			break
		} else {
			seen[rd.String()] = i
		}
		i++
	}
	diff := i - matchingCycle
	remainingToGetToNum := num - i
	multiple := remainingToGetToNum / diff
	closestMultiple := multiple * diff
	remaining := remainingToGetToNum - closestMultiple

	for i := 1; i <= remaining; i++ {
		rd.tiltNorth()
		rd.tiltWest()
		rd.tiltSouth()
		rd.tiltEast()
	}

}

func (rd ReflectorDish) tiltNorth() {
	// for each col
	for c := 0; c < len(rd[0]); c++ {
		// for each row
		for r := 0; r < len(rd); r++ {
			el := rd[r][c]
			// if its a round rock and can't move any more north, move it north as much as possible
			if el == roundRock && r > 0 && rd[r-1][c] != roundRock {
				// init pointer to one above
				i := r - 1
				// check if its a round rock
				swapped := false
				for !swapped && i >= -1 {
					var newEl byte
					if i > -1 {
						newEl = rd[i][c]
					}
					// if its a round rock or cube rock, swap round rock found with down one from cur pos
					if newEl == roundRock || newEl == cubeRock || i == -1 {
						rd[r][c], rd[i+1][c] = rd[i+1][c], rd[r][c]
						swapped = true
					} else { // if not round rock, move up
						i--
					}
				}
			}
		}
	}
}

func (rd ReflectorDish) tiltSouth() {
	// for each col
	for c := 0; c < len(rd[0]); c++ {
		// for each row
		for r := len(rd) - 1; r >= 0; r-- {
			el := rd[r][c]
			// if its a round rock and can't move any more south, move it south as much as possible
			if el == roundRock && r < len(rd)-1 && rd[r+1][c] != roundRock {
				// init pointer to one above
				i := r + 1
				// check if its a round rock
				swapped := false
				for !swapped && i <= len(rd) {
					var newEl byte
					if i < len(rd) {
						newEl = rd[i][c]
					}
					// if its a round rock or cube rock, swap round rock found with up one from cur pos
					if newEl == roundRock || newEl == cubeRock || i == len(rd) {
						rd[r][c], rd[i-1][c] = rd[i-1][c], rd[r][c]
						swapped = true
					} else { // if not round rock, move down
						i++
					}
				}
			}
		}
	}
}

func (rd ReflectorDish) tiltWest() {
	// for each row
	for r := 0; r < len(rd); r++ {
		// for each col
		for c := 0; c < len(rd[0]); c++ {
			el := rd[r][c]
			// if its a round rock and can't move any more west, move it west as much as possible
			if el == roundRock && c > 0 && rd[r][c-1] != roundRock {
				// init pointer to one to the left
				i := c - 1
				// check if its a round rock
				swapped := false
				for !swapped && i >= -1 {
					var newEl byte
					if i > -1 {
						newEl = rd[r][i]
					}
					// if its a round rock or cube rock, swap round rock found with one to the right from cur pos
					if newEl == roundRock || newEl == cubeRock || i == -1 {
						rd[r][c], rd[r][i+1] = rd[r][i+1], rd[r][c]
						swapped = true
					} else { // if not round rock, move up
						i--
					}
				}
			}
		}
	}
}

func (rd ReflectorDish) tiltEast() {
	// for each row
	for r := 0; r < len(rd); r++ {
		// for each col
		for c := len(rd[0]) - 1; c >= 0; c-- {
			el := rd[r][c]
			// if its a round rock and can't move any more east, move it east as much as possible
			if el == roundRock && c < len(rd[0])-1 && rd[r][c+1] != roundRock {
				// init pointer to one above
				i := c + 1
				// check if its a round rock
				swapped := false
				for !swapped && i <= len(rd[0]) {
					var newEl byte
					if i < len(rd[0]) {
						newEl = rd[r][i]
					}
					// if its a round rock or cube rock, swap round rock found with one to the left from cur pos
					if newEl == roundRock || newEl == cubeRock || i == len(rd[0]) {
						rd[r][c], rd[r][i-1] = rd[r][i-1], rd[r][c]
						swapped = true
					} else { // if not round rock, move up
						i++
					}
				}
			}
		}
	}
}

func (rd ReflectorDish) loadNorth() int {
	var result int
	// for each row
	for r := range rd {
		// load is len of rd - row index
		load := len(rd) - r
		// for each col
		for c := range rd[r] {
			// if el is round rock, add load to result
			el := rd[r][c]
			if el == roundRock {
				result += load
			}
		}
	}
	return result
}

func newReflectorDish(scanner *bufio.Scanner) ReflectorDish {
	var dish [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		dish = append(dish, []byte(line))
	}
	return dish
}
