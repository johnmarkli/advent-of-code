package aoc2023

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	springUnknown = '?'
	springBroken  = '#'
	springOp      = '.'
)

// Day12Part1 ...
func Day12Part1(filepath string) any {
	var result int

	file, scanner := readFile(filepath)
	defer file.Close()

	sm := newSpringMap(scanner)

	for _, sr := range sm {
		arrs := sr.Arrangements()
		result += arrs
	}

	return result
}

// Day12Part2 ...
func Day12Part2(filepath string) any {
	var result int

	file, scanner := readFile(filepath)
	defer file.Close()

	sm := newSpringMap(scanner)
	sm.unfold()

	for _, sr := range sm {
		arrs := sr.Arrangements()
		result += arrs
	}

	return result
}

type springMap []*SpringRow

func (sm springMap) String() string {
	var out string
	for _, r := range sm {
		out += r.String() + "\n"
	}
	return out
}

func (sm springMap) unfold() {
	for _, sr := range sm {
		sr.unfold()
	}
}

// SpringRow ...
type SpringRow struct {
	Springs []byte
	Groups  []int
	memo    map[string]int
}

// String ...
func (sr *SpringRow) String() string {
	return fmt.Sprintf("%s %v", sr.Springs, sr.Groups)
}

// unfold by replacing springs with 5 copies of itself with a ? in between
// then replace groups with 5 copies of itself
func (sr *SpringRow) unfold() {
	var springs []byte
	var groups []int

	for i := 1; i <= 5; i++ {
		springs = append(springs, sr.Springs...)
		if i < 5 {
			springs = append(springs, springUnknown)
		}
	}

	for i := 1; i <= 5; i++ {
		groups = append(groups, sr.Groups...)
	}

	sr.Springs = springs
	sr.Groups = groups
}

// Arrangements ...
func (sr *SpringRow) Arrangements() int {
	// follow tree of possibilities until the end and count how many satisfy the choosing of groups
	// - how to traverse tree of possibilities
	// - how to choose groups and apply skip logic
	// - how to count num of valid arrangements that get to the length of row
	// - after a group is done, need to have at least one . before starting the next group

	result := sr.springDFS(0, sr.Groups, byte(0))

	return result
}

func (sr *SpringRow) springDFS(start int, groups []int, prevEl byte) int {
	// fmt.Println("\nstart, curRow, groups", start, string(curRow), origGroups)
	var result int
	// base case
	if start > len(sr.Springs)-1 {
		// fmt.Println("hit end of springs with groups", string(curRow), origGroups)
		if len(groups) == 0 || (len(groups) == 1 && groups[0] == 0) {
			// fmt.Println("no more groups, add to arrs", string(curRow), origGroups)
			return 1
		}
		// fmt.Println("still have groups, don't add to arrs")
		return 0
	}

	curGroup := -1
	if len(groups) > 0 {
		curGroup = groups[0]
	}

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
	if el == springUnknown {
		// choose
		if curGroup == 0 {
			choices[springOp] = groupsWithoutFirst
		} else if curGroup == -1 || (curGroup > 0 && prevEl != springBroken) {
			choices[springOp] = newGroups
		}

		if curGroup > 0 {
			choices[springBroken] = groupsWithDecr
		}
	} else { // if el is not ?
		// take el
		if el == springOp {
			if curGroup == 0 {
				choices[el] = groupsWithoutFirst
			} else if curGroup == -1 || (curGroup > 0 && prevEl != springBroken) {
				choices[el] = newGroups
			}
		} else if el == springBroken && curGroup > 0 {
			choices[el] = groupsWithDecr
		}
	}

	if len(choices) > 0 {
		for el, gs := range choices {
			// check memo for arrs first, if doesn't exist then dfs and save result to memo
			var arrs int
			foundMemo := false
			cacheKey := sr.cacheKey(start, gs, el)
			if res, ok := sr.memo[cacheKey]; ok {
				foundMemo = true
				arrs = res
			}
			if !foundMemo {
				arrs = sr.springDFS(start+1, gs, el)
				if sr.memo == nil {
					sr.memo = map[string]int{}
				}
				sr.memo[cacheKey] = arrs
			}
			result += arrs
		}
	}

	return result
}

func (sr *SpringRow) cacheKey(start int, groups []int, el byte) string {
	curGroup := -1
	if len(groups) > 0 {
		curGroup = groups[0]
	}
	return fmt.Sprintf("%d-%d-%d-%d", start, curGroup, len(groups), el)
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
			memo:    map[string]int{},
		}
		springRows = append(springRows, sr)
	}
	return springRows
}
