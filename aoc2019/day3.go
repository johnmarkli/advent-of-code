package aoc2019

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type point struct {
	x int
	y int
}

// Day3Part1 ...
func Day3Part1(filepath string) any {
	result := 0
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// lines := make([]string, 2)
	// lineCount := 0
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	lines[lineCount] = string(scanner.Text())
	// 	lineCount++
	// }
	// line1Str := lines[0]
	// line2Str := lines[1]
	// line1 := strings.Split(line1Str, ",")
	// line2 := strings.Split(line2Str, ",")
	// TODO: fix ClosestIntersectionDist
	// result := ClosestIntersectionDist(line1, line2)
	return result
}

func intersectionLeastSteps(line1 []string, line2 []string) int {
	l1 := buildLine(line1)
	l2 := buildLine(line2)
	_ = findIntersections(l1, l2)
	return 0
}

func numSteps(lineSteps []string, intersection point) int {
	for _, move := range lineSteps {
		dir := string(move[0])
		_, _ = strconv.Atoi(string(move[1:]))
		switch dir {
		case "U":
		case "D":
		case "L":
		case "R":
		default:
			fmt.Println("Invalid line")
		}
	}
	return 0
}

// ClosestIntersectionDist ...
func ClosestIntersectionDist(line1 []string, line2 []string) int {
	l1 := buildLine(line1)
	l2 := buildLine(line2)
	intersections := findIntersections(l1, l2)
	var closestDist int
	for i, p := range intersections {
		if i == 0 {
			closestDist = ManhattanDistance(point{0, 0}, intersections[0])
		} else {
			d := ManhattanDistance(point{0, 0}, p)
			if d < closestDist {
				closestDist = d
			}
		}
	}
	return closestDist
}

func buildLine(line []string) (result []point) {
	currPos := point{0, 0}
	result = append(result, currPos)
	for _, move := range line {
		dir := string(move[0])
		steps, _ := strconv.Atoi(string(move[1:]))
		switch dir {
		case "U":
			// fmt.Println("Up ", steps)
			for i := 0; i < steps; i++ {
				currPos = point{currPos.x, currPos.y + 1}
				result = append(result, currPos)
			}
		case "D":
			// fmt.Println("Down ", steps)
			for i := 0; i < steps; i++ {
				currPos = point{currPos.x, currPos.y - 1}
				result = append(result, currPos)
			}
		case "L":
			// fmt.Println("Left ", steps)
			for i := 0; i < steps; i++ {
				currPos = point{currPos.x - 1, currPos.y}
				result = append(result, currPos)
			}
		case "R":
			// fmt.Println("Right ", steps)
			for i := 0; i < steps; i++ {
				currPos = point{currPos.x + 1, currPos.y}
				result = append(result, currPos)
			}
		default:
			fmt.Errorf("Invalid line")
		}
	}
	return result
}

func findIntersections(line1 []point, line2 []point) (result []point) {
	for i := 0; i < len(line1); i++ {
		for j := 0; j < len(line2); j++ {
			if line1[i].x == line2[j].x && line1[i].y == line2[j].y {
				// fmt.Println("found intersection", line1[i])
				result = append(result, line1[i])
			}
		}
	}
	return result[1:]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ManhattanDistance ...
func ManhattanDistance(point1, point2 point) int {
	return abs(point1.x-point2.x) + abs(point1.y-point2.y)
}
