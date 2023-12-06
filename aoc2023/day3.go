package aoc2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
You and the Elf eventually reach a gondola lift station; he says the gondola
lift will take you up to the water source, but this is as far as he can bring
you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem:
they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of
surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working
right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine,
but nobody can figure out which one. If you can add up all the part numbers in
the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of
the engine. There are lots of numbers and symbols you don't really understand,
but apparently any number adjacent to a symbol, even diagonally, is a "part
number" and should be included in your sum. (Periods (.) do not count as a
symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

In this schematic, two numbers are not part numbers because they are not
adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number
is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all
of the part numbers in the engine schematic?
*/

// Day3Part1 ...
func Day3Part1(filepath string) any {
	result := 0

	readFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	schematic := [][]byte{}
	// convert input to 2D array
	for fileScanner.Scan() {
		line := fileScanner.Bytes()
		row := []byte{}
		for _, c := range line {
			row = append(row, c)
		}
		schematic = append(schematic, row)
	}

	if err := fileScanner.Err(); err != nil {
		panic(err)
	}

	// run through 2D array to find numbers
	curNum := []byte{}
	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			c := schematic[i][j]
			// if number is found, look for a symbol adjacent, if symbol is found, add to result
			isNumber := false
			if c >= '0' && c <= '9' {
				curNum = append(curNum, c)
				isNumber = true
			}
			if !isNumber || j == len(schematic[i])-1 {
				endNum := j
				if !isNumber {
					endNum = j - 1
				}
				if len(curNum) > 0 {
					sym, _, _ := isPartNumber(schematic, curNum, i, endNum)
					if sym != byte(0) {
						num, err := strconv.Atoi(string(curNum))
						if err != nil {
							panic(err)
						}
						result += num
					}
					curNum = []byte{}
				}
			}
		}
	}
	return result
}

/*
The engineer finds the missing part and installs it in the engine! As the
engine springs to life, you jump in the closest gondola, finally ready to
ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong?
Fortunately, the gondola has a phone labeled "help", so you pick it up and the
engineer answers.

Before you can explain the situation, she suggests that you look out the
window. There stands the engineer, holding a phone in one hand and waving with
the other. You're going so slowly that you haven't even left the station. You
exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is
wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its
gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so
that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

In this schematic, there are two gears. The first is in the top left; it has
part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the
lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear
because it is only adjacent to one part number.) Adding up all of the gear
ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?
*/

// Day3Part2 ...
func Day3Part2(filepath string) any {
	result := 0
	readFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	schematic := [][]byte{}
	// convert input to 2D array
	for fileScanner.Scan() {
		line := fileScanner.Bytes()
		row := []byte{}
		for _, c := range line {
			row = append(row, c)
		}
		schematic = append(schematic, row)
	}

	if err := fileScanner.Err(); err != nil {
		panic(err)
	}

	gearMap := map[string][]int{}

	// run through 2D array to find numbers
	curNum := []byte{}
	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			// if has * symbol adjacent, add star pos to map with number
			c := schematic[i][j]
			// if number is found, look for a symbol adjacent, if symbol is found, add to result
			isNumber := false
			if c >= '0' && c <= '9' {
				curNum = append(curNum, c)
				isNumber = true
			}

			if !isNumber || j == len(schematic[i])-1 {
				endNum := j
				if !isNumber {
					endNum = j - 1
				}
				if len(curNum) > 0 {
					sym, x, y := isPartNumber(schematic, curNum, i, endNum)
					if sym == '*' {
						num, err := strconv.Atoi(string(curNum))
						if err != nil {
							panic(err)
						}
						pos := fmt.Sprintf("%d,%d", x, y)
						if _, ok := gearMap[pos]; !ok {
							gearMap[pos] = []int{}
						}
						gearMap[pos] = append(gearMap[pos], num)
					}
					curNum = []byte{}
				}
			}
		}
	}

	// loop all star positions and if array of numbers is len 2, multiply and add to result
	for _, gearNums := range gearMap {
		if len(gearNums) == 2 {
			gearRatio := gearNums[0] * gearNums[1]
			result += gearRatio
		}
	}

	return result
}

func isPartNumber(schematic [][]byte, num []byte, x, y int) (byte, int, int) {
	var result byte
	startY := y - len(num) + 1
	endY := y
	checkStartY := startY - 1
	checkEndY := endY + 1
	aboveX := x - 1
	belowX := x + 1
	// check above
	if aboveX >= 0 {
		for k := checkStartY; k <= checkEndY; k++ {
			if k >= 0 && k < len(schematic[aboveX]) && isSymbol(schematic[aboveX][k]) {
				return schematic[aboveX][k], aboveX, k
			}
		}
	}
	// check to the sides
	if checkStartY >= 0 && isSymbol(schematic[x][checkStartY]) {
		return schematic[x][checkStartY], x, checkStartY
	}
	if checkEndY < len(schematic[x]) && isSymbol(schematic[x][checkEndY]) {
		return schematic[x][checkEndY], x, checkEndY
	}
	// check below
	if belowX < len(schematic) {
		for k := checkStartY; k <= checkEndY; k++ {
			if k >= 0 && k < len(schematic[belowX]) && isSymbol(schematic[belowX][k]) {
				return schematic[belowX][k], belowX, k
			}
		}
	}
	return result, 0, 0
}

func isSymbol(c byte) bool {
	if (c >= '0' && c <= '9') || c == '.' {
		return false
	}
	return true
}
