package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day4() {
	fmt.Println("Day 4")
	input := "272091-815432"
	pRange := strings.Split(input, "-")
	low, _ := strconv.Atoi(pRange[0])
	high, _ := strconv.Atoi(pRange[1])
	// fmt.Println(low, high)

	passCount := 0

	// loop through pRange
	for num := low; num <= high; num++ {
		// fmt.Println(num)
		digits := strconv.Itoa(num)

		hasAdjacent := false
		isIncreasing := true

		prev := 0
		for i := 0; i < len(digits); i++ {
			// test for adjacent digits
			if i > 0 {
				if digits[i] == digits[prev] {
					hasAdjacent = true
				}
				if digits[i] < digits[prev] {
					isIncreasing = false
				}
			}

			// test for increasing digits
			prev = i
		}

		if hasAdjacent && isIncreasing {
			// fmt.Println("matching password", num)
			passCount++
		}

	}
	fmt.Println("Day 4 Part 1: ", passCount)
}
