package aoc2023

import (
	"bufio"
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

	cardStrengthsWithJoker = map[byte]int{
		'J': 1,
		'T': 10,
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
	sort.Sort(Hands(hands))

	for i, hand := range hands {
		winnings := (i + 1) * hand.bid
		result += winnings
	}

	return result
}

// Day7Part2 ...
func Day7Part2(filepath string) any {
	result := 0

	file, fileScanner := readFile(filepath)
	defer file.Close()

	hands := NewHands(fileScanner)
	// fmt.Println("original", hands)
	sort.Sort(HandsWithJoker(hands))
	// fmt.Println("sorted", hands)

	for i, hand := range hands {
		winnings := (i + 1) * hand.bid
		result += winnings
	}

	return result
}

// Hands ...
type Hands []Hand

func (h Hands) Len() int      { return len(h) }
func (h Hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Hands) Less(i, j int) bool {
	handTypeI := h[i].handType()
	handTypeJ := h[j].handType()

	if handTypeI != handTypeJ {
		return handTypeI.less(handTypeJ)
	}
	return h[i].orderLess(h[j], cardStrengths)
}

// HandsWithJoker ...
type HandsWithJoker []Hand

func (h HandsWithJoker) Len() int      { return len(h) }
func (h HandsWithJoker) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h HandsWithJoker) Less(i, j int) bool {
	handTypeI := h[i].handTypeWithJoker()
	handTypeJ := h[j].handTypeWithJoker()

	if handTypeI != handTypeJ {
		return handTypeI.less(handTypeJ)
	}
	return h[i].orderLess(h[j], cardStrengthsWithJoker)
}

// Hand ...
type Hand struct {
	cards string
	bid   int
}

// returns true of h hand order is less than h2 hand
func (h Hand) orderLess(h2 Hand, strengths map[byte]int) bool {
	for i := 0; i < len(h.cards); i++ {
		card1 := h.cards[i]
		card2 := h2.cards[i]

		card1Strength := 0
		card2Strength := 0
		if strength, ok := strengths[card1]; ok {
			card1Strength = strength
		} else {
			card1Strength, _ = strconv.Atoi(string(card1))
		}

		if strength, ok := strengths[card2]; ok {
			card2Strength = strength
		} else {
			card2Strength, _ = strconv.Atoi(string(card2))
		}

		if card1Strength < card2Strength {
			return true
		} else if card2Strength < card1Strength {
			return false
		}
	}
	return false
}

// HandType ...
type HandType int64

// String ...
func (ht HandType) String() string {
	switch ht {
	case HighCard:
		return "high card"
	case OnePair:
		return "one pair"
	case TwoPair:
		return "two pair"
	case ThreeOfAKind:
		return "three of a kind"
	case FullHouse:
		return "full house"
	case FourOfAKind:
		return "four of a kind"
	case FiveOfAKind:
		return "five of a kind"
	}
	return "unknown"
}

func (ht HandType) less(ht2 HandType) bool {
	return ht < ht2
}

const (
	// UnknownHand ...
	UnknownHand HandType = iota

	// HighCard ...
	HighCard

	// OnePair ...
	OnePair

	// TwoPair ...
	TwoPair

	// ThreeOfAKind ...
	ThreeOfAKind

	// FullHouse ...
	FullHouse

	// FourOfAKind ...
	FourOfAKind

	// FiveOfAKind ...
	FiveOfAKind
)

// 5, 0 - five of a kind
// 4, 1 - four of a kind
// 3, 2 - full house
// 3, 1 - three of a kind
// 2, 2 - two pair
// 2, 1 - one pair
// 1, 1 - high card
func (h Hand) handType() HandType {
	cardMap := h.getCardMap()
	var freq []int
	for _, v := range cardMap {
		freq = append(freq, v)
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})
	if freq[0] == 5 {
		return FiveOfAKind
	} else if freq[0] == 4 && freq[1] == 1 {
		return FourOfAKind
	} else if freq[0] == 3 && freq[1] == 2 {
		return FullHouse
	} else if freq[0] == 3 && freq[1] == 1 {
		return ThreeOfAKind
	} else if freq[0] == 2 && freq[1] == 2 {
		return TwoPair
	} else if freq[0] == 2 && freq[1] == 1 {
		return OnePair
	} else if freq[0] == 1 && freq[1] == 1 {
		return HighCard
	}
	return UnknownHand
}

func (h Hand) handTypeWithJoker() HandType {
	cardMap := h.getCardMap()
	jokers := 0
	var freq []int
	for k, v := range cardMap {
		if k == 'J' {
			jokers += v
		} else {
			freq = append(freq, v)
		}
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})
	if jokers == 5 {
		return FiveOfAKind
	}
	freq[0] += jokers
	if freq[0] == 5 {
		return FiveOfAKind
	} else if freq[0] == 4 && freq[1] == 1 {
		return FourOfAKind
	} else if freq[0] == 3 && freq[1] == 2 {
		return FullHouse
	} else if freq[0] == 3 && freq[1] == 1 {
		return ThreeOfAKind
	} else if freq[0] == 2 && freq[1] == 2 {
		return TwoPair
	} else if freq[0] == 2 && freq[1] == 1 {
		return OnePair
	} else if freq[0] == 1 && freq[1] == 1 {
		return HighCard
	}
	return UnknownHand
}

// returns desc map of freq of cards
func (h Hand) getCardMap() map[rune]int {
	cardMap := map[rune]int{}
	for _, c := range h.cards {
		cardMap[c]++
	}
	return cardMap
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
