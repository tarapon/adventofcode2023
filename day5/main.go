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
	content, err := os.ReadFile("day5/input.txt")
	if err != nil {
		panic(err)
	}

	numbers := parseNumbers(strings.Split(string(content), "\n")[0])
	ranges := h.Map[[]int, tRange](h.InGroupsOf[int](numbers, 2), func(a []int) tRange {
		return tRange{min: a[0], max: a[1]}
	})
	maps := parseInput(strings.Split(string(content), "\n")[1:])

	for _, m := range maps {
		numbers = m.translateNumbers(numbers)
	}

	for _, m := range maps {
		ranges = m.translateRanges(ranges)
	}

	log.Println("Part 1:", h.Min(numbers))
	log.Println("Part 2:", h.Min(h.Map[tRange, int](ranges, func(r tRange) int { return r.min })))
}

func parseNumbers(input string) []int {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	return h.Map[[]string, int](matches, func(match []string) int {
		return h.Must[int](strconv.Atoi(match[0]))
	})
}

func parseInput(input []string) []*tMap {
	var result []*tMap

	for _, s := range input {
		if s == "" {
			continue
		}

		if strings.Contains(s, "map") {
			result = append(result, &tMap{})
			continue
		}

		last := h.Last(result)
		last.AddRange(parseRange(s))
	}

	return result
}

func parseRange(input string) tRange {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	to := h.Must[int](strconv.Atoi(matches[0][0]))
	min := h.Must[int](strconv.Atoi(matches[1][0]))
	length := h.Must[int](strconv.Atoi(matches[2][0]))

	return tRange{
		to:  to,
		min: min,
		max: min + length - 1,
	}
}

type tRange struct {
	min, max, to int
}

func (r tRange) include(a int) bool {
	return a >= r.min && a <= r.max
}

func (r tRange) translate(a int) int {
	if r.include(a) {
		return r.to + (a - r.min)
	}

	panic("can not translate")
}

type tMap struct {
	ranges []tRange
}

func (m *tMap) AddRange(r tRange) {
	m.ranges = append(m.ranges, r)
}

func (m *tMap) translate(a int) int {
	for _, r := range m.ranges {
		if r.include(a) {
			return r.translate(a)
		}
	}

	return a
}

func (m *tMap) translateNumbers(values []int) []int {
	return h.Map[int](values, func(a int) int {
		return m.translate(a)
	})
}

func (m *tMap) translateRanges(values []tRange) []tRange {
	return values
}
