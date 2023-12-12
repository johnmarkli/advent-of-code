package aoc2023

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	cardStrengths = map[byte]int{
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
)

// Day7Part1 ...
func Day7Part1(filepath string) any {
	result := 0

	file, fileScanner := readFile(filepath)
	defer file.Close()

	hands := NewHands(fileScanner)
	fmt.Println("original", hands)
	sort.Sort(Hands(hands))
	fmt.Println("sorted", hands)

	for i, hand := range hands {
		winnings := (i + 1) * hand.bid
		result += winnings
	}

	return result
}

// Day7Part2 ...
func Day7Part2(filepath string) any {
	result := 0

	// file, fileScanner := readFile(filepath)
	// defer file.Close()

	return result
}

// Hands ...
type Hands []Hand

func (h Hands) Len() int      { return len(h) }
func (h Hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Hands) Less(i, j int) bool {
	numKindsI := h[i].getNumKinds()
	numKindsJ := h[j].getNumKinds()
	fmt.Println("I", numKindsI)
	fmt.Println("J", numKindsJ)
	return (numKindsI[0] < numKindsJ[0]) || (numKindsI[0] == numKindsJ[0] &&
		numKindsI[1] < numKindsJ[1]) && compareHands(h[i], h[j])
}

// Hand ...
type Hand struct {
	cards    string
	bid      int
	numKinds []int
}

func compareHands(h1 Hand, h2 Hand) bool {
	if len(h1.cards) != len(h1.cards) {
		return false
	}

	for i := 0; i < len(h1.cards); i++ {
		h1Strength := 0
		h2Strength := 0

		if strength, ok := cardStrengths[h1.cards[i]]; ok {
			h1Strength = strength
		} else {
			h1Strength, _ = strconv.Atoi(string(h1.cards[i]))
		}

		if strength, ok := cardStrengths[h2.cards[i]]; ok {
			h2Strength = strength
		} else {
			h2Strength, _ = strconv.Atoi(string(h2.cards[i]))
		}
		if h1Strength < h2Strength {
			return true
		}
	}

	return false
}

// returns top 2 frequent cards
// 5, 0 - five of a kind
// 4, 1 - four of a kind
// 3, 2 - full house
// 3, 1 - three of a kind
// 2, 2 - two pair
// 2, 1 - one pair
// 1, 1 - high card
func (h Hand) getNumKinds() []int {
	if len(h.numKinds) > 0 {
		return h.numKinds
	}
	cardMap := map[rune]int{}
	for _, c := range h.cards {
		cardMap[c]++
	}
	var sortedSlice []int
	for _, v := range cardMap {
		sortedSlice = append(sortedSlice, v)
	}
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i] > sortedSlice[j]
	})
	h.numKinds = sortedSlice
	return h.numKinds
}

func (h Hands) sort() Hands {
	return h
}

// NewHands ...
func NewHands(fileScanner *bufio.Scanner) Hands {
	hands := []Hand{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fields := strings.Fields(line)
		bid, _ := strconv.Atoi(fields[1])
		hands = append(hands, Hand{
			cards: fields[0],
			bid:   bid,
		})
	}
	return hands
}
