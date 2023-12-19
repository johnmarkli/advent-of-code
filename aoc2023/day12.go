package aoc2023

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	springUnknown     = '?'
	springBroken      = '#'
	springOperational = '.'
)

// Day12Part1 ...
func Day12Part1(filepath string) any {
	var result int

	file, scanner := readFile(filepath)
	defer file.Close()

	sm := newSpringMap(scanner)
	// fmt.Println(sm)

	for _, sr := range sm {
		arrs := sr.Arrangements()
		fmt.Println(sr, arrs)
		result += arrs
	}

	return result
}

// Day12Part2 ...
func Day12Part2(filepath string) any {
	var result int

	// file, scanner := readFile(filepath)
	// defer file.Close()

	return result
}

type springMap []*SpringRow

func (sm springMap) String() string {
	var out string
	for _, r := range sm {
		out += r.String()
	}
	return out
}

// SpringRow ...
type SpringRow struct {
	Springs []byte
	Groups  []int
}

// String ...
func (sr *SpringRow) String() string {
	return fmt.Sprintf("%s %v", sr.Springs, sr.Groups)
}

// Arrangements ...
func (sr *SpringRow) Arrangements() int {
	// follow tree of possibilities until the end and count how many satisfy the choosing of groups

	// - how to traverse tree of possibilities
	// - how to choose groups and apply skip logic
	// - how to count num of valid arrangements that get to the length of row
	// after a group is done, need to have at least one . before starting the next group

	result := sr.springDFS(0, sr.Groups, []byte{})

	return result
}

func (sr *SpringRow) springDFS(start int, origGroups []int, curRow []byte) int {
	fmt.Println("\nstart, curRow, groups", start, string(curRow), origGroups)
	var arrs int
	// base case
	if start > len(sr.Springs)-1 {
		fmt.Println("hit end of springs with groups", string(curRow), origGroups)
		if len(origGroups) == 0 || (len(origGroups) == 1 && origGroups[0] == 0) {
			fmt.Println("no more groups, add to arrs", string(curRow), origGroups)
			return 1
		}
		fmt.Println("still have groups, don't add to arrs")
		return 0
	}

	groups := make([]int, len(origGroups))
	copy(groups, origGroups)
	curGroup := -1
	if len(groups) > 0 {
		curGroup = groups[0]
	}
	fmt.Println("curGroup start", curGroup, groups)

	// iterative case

	el := sr.Springs[start]
	choices := map[byte][]int{}

	groupsWithoutFirst := make([]int, len(groups))
	copy(groupsWithoutFirst, groups)
	if len(groupsWithoutFirst) > 0 {
		groupsWithoutFirst = groups[1:]
	}

	newGroups := make([]int, len(groups))
	copy(newGroups, groups)

	groupsWithDecr := make([]int, len(groups))
	copy(groupsWithDecr, groups)
	if len(groupsWithDecr) > 0 {
		groupsWithDecr[0]--
	}

	// if el is ?
	if el == '?' {
		// choose
		if curGroup == 0 {
			choices['.'] = groupsWithoutFirst
		} else {
			choices['.'] = groups
		}

		if curGroup > 0 {
			choices['#'] = groupsWithDecr
		}
	} else { // if el is not ?
		// take el
		if el == '.' {
			if curGroup == 0 {
				choices['.'] = groupsWithoutFirst
			} else {
				choices['.'] = newGroups
			}
		} else if el == '#' && curGroup > 0 {
			choices['#'] = groupsWithDecr
		}
	}

	if len(choices) > 0 {
		for el, gs := range choices {
			newRow := append(curRow, el)
			fmt.Println("chose", string(el), "with groups", curGroup, gs, string(newRow))
			arrs += sr.springDFS(start+1, gs, newRow)
		}
	}

	return arrs
}

func newSpringMap(scanner *bufio.Scanner) springMap {
	springRows := []*SpringRow{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		springsStr := fields[0]
		groupsStr := fields[1]
		groups := []int{}
		groupFields := strings.Split(groupsStr, ",")
		for _, gf := range groupFields {
			g, _ := strconv.Atoi(gf)
			groups = append(groups, g)
		}
		sr := &SpringRow{
			Springs: []byte(springsStr),
			Groups:  groups,
		}
		springRows = append(springRows, sr)
	}
	return springRows
}

// attempt #1
// if sr.Springs[start] == '?' {
// 	fmt.Println("curGroup check ?", curGroup, groups)
//
// 	// if ?, then explore branches
//
// 	// if choosing . and curGroup is 0, then remove 0 group
// 	// don't choose a . if curGroup > 0 and last chosen was a #
// 	if len(curRow) == 0 || (len(curRow) > 0 && curRow[len(curRow)-1] != '#') {
// 		fmt.Println("curGroup check ?", curGroup, groups)
// 		if curGroup == 0 {
// 			groups = groups[1:]
// 		}
// 		newRow2 := append(curRow, '.')
// 		fmt.Println("chose . with groups", curGroup, groups, string(newRow2))
// 		arrs += sr.springDFS(start+1, groups, newRow2)
// 	}
//
// 	if curGroup > 0 {
// 		fmt.Println("curgroup beforfe adding #", curGroup)
// 		newRow1 := append(curRow, '#')
// 		curGroup--
// 		groups[0] = curGroup
// 		fmt.Println("chose # with groups", curGroup, groups, string(newRow1))
// 		arrs += sr.springDFS(start+1, groups, newRow1)
// 	}
//
// } else { // not ?, so must choose
// 	el := sr.Springs[start]
// 	// fmt.Println("need to choose", string(el), "with groups", groups, string(curRow))
// 	if el == '#' {
// 		if curGroup == 0 {
// 			fmt.Println("can't choose # since need a new group, returning")
// 			return 0
// 		}
// 		curGroup--
// 		groups[0] = curGroup
// 		// if curGroup > 0 {
// 		// 	groups[0] = curGroup
// 		// } else {
// 		// 	groups = groups[1:]
// 		// }
// 		curRow = append(curRow, el)
// 		fmt.Println("need to chose # with groups", curGroup, groups, string(curRow))
// 		arrs += sr.springDFS(start+1, groups, curRow)
// 	} else if curGroup == 0 {
// 		groups = groups[1:]
// 		curRow = append(curRow, el)
// 		fmt.Println("need to chose . with groups", curGroup, groups, string(curRow))
// 		arrs += sr.springDFS(start+1, groups, curRow)
// 	} else if curGroup > 0 {
// 		fmt.Println("can't choose . since need to finish group")
// 		return 0
// 	}
// }

// attempt #2
// // if current element is ?
// if el == '?' {
// 	// choose . and # if valid
//
// 	// if first el in curRow
// 	if len(curRow) == 0 {
// 		// add . to curRow - dfs
// 		newRow1 := append(curRow, '.')
// 		fmt.Println("chose . with groups", curGroup, groups, string(newRow1))
//
// 		// decr curGroup and add # to curRow - dfs
// 		curGroup--
// 		groups[0] = curGroup
// 		newRow2 := append(curRow, '#')
// 		fmt.Println("chose # with groups", curGroup, groups, string(newRow2))
// 		arrs += sr.springDFS(start+1, groups, newRow1)
// 	} else {
// 		prevEl := curRow[len(curRow)-1]
// 		// if previous el was a # and curGroup > 0
// 		if prevEl == '#' {
// 			if curGroup > 0 {
// 				// decr curGroup and add # to curRow - dfs
// 				curGroup--
// 				groups[0] = curGroup
// 				newRow := append(curRow, '#')
// 				fmt.Println("chose # with groups", curGroup, groups, string(newRow))
// 				arrs += sr.springDFS(start+1, groups, newRow)
// 			} else { // if previous el was a # and curGroup == 0
// 				// remove curGroup and add . to curRow - dfs
// 				newGroups := groups[1:]
// 				newRow := append(curRow, '.')
// 				fmt.Println("chose . with groups", curGroup, newGroups, string(newRow))
// 			}
// 		} else if prevEl == '.' {
// 			if curGroup == 0 {
// 				// remove curGroup and add . to curRow - dfs
// 				newGroups := groups[1:]
// 				newRow := append(curRow, '.')
// 				fmt.Println("chose . with groups", curGroup, newGroups, string(newRow))
// 			} else {
// 				// add . to curRow - dfs
// 				newRow := append(curRow, '.')
// 				fmt.Println("chose . with groups", curGroup, groups, string(newRow))
// 			}
//
// 		}
// 	}
// } else { // if current element is not ?
// 	// need to choose element if valid
//
// 	// if el is # and curGroup > 0
// 	if el == '#' && curGroup > 0 {
// 		// decr curGroup and add # to curRow - dfs
// 		curGroup--
// 		groups[0] = curGroup
// 		newRow := append(curRow, '#')
// 		fmt.Println("needed to chose # with groups", curGroup, groups, string(newRow))
// 		arrs += sr.springDFS(start+1, groups, newRow)
// 	} else if el == '.' && curGroup == 0 { // if el is . and curGroup == 0
// 		// remove curGroup and add . to curRow - dfs
// 		// remove curGroup and add . to curRow - dfs
// 		newGroups := groups[1:]
// 		newRow := append(curRow, '.')
// 		fmt.Println("chose . with groups", curGroup, newGroups, string(newRow))
// 	}
// }
