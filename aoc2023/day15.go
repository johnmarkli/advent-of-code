package aoc2023

import (
	"bufio"
	"strings"
)

// Day15Part1 ...
func Day15Part1(filepath string) any {
	var result int

	file, scanner := readFile(filepath)
	defer file.Close()

	seq := newInitSequence(scanner)
	result = seq.hashSum()

	return result
}

// Day15Part2 ...
func Day15Part2(filepath string) any {
	var result int

	// file, scanner := readFile(filepath)
	// defer file.Close()

	return result
}

type initSequence []InitStr

func (seq initSequence) hashSum() int {
	var result int
	for _, is := range seq {
		result += is.Hash()
	}
	return result
}

// InitStr ..
type InitStr string

// Hash ...
// Determine the ASCII code for the current character of the string.
// Increase the current value by the ASCII code you just determined.
// Set the current value to itself multiplied by 17.
// Set the current value to the remainder of dividing itself by 256.
func (is InitStr) Hash() int {
	var val int
	for i := 0; i < len(is); i++ {
		asciiVal := int(is[i])
		val += asciiVal
		val = val * 17
		rem := val % 256
		val = rem
	}
	return val
}

func newInitSequence(scanner *bufio.Scanner) initSequence {
	var seq initSequence
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		for _, el := range split {
			seq = append(seq, InitStr(el))
		}
	}
	return seq
}
