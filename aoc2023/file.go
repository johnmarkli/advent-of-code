package aoc2023

import (
	"bufio"
	"os"
)

func readFile(filepath string) (*os.File, *bufio.Scanner) {
	readFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return readFile, fileScanner
}
