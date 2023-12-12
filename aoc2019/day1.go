// Package aoc2019 ...
package aoc2019

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Day1Part1 ...
func Day1Part1(filepath string) any {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fuelSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		mass, _ := strconv.Atoi(string(scanner.Text()))
		fuel := (mass / 3) - 3
		fuelSum += fuel
	}

	return fuelSum
}

// Day1Part2 ...
func Day1Part2(filepath string) any {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fuelSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, _ := strconv.Atoi(string(scanner.Text()))
		fuel := calcFuel(mass)
		fuelSum += fuel
	}

	return fuelSum
}

func calcFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel + calcFuel(fuel)
}
