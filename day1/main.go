package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}

	res1, res2 := 0, 0
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		res1 += extractNumber(line, false)
		res2 += extractNumber(line, true)
	}

	log.Println("calibrate (simple):", res1)
	log.Println("calibrate (advanced):", res2)
}

func extractNumber(s string, advanced bool) int {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// one, two, thee, four, five, six, seven, eight, nine

	first, last := -1, -1
	for idx, c := range s {
		if c >= '0' && c <= '9' {
			if first == -1 {
				first = int(c) - '0'
			}

			last = int(c) - '0'
		} else if advanced {
			for i, w := range words {
				if strings.HasPrefix(s[idx:], w) {
					if first == -1 {
						first = i + 1
					}

					last = i + 1
					break
				}
			}
		}
	}

	return first*10 + last
}
