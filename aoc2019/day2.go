package aoc2019

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	op1  = 1
	op2  = 2
	op99 = 99
)

// Day2Part1 ...
func Day2Part1(filepath string) any {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	arrStrData := strings.Split(string(fileBytes), ",")
	var arrData []int
	for _, el := range arrStrData {
		num, _ := strconv.Atoi(el)
		arrData = append(arrData, num)
	}

	arrPart1 := make([]int, len(arrData))
	copy(arrPart1, arrData)
	arrPart1[1] = 12
	arrPart1[2] = 2
	part1(arrPart1)
	return arrPart1[0]
}

// Day2Part2 ...
func Day2Part2(filepath string) any {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	arrStrData := strings.Split(string(fileBytes), ",")
	var arrData []int
	for _, el := range arrStrData {
		num, _ := strconv.Atoi(el)
		arrData = append(arrData, num)
	}

	arrPart1 := make([]int, len(arrData))
	copy(arrPart1, arrData)
	arrPart1[1] = 12
	arrPart1[2] = 2
	part1(arrPart1)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			arrCopy := make([]int, len(arrData))
			copy(arrCopy, arrData)
			arrCopy[1] = i
			arrCopy[2] = j
			part1(arrCopy)
			if arrCopy[0] == 19690720 {
				// fmt.Println("FOUND", i, j)
				result := (100 * i) + j
				return result
			}
		}
	}
	fmt.Println("not found..")
	return 0
}

// func day2() {
// 	fmt.Println("Day 2")
// 	fileBytes, err := ioutil.ReadFile("testdata/input2")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	arrStrData := strings.Split(string(fileBytes), ",")
// 	var arrData []int
// 	for _, el := range arrStrData {
// 		num, _ := strconv.Atoi(el)
// 		arrData = append(arrData, num)
// 	}
//
// 	arrPart1 := make([]int, len(arrData))
// 	copy(arrPart1, arrData)
// 	arrPart1[1] = 12
// 	arrPart1[2] = 2
// 	part1(arrPart1)
// 	fmt.Println("Day 2 Part 1 Answer: ", arrPart1[0])
//
// 	part2(arrData)
// }

func part1(arrData []int) {
	for i := 0; i < (len(arrData)); i += 4 {
		num := arrData[i]
		pos1 := arrData[i+1]
		pos2 := arrData[i+2]
		pos3 := arrData[i+3]
		if pos1 < len(arrData) && pos2 < len(arrData) && pos3 < len(arrData) {
			switch num {
			case op1:
				arrData[pos3] = arrData[pos1] + arrData[pos2]
			case op2:
				arrData[pos3] = arrData[pos1] * arrData[pos2]
			case op99:
				return
			default:
			}
		}
	}

}

// func part2(arrData []int) {
// 	for i := 0; i < 100; i++ {
// 		for j := 0; j < 100; j++ {
// 			arrCopy := make([]int, len(arrData))
// 			copy(arrCopy, arrData)
// 			arrCopy[1] = i
// 			arrCopy[2] = j
// 			part1(arrCopy)
// 			if arrCopy[0] == 19690720 {
// 				// fmt.Println("FOUND", i, j)
// 				result := (100 * i) + j
// 				fmt.Println("Day 2 Part 2 Answer: ", result)
// 				return
// 			}
// 		}
// 	}
// 	fmt.Println("not found..")
// }
