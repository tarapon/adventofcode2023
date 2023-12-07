package main

import (
	h "advcode2025/helper"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("day7/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	decks := parseDecks(lines)

	res1, res2 := score(decks, false), score(decks, true)

	log.Println("score (simple):", res1)
	log.Println("score (jokers):", res2)
}

func score(decks []hand, joker bool) int {
	res := 0
	slices.SortFunc(decks, func(a, b hand) int {
		return a.compare(b, joker)
	})

	for i, deck := range decks {
		res += deck.bet * (i + 1)
	}

	return res
}

type hand struct {
	deck []byte
	bet  int
}

func (h hand) setScore(m map[byte]int) int {
	keys := make([]byte, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	switch len(keys) {
	case 1:
		return 7 // five of a kind
	case 2:
		if m[keys[0]] == 4 || m[keys[1]] == 4 {
			return 6 // four of a kind
		} else {
			return 5 // full house
		}
	case 3:
		if m[keys[0]] == 3 || m[keys[1]] == 3 || m[keys[2]] == 3 {
			return 4 // three of a kind
		} else {
			return 3 // two pairs
		}
	case 4:
		return 2 // one pair
	default:
		return 1 // high card
	}
}

func (h hand) rank(joker bool) int {
	m := make(map[byte]int)
	kMax, vMax := byte(0), 0

	for _, card := range h.deck {
		m[card]++

		if m[card] > vMax && card != 'J' {
			vMax = m[card]
			kMax = card
		}
	}

	if joker {
		m[kMax] += m['J']
		delete(m, 'J')
	}

	return h.setScore(m)
}

func (h hand) compare(other hand, joker bool) int {
	rankA := h.rank(joker)
	rankB := other.rank(joker)

	if rankA != rankB {
		return rankA - rankB
	}

	cards := map[byte]int{
		'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10,
		'9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
	}

	if joker {
		cards['J'] = 1
	}

	for i := 0; i < 5; i++ {
		if h.deck[i] != other.deck[i] {
			return cards[h.deck[i]] - cards[other.deck[i]]
		}
	}

	return 0
}

func parseDecks(lines []string) []hand {
	return h.Map(lines, func(line string) hand {
		parts := strings.Split(line, " ")

		return hand{
			deck: []byte(parts[0]),
			bet:  h.Must(strconv.Atoi(parts[1])),
		}
	})
}
