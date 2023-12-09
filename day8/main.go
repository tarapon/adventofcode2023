package main

import (
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
