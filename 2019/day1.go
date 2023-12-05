// Package main ...
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day1() {
	fmt.Println("Day 1")
	day1Part1()
	day1Part2()
}

func day1Part1() {
	file, err := os.Open("testdata/input1")
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

	fmt.Println("Day 1 Part 1 Answer: ", fuelSum)
}

func day1Part2() {
	file, err := os.Open("testdata/input1")
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

	fmt.Println("Day 1 Part 2 Answer: ", fuelSum)
}

func calcFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel + calcFuel(fuel)
}
