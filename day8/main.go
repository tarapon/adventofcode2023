package main

import (
	h "advcode2025/helper"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	content, err := os.ReadFile("day8/input.txt")
	if err != nil {
		panic(err)
	}

	res1, res2 := 0, 0
	lines := strings.Split(string(content), "\n")

	steps := tSteps{steps: strings.Split(lines[0], "")}
	mp := parseMap(lines[2:])

	for mp.cur != "ZZZ" {
		res1++
		mp.move(steps.next())
	}

	maps := make([]*tMap, 0)
	for key, _ := range mp.data {
		if strings.HasSuffix(key, "A") {
			maps = append(maps, &tMap{
				cur:  key,
				data: mp.data,
			})
		}
	}

	steps.reset()
	allZeds := false
	values := make([]int, len(maps))

	for !allZeds {
		res2++
		step := steps.next()
		for i, m := range maps {
			m.move(step)
			if values[i] == 0 && strings.HasSuffix(m.cur, "Z") {
				values[i] = res2
			}
		}

		allZeds = h.All(values, func(v int) bool { return v > 0 })
	}

	res2 = lcm(values[0], values[1], values[2:]...)

	log.Println("distance (simple):", res1)
	log.Println("distance (advanced):", res2)
}

type tSteps struct {
	steps []string
	idx   int
}

func (s *tSteps) next() string {
	if s.idx >= len(s.steps) {
		s.idx = 0
	}
	s.idx++
	return s.steps[s.idx-1]
}

func (s *tSteps) reset() {
	s.idx = 0
}

type tMap struct {
	data map[string][]string
	cur  string
}

func (m *tMap) move(dir string) {
	if dir == "L" {
		m.cur = m.data[m.cur][0]
	} else if dir == "R" {
		m.cur = m.data[m.cur][1]
	} else {
		panic("unknown direction")
	}
}

func parseMap(lines []string) tMap {
	m := tMap{cur: "AAA"}
	m.data = make(map[string][]string)

	re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		m.data[match[1]] = []string{match[2], match[3]}
	}

	return m
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
