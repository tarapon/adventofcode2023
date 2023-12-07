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

	res1, res2 := 0, 0
	lines := strings.Split(string(content), "\n")
	races := parseInput(lines)

	//for _, race := range races {
	//	res1 += 0
	//	res2 += 0
	//}

	log.Println(races)

	log.Println("win factor (simple):", res1)
	log.Println("win factor (advanced):", res2)
}

type race struct {
	time     int
	distance int
}

func parseInput(lines []string) []race {
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
