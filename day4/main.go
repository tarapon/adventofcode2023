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
}

func (c *card) interception() []int {
	interception := make([]int, 0)

	for _, a := range c.numbers {
		for _, b := range c.winning {
			if a == b {
				interception = append(interception, a)
				break
			}
		}
	}

	return interception
}

func (c *card) points() int {
	nums := c.interception()
	return int(math.Pow(float64(2), float64(len(nums)-1)))
}

func main() {
	content, err := os.ReadFile("day4/input.txt")
	if err != nil {
		panic(err)
	}

	res1, res2 := 0, 0
	lines := strings.Split(string(content), "\n")

	cards := parseInput(lines)

	for _, card := range cards {
		res1 += card.points()
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

		cards = append(cards, card)
	}

	return cards
}
