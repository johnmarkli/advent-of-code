package aoc2023

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	removeLens = '-'
	addLens    = '='
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

	file, scanner := readFile(filepath)
	defer file.Close()

	lbs := newLightBoxes(scanner)
	result = lbs.focusingPower()

	return result
}

type lightBoxes []*lightBox

func (lbs lightBoxes) String() string {
	var out string
	for i, lb := range lbs {
		if lb != nil && len(*lb) > 0 {
			out += fmt.Sprintf("\nBox %d: %s", i, lb.String())
		}
	}
	return out
}

func newLightBoxes(scanner *bufio.Scanner) lightBoxes {
	lbs := lightBoxes(make([]*lightBox, 256))
	for scanner.Scan() {
		line := scanner.Text()
		initSteps := strings.Split(line, ",")
		for _, initStep := range initSteps {
			var label string
			var op byte
			var focalLength int

			if initStep[len(initStep)-1] == removeLens { // removeLens
				label = initStep[:len(initStep)-1]
				op = removeLens
			} else { // addLens
				equalSplit := strings.Split(initStep, string(addLens))
				if len(equalSplit) == 2 {
					label = equalSplit[0]
					focalLength, _ = strconv.Atoi(equalSplit[1])
				}
				op = addLens
			}

			boxHash := InitStr(label).Hash()
			// fmt.Println("label, op, focalLength, boxHash", label, string(op), focalLength, boxHash)
			if lbs[boxHash] == nil {
				lbs[boxHash] = &lightBox{}
			}
			lb := lbs[boxHash]
			switch op {
			case addLens:
				lbs[boxHash] = lb.addLens(label, focalLength)
				// fmt.Println("addLens", lbs)
			case removeLens:
				lbs[boxHash] = lb.removeLens(label)
				// fmt.Println("removeLens", lbs)
			}

		}
	}
	return lbs
}

func (lbs lightBoxes) focusingPower() int {
	var result int
	for boxNum, lb := range lbs {
		if lb != nil {
			result += lb.focusingPower(boxNum)
		}
	}
	return result
}

type lightBox []lens

func (lb lightBox) String() string {
	var out string
	for _, l := range lb {
		out += fmt.Sprintf("[%s %d] ", l.label, l.focalLength)
	}
	return out
}

func (lb lightBox) focusingPower(boxNum int) int {
	var result int
	for slotNum, l := range lb {
		result += l.focusingPower(boxNum, slotNum+1)
	}
	return result
}

func (lb lightBox) addLens(label string, focalLength int) *lightBox {
	newLens := lens{label, focalLength}
	for i, l := range lb {
		if l.label == label {
			lb[i] = newLens // replace lens
			return &lb
		}
	}
	lb = append(lb, newLens) // add lens to back of box
	return &lb
}

func (lb lightBox) removeLens(label string) *lightBox {
	// find lens with label and remove it
	for i, l := range lb {
		if l.label == label {
			if i == 0 {
				lb = lb[1:] // if beginning, remove from beginning
				return &lb
			}
			if i == len(lb)-1 {
				lb = lb[0:i] // if end, remove from end
				return &lb
			}
			lb = append(lb[0:i], lb[i+1:]...) // if middle, remove from middle
			return &lb
		}
	}
	return &lb
}

type lens struct {
	label       string
	focalLength int
}

// The focusing power of a single lens is the result of multiplying together:
// One plus the box number of the lens in question.
// The slot number of the lens within the box: 1 for the first lens, 2 for the second lens, and so on.
// The focal length of the lens.
func (l lens) focusingPower(boxNum, slotNum int) int {
	result := (1 + boxNum) * slotNum * l.focalLength
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
