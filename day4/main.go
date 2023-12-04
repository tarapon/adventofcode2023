package main

import (
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type card struct {
	id      int
	winning []int
	numbers []int
	copies  int
}

func (c *card) inc(value int) int {
	c.copies += value
	return c.copies
}

func (c *card) winningCount() int {
	return len(c.intersection())
}

func (c *card) intersection() []int {
	intersection := make([]int, 0)

	for _, a := range c.numbers {
		for _, b := range c.winning {
			if a == b {
				intersection = append(intersection, a)
				break
			}
		}
	}

	return intersection
}

func (c *card) points() int {
	return int(math.Pow(2, float64(c.winningCount()-1)))
}

func main() {
	content, err := os.ReadFile("day4/input.txt")
	if err != nil {
		panic(err)
	}

	res1, res2 := 0, 0
	lines := strings.Split(string(content), "\n")

	cards := parseInput(lines)

	for i, card := range cards {
		res1 += card.points()

		for j := i + 1; j < len(cards) && j-i-1 < card.winningCount(); j++ {
			cards[j].inc(card.copies)
		}

		res2 += card.copies
	}

	log.Println("scratchcards (simple):", res1)
	log.Println("scratchcards (advanced):", res2)
}

func mustParseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return val
}

func parseNumbers(s string) []int {
	numbers := make([]int, 0)
	s = strings.Trim(s, " ")

	for _, part := range regexp.MustCompile(`\s+`).Split(s, -1) {
		numbers = append(numbers, mustParseInt(part))
	}

	return numbers
}

func parseInput(lines []string) []card {
	cards := make([]card, 0)

	for _, line := range lines {
		card := card{}

		re := regexp.MustCompile(`^Card\s+(\d+): ([\d\s]+) \| ([\d\s]+)$`)
		parts := re.FindStringSubmatch(line)
		card.id = mustParseInt(parts[1])
		card.winning = parseNumbers(parts[2])
		card.numbers = parseNumbers(parts[3])
		card.copies = 1

		cards = append(cards, card)
	}

	return cards
}
