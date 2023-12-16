package aoc2023

import (
	"bufio"
	"strconv"
	"strings"
)

// Day9Part1 ...
func Day9Part1(filepath string) any {
	var result int

	file, fileScanner := readFile(filepath)
	defer file.Close()

	report := NewReport(fileScanner)

	for _, history := range report {
		nextVal := history.getNextVal()
		result += nextVal
	}

	return result
}

// Day9Part2 ...
func Day9Part2(filepath string) any {
	var result int

	file, fileScanner := readFile(filepath)
	defer file.Close()

	report := NewReport(fileScanner)

	for _, history := range report {
		nextVal := history.getPrevVal()
		result += nextVal
	}

	return result
}

// Report ...
type Report []valueHistory

type valueHistory []int

func (vh valueHistory) getNextVal() int {
	diffSet := vh.buildDiffs()

	// extrapolate
	// loop from last set to first
	lastDiffSet := len(diffSet) - 1
	lastElInLastDiffSet := len(diffSet[lastDiffSet]) - 1
	nextVal := diffSet[lastDiffSet][lastElInLastDiffSet]
	diffSet[lastDiffSet] = append(diffSet[lastDiffSet], 0)

	for i := len(diffSet) - 1; i >= 1; i-- {
		// next val equals last number in next set up + last number in current set
		nextDiffSet := i - 1
		lastElInNextDiffSet := len(diffSet[nextDiffSet]) - 1
		lastElInCurDiffSet := len(diffSet[i]) - 1
		lastNumInNextSet := diffSet[nextDiffSet][lastElInNextDiffSet]
		lastNumInCurSet := diffSet[i][lastElInCurDiffSet]
		nextVal = lastNumInNextSet + lastNumInCurSet
		diffSet[nextDiffSet] = append(diffSet[nextDiffSet], nextVal)
	}

	return nextVal
}

func (vh valueHistory) getPrevVal() int {
	diffSet := vh.buildDiffs()

	// extrapolate
	// loop from last set to first
	lastDiffSet := len(diffSet) - 1
	firstElInLastDiffSet := 0
	prevVal := diffSet[lastDiffSet][firstElInLastDiffSet]
	diffSet[lastDiffSet] = append([]int{0}, diffSet[lastDiffSet]...)

	for i := len(diffSet) - 1; i >= 1; i-- {
		// prev val equals first number in next set up - first number in current set
		nextDiffSet := i - 1
		firstElInNextDiffSet := 0
		firstElInCurDiffSet := 0
		firstNumInNextSet := diffSet[nextDiffSet][firstElInNextDiffSet]
		firstNumInCurSet := diffSet[i][firstElInCurDiffSet]
		prevVal = firstNumInNextSet - firstNumInCurSet
		diffSet[nextDiffSet] = append([]int{prevVal}, diffSet[nextDiffSet]...)
	}

	return prevVal
}

func (vh valueHistory) buildDiffs() [][]int {
	diffSet := [][]int{vh}
	vals := vh
	diffHasAllZero := false
	for !diffHasAllZero {
		allZero := true
		diffs := []int{}
		for i := 0; i < len(vals)-1; i++ {
			diff := vals[i+1] - vals[i]
			if diff != 0 {
				allZero = false
			}
			diffs = append(diffs, diff)
		}
		vals = diffs
		if allZero {
			diffHasAllZero = true
		}
		diffSet = append(diffSet, diffs)
	}
	return diffSet
}

// NewReport ...
func NewReport(fileScanner *bufio.Scanner) Report {
	report := []valueHistory{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fields := strings.Fields(line)
		vh := []int{}
		for _, field := range fields {
			val, _ := strconv.Atoi(field)
			vh = append(vh, val)
		}
		report = append(report, vh)
	}
	return Report(report)
}
