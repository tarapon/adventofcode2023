package main

import (
	h "advcode2025/helper"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("day11/input.txt")
	if err != nil {
		panic(err)
	}

	res1, res2 := 0, 0
	lines := strings.Split(string(content), "\n")
	m := tMap(parseInput(lines))
	expand(m)

	galaxies := findGalaxies(m)

	forEachPair(galaxies, func(g1, g2 galaxy) {
		res1 += m.distance(g1, g2, 2)
		res2 += m.distance(g1, g2, 1000000)
	})

	log.Println("distance (simple):", res1)
	log.Println("distance (advanced):", res2)
}

type tMap [][]byte

type galaxy struct {
	i, j int
}

func (m tMap) distance(a, b galaxy, w int) int {
	hDist, vDist := 0, 0

	for i := minInt(a.i, b.i); i < maxInt(a.i, b.i); i++ {
		if m[i][0] == '+' {
			vDist += w
		} else {
			vDist++
		}
	}

	for j := minInt(a.j, b.j); j < maxInt(a.j, b.j); j++ {
		if m[0][j] == '+' {
			hDist += w
		} else {
			hDist++
		}
	}

	return hDist + vDist
}

func forEachPair(galaxies []galaxy, f func(g1, g2 galaxy)) {
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			f(galaxies[i], galaxies[j])
		}
	}
}

func parseInput(lines []string) [][]byte {
	return h.Map(lines, func(line string) []byte {
		return []byte(line)
	})
}

func expand(s [][]byte) {
	// expand vertically
	for i := 0; i < len(s); i++ {
		allDots := true

		for j := 0; j < len(s[i]); j++ {
			if s[i][j] != '.' {
				allDots = false
				break
			}
		}

		if allDots {
			for j := 0; j < len(s[i]); j++ {
				s[i][j] = '+'
			}
		}
	}

	// expand horizontally
	for j := 0; j < len(s[0])-1; j++ {
		allDots := true

		for i := 0; i < len(s); i++ {
			if s[i][j] != '.' && s[i][j] != '+' {
				allDots = false
				break
			}
		}

		if allDots {
			for i := 0; i < len(s); i++ {
				s[i][j] = '+'
			}
		}
	}
}

func findGalaxies(m [][]byte) []galaxy {
	galaxies := make([]galaxy, 0)

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == '#' {
				galaxies = append(galaxies, galaxy{i, j})
			}
		}
	}

	return galaxies
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
