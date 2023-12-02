package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}

	games := parseInput(string(content))

	res1, res2 := 0, 0

	for _, game := range games {
		if game.IsPossible(dice{red: 12, green: 13, blue: 14}) {
			res1 += game.id
		}
	}

	log.Println("game (simple):", res1)
	log.Println("game (advanced):", res2)
}

type dice struct {
	red   int
	green int
	blue  int
}

func (d dice) Gte(o dice) bool {
	return d.red >= o.red && d.green >= o.green && d.blue >= o.blue
}

type game struct {
	id    int
	dices []dice
}

func (g game) IsPossible(o dice) bool {
	for _, d := range g.dices {
		if !o.Gte(d) {
			return false
		}
	}

	return true
}

func mustInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return res
}

func parseInput(s string) []game {
	lines := strings.Split(s, "\n")
	res := make([]game, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ":")
		idTag := parts[0]
		dicesTags := strings.Split(parts[1], ";")

		id := mustInt(strings.Split(idTag, " ")[1])
		dices := make([]dice, 0, len(dicesTags))

		for _, diceTag := range dicesTags {
			dice := dice{}

			re := regexp.MustCompile("(\\d+) (red|green|blue)")
			for _, m := range re.FindAllStringSubmatch(diceTag, -1) {
				switch m[2] {
				case "red":
					dice.red = mustInt(m[1])
				case "green":
					dice.green = mustInt(m[1])
				case "blue":
					dice.blue = mustInt(m[1])
				}
			}

			dices = append(dices, dice)
		}

		res = append(res, game{id: id, dices: dices})
	}

	return res
}
