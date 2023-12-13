package aoc2023

import (
	"bufio"
	"fmt"
	"strings"
)

// Day8Part1 ...
func Day8Part1(filepath string) any {
	result := 0

	file, fileScanner := readFile(filepath)
	defer file.Close()

	network := NewNetwork(fileScanner)

	curNode := network.nodes["AAA"]
	i := 0
	for curNode.val != "ZZZ" {
		instruction := network.instructions[i]
		switch instruction {
		case "R":
			curNode = curNode.right
		case "L":
			curNode = curNode.left
		}
		result++
		if i == len(network.instructions)-1 {
			i = 0
		} else {
			i++
		}
	}

	return result
}

// Day8Part2 ...
func Day8Part2(filepath string) any {
	result := 0

	file, fileScanner := readFile(filepath)
	defer file.Close()

	network := NewNetwork(fileScanner)

	var curNodes nodeGroup
	// find nodes that end in A
	for _, node := range network.nodes {
		if node.endsIn("A") {
			curNodes = append(curNodes, node)
		}
	}

	// trick: each starting node will get to Z at different periods, but they
	// should all eventually line up -> calculate Lowest Common Multiple

	// find multiple for each starting xxA node
	multiples := map[int]int{}

	// while nodes don't all end in Z
	inst := 0
	for len(multiples) < len(curNodes) {
		// loop through nodes and follow instructions for each
		instruction := network.instructions[inst]
		for i, curNode := range curNodes {
			switch instruction {
			case "R":
				curNodes[i] = curNode.right
			case "L":
				curNodes[i] = curNode.left
			}
			if curNodes[i].endsIn("Z") {
				if _, ok := multiples[i]; !ok {
					multiples[i] = result + 1
				}
			}
		}
		result++
		if inst == len(network.instructions)-1 {
			inst = 0
		} else {
			inst++
		}
	}
	var nums []int
	for _, num := range multiples {
		nums = append(nums, num)
	}

	result = lcm(nums)

	return result
}

func lcm(nums []int) int {
	result := lcmSingle(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		result = lcmSingle(result, nums[i])
	}
	return result
}

func lcmSingle(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type node struct {
	val      string
	leftVal  string
	rightVal string
	left     *node
	right    *node
}

func (n *node) endsIn(c string) bool {
	return string(n.val[len(n.val)-1]) == c
}

type nodeGroup []*node

func (ng nodeGroup) endsIn(c string) bool {
	for _, node := range ng {
		if !node.endsIn(c) {
			return false
		}
	}
	return true
}

func (n *node) String() string {
	return fmt.Sprintf("%s = (%s, %s)", n.val, n.left.val, n.right.val)
}

// Network ...
type Network struct {
	nodes        map[string]*node
	instructions []string
}

// NewNetwork ...
func NewNetwork(fileScanner *bufio.Scanner) Network {
	nodes := map[string]*node{}

	// instructions line
	fileScanner.Scan()
	line := fileScanner.Text()
	var instructions []string
	for _, c := range line {
		instructions = append(instructions, string(c))
	}

	// empty line
	fileScanner.Scan()
	_ = fileScanner.Text()

	// rest of the nodes
	for fileScanner.Scan() {
		line = fileScanner.Text()
		node := makeNode(line)
		nodes[node.val] = node
	}

	// connect nodes
	for _, node := range nodes {
		node.left = nodes[node.leftVal]
		node.right = nodes[node.rightVal]
	}

	return Network{
		nodes:        nodes,
		instructions: instructions,
	}
}

func makeNode(line string) *node {
	split := strings.Split(line, " = ")
	val := split[0]
	leavesStart := 1
	leavesEnd := len(split[1]) - 1
	leaves := split[1][leavesStart:leavesEnd]
	leavesSplit := strings.Split(leaves, ", ")
	leftVal := leavesSplit[0]
	rightVal := leavesSplit[1]
	return &node{
		val:      val,
		rightVal: rightVal,
		leftVal:  leftVal,
	}
}
