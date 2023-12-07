package main

import (
	h "advcode2025/helper"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("day6/input.txt")
	if err != nil {
		panic(err)
	}

	res1, res2 := 1, 0
	lines := strings.Split(string(content), "\n")
	races := parseInput1(lines)
	bigRace := parseInput2(lines)

	for _, race := range races {
		res1 *= race.winVariants()
	}

	res2 = bigRace.winVariants()

	log.Println(races)

	log.Println("win factor (simple):", res1)
	log.Println("win factor (advanced):", res2)
}

type race struct {
	time     int
	distance int
}

func (r race) holdAndGo(t int) int {
	return (r.time - t) * t
}

func (r race) winVariants() int {
	k := 0

	for i := 1; i < r.time; i++ {
		if r.holdAndGo(i) > r.distance {
			k++
		}
	}

	return k
}

func parseInput1(lines []string) []race {
	re := regexp.MustCompile(`(\d+)+`)

	times := h.Map[[]string, int](re.FindAllStringSubmatch(lines[0], -1), func(match []string) int {
		return h.Must[int](strconv.Atoi(match[1]))
	})

	distances := h.Map[[]string, int](re.FindAllStringSubmatch(lines[1], -1), func(match []string) int {
		return h.Must[int](strconv.Atoi(match[1]))
	})

	races := make([]race, 0, len(times))

	for i, t := range times {
		races = append(races, race{time: t, distance: distances[i]})
	}

	return races
}

func parseInput2(lines []string) race {
	re := regexp.MustCompile(`\s+`)
	line1 := re.ReplaceAllString(lines[0], "")
	line2 := re.ReplaceAllString(lines[1], "")

	return race{
		time:     h.Must[int](strconv.Atoi(strings.Split(line1, ":")[1])),
		distance: h.Must[int](strconv.Atoi(strings.Split(line2, ":")[1])),
	}
}
