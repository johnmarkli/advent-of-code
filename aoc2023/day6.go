package aoc2023

import (
	"bufio"
	"strconv"
	"strings"
)

// Day6Part1 ...
func Day6Part1(filepath string) any {
	result := 1
	numWaysToWin := []int{}

	file, fileScanner := readFile(filepath)
	defer file.Close()

	// parse races
	rrs := NewRaceRecords(fileScanner)

	// for each race record
	for _, rr := range rrs {
		wins := 0
		// calculate distances for each number of ms the button is held
		for i := 0; i < rr.maxTime; i++ {
			// if distance is greater than record, incr num wins
			if rr.holdWin(i) {
				wins++
			}
		}
		numWaysToWin = append(numWaysToWin, wins)
	}

	for _, num := range numWaysToWin {
		result *= num
	}
	return result
}

// Day6Part2 ...
func Day6Part2(filepath string) any {
	result := 0

	file, fileScanner := readFile(filepath)
	defer file.Close()

	// parse races
	rr := NewRaceRecord(fileScanner)

	// calculate distances for each number of ms the button is held
	for i := 0; i < rr.maxTime; i++ {
		// if distance is greater than record, incr num wins
		if rr.holdWin(i) {
			result++
		}
	}

	return result
}

// RaceRecords ..
type RaceRecords []RaceRecord

// RaceRecord ..
type RaceRecord struct {
	maxTime     int
	maxDistance int
}

func (rr *RaceRecord) holdWin(held int) bool {
	speed := held
	timeLeft := rr.maxTime - held
	dist := speed * timeLeft
	if dist > rr.maxDistance {
		return true
	}
	return false
}

// NewRaceRecord ...
func NewRaceRecord(fileScanner *bufio.Scanner) RaceRecord {
	// First line times
	fileScanner.Scan()
	timeLine := fileScanner.Text()
	timeStrs := strings.Fields(timeLine)
	timeStr := strings.Join(timeStrs[1:], "")

	// Second line distances
	fileScanner.Scan()
	distanceLine := fileScanner.Text()
	distanceStrs := strings.Fields(distanceLine)
	distanceStr := strings.Join(distanceStrs[1:], "")

	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)
	return RaceRecord{
		maxTime:     time,
		maxDistance: distance,
	}
}

// NewRaceRecords ...
func NewRaceRecords(fileScanner *bufio.Scanner) []RaceRecord {
	rrs := []RaceRecord{}

	// First line times
	fileScanner.Scan()
	timeLine := fileScanner.Text()
	timeStrs := strings.Fields(timeLine)

	// Second line distances
	fileScanner.Scan()
	distanceLine := fileScanner.Text()
	distanceStrs := strings.Fields(distanceLine)

	if len(timeStrs) == len(distanceStrs) {
		for i := 1; i < len(timeStrs); i++ {
			time, _ := strconv.Atoi(timeStrs[i])
			distance, _ := strconv.Atoi(distanceStrs[i])
			rrs = append(rrs, RaceRecord{
				maxTime:     time,
				maxDistance: distance,
			})
		}
	}
	return rrs
}
