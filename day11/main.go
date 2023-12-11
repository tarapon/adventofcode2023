package main

import (
	h "advcode2025/helper"
	"log"
	"math"
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
	m := parseInput(lines)

	galaxies := findGalaxies(expand(m))

	forEachPair(galaxies, func(g1, g2 galaxy) {
		res1 += g1.distanceTo(g2)
	})

	log.Println("distance (simple):", res1)
	log.Println("distance (advanced):", res2)
}

type galaxy struct {
	i, j int
}

func forEachPair(galaxies []galaxy, f func(g1, g2 galaxy)) {
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			f(galaxies[i], galaxies[j])
		}
	}
}

func (g galaxy) distanceTo(o galaxy) int {
	return int(math.Abs(float64(g.i-o.i)) + math.Abs(float64(g.j-o.j)))
}

func parseInput(lines []string) [][]byte {
	return h.Map(lines, func(line string) []byte {
		return []byte(line)
	})
}

func insertAt(a []byte, idx int, v byte) []byte {
	a = append(a, 0)
	copy(a[idx+1:], a[idx:])
	a[idx] = v
	return a
}

func expand(s [][]byte) [][]byte {
	d := make([][]byte, 0, len(s))

	// expand vertically
	for i := 0; i < len(s); i++ {
		d = append(d, s[i])

		if h.All(s[i], func(b byte) bool { return b == '.' }) {
			d = append(d, s[i])
		}
	}

	// expand horizontally
	for j := len(d[0]) - 1; j >= 0; j-- {
		allDots := true

		for i := 0; i < len(d); i++ {
			if d[i][j] != '.' {
				allDots = false
				break
			}
		}

		if allDots {
			for i := 0; i < len(d); i++ {
				d[i] = insertAt(d[i], j, '.')
			}
		}
	}

	return d
}

func findGalaxies(m [][]byte) []galaxy {
	galaxies := make([]galaxy, 0)

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == '#' {
				galaxies = append(galaxies, galaxy{j, i})
			}
		}
	}

	return galaxies
}
