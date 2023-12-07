package day7

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

type Hand struct {
	cards []int
}

const (
	FiveOfAKind  = 7
	FourOfAKind  = 6
	FullHouse    = 5
	ThreeOfAKind = 4
	TwoPair      = 3
	OnePair      = 2
	HighCard     = 1
	InvalidHand  = 0
)

func classigyCounts(cardCounts map[int]int) int {
	switch len(cardCounts) {
	case 1:
		return FiveOfAKind
	case 2:
		if containsCardCount(cardCounts, 4) {
			return FourOfAKind
		}
		return FullHouse
	case 3:
		if containsCardCount(cardCounts, 3) {
			return ThreeOfAKind
		}
		return TwoPair
	case 4:
		return OnePair
	case 5:
		return HighCard
	}
	return InvalidHand
}
func (h Hand) Type() int {
	cardCounts := make(map[int]int)
	for _, c := range h.cards {
		cardCounts[c] = 0
	}
	for _, c := range h.cards {
		cardCounts[c]++
	}

	return classigyCounts(cardCounts)
}

type Play struct {
	hand Hand
	bid  int
}

func CompareHands(a, b Hand) int {
	hType := a.Type()
	oType := b.Type()
	if hType != oType {
		return cmp.Compare(hType, oType)
	}
	for i := 0; i < len(a.cards); i++ {
		if a.cards[i] != b.cards[i] {
			return cmp.Compare(a.cards[i], b.cards[i])
		}
	}
	return 0
}

func HandTypeWithJokers(h Hand) int {
	cardCounts := make(map[int]int)
	for _, c := range h.cards {
		cardCounts[c] = 0
	}
	for _, c := range h.cards {
		cardCounts[c]++
	}

	numJokers := cardCounts[1] // joker == 1
	if numJokers == 5 {
		return FiveOfAKind
	}
	delete(cardCounts, 1)
	// add jokers to the highest count card
	maxCount := 0
	for _, v := range cardCounts {
		if v > maxCount {
			maxCount = v
		}
	}
	for k, v := range cardCounts {
		if v == maxCount {
			cardCounts[k] += numJokers
			break
		}
	}

	return classigyCounts(cardCounts)
}

func CompareHandsWithJokers(a, b Hand) int {
	hType := HandTypeWithJokers(a)
	oType := HandTypeWithJokers(b)
	if hType != oType {
		return cmp.Compare(hType, oType)
	}
	for i := 0; i < len(a.cards); i++ {
		if a.cards[i] != b.cards[i] {
			return cmp.Compare(a.cards[i], b.cards[i])
		}
	}
	return 0
}

func containsCardCount(cc map[int]int, count int) bool {
	for _, v := range cc {
		if v == count {
			return true
		}
	}
	return false
}

func parseInput(data string) []Play {
	var pp []Play

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		var cc []int
		for _, c := range parts[0] {
			val := 0
			switch string(c) {
			case "A":
				val = 14
			case "K":
				val = 13
			case "Q":
				val = 12
			case "J":
				val = 11
			case "T":
				val = 10
			default:
				val, _ = strconv.Atoi(string(c))
			}
			cc = append(cc, val)
		}

		pp = append(pp, Play{bid: bid, hand: Hand{cc}})
	}

	return pp
}

func calculateWinnings(plays []Play, cmpFn func(a, b Hand) int) int {
	slices.SortFunc(plays, func(a, b Play) int {
		return cmpFn(a.hand, b.hand)
	})
	winnings := 0
	for i, p := range plays {
		winnings += (i + 1) * p.bid
	}
	return winnings
}

func Part1() string {
	plays := parseInput(input)
	return strconv.Itoa(calculateWinnings(plays, CompareHands))
}

func Part2() string {
	plays := parseInput(input)
	// special rule: J are worth 1
	for _, p := range plays {
		for i, c := range p.hand.cards {
			if c == 11 {
				p.hand.cards[i] = 1
			}
		}
	}
	winnings := calculateWinnings(plays, CompareHandsWithJokers)
	fmt.Println("correct answer: 247899149")
	return strconv.Itoa(winnings)
}
