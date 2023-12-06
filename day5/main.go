package main

import (
	h "advcode2025/helper"
	"golang.org/x/exp/slices"
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
		return tRange{min: a[0], max: a[0] + a[1] - 1}
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

func (r tRange) empty() bool {
	return r.min > r.max
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

	slices.SortStableFunc(m.ranges, func(a, b tRange) int {
		return a.min - b.min
	})
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

func (m *tMap) split(value tRange) []tRange {
	result := make([]tRange, 0, 1)

	for _, r := range m.ranges {
		if value.empty() || r.min > value.max {
			break
		}

		if r.max < value.min {
			continue
		}

		if value.min < r.min {
			result = append(result, tRange{min: value.min, max: r.min - 1})
			if value.max <= r.max {
				result = append(result, tRange{min: r.min, max: value.max})
				value.min = value.max + 1
			} else {
				result = append(result, tRange{min: r.min, max: r.max})
				value.min = r.max + 1
			}
		} else {
			if value.max <= r.max {
				result = append(result, tRange{min: value.min, max: value.max})
				value.min = value.max + 1
			} else {
				result = append(result, tRange{min: value.min, max: r.max})
				value.min = r.max + 1
			}
		}
	}

	if !value.empty() {
		result = append(result, value)
	}

	return result
}

func (m *tMap) translateRanges(values []tRange) []tRange {
	result := make([]tRange, 0, len(values))

	for _, r := range values {
		for _, r1 := range m.split(r) {
			result = append(result, tRange{
				min: m.translate(r1.min),
				max: m.translate(r1.max),
			})
		}
	}

	return result
}
