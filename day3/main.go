package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	res1 := 0

	ch := extractNumbers(strings.Split(string(content), "\n"))

	for num := range ch {
		res1 += num
	}

	log.Println("Parts sum:", res1)
}

func mustParseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return val
}

func isSymbol(c int32) bool {
	return c != '.' && !(c >= '0' && c <= '9')
}

func hasSymbol(s string) bool {
	for _, c := range s {
		if isSymbol(c) {
			return true
		}
	}

	return false
}

func extractNumbers(data []string) <-chan int {
	re := regexp.MustCompile(`\d+`)
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i, line := range data {
			matches := re.FindAllSubmatchIndex([]byte(line), -1)
			for _, match := range matches {
				i1, i2 := match[0], match[1]
				isAdjacent := false

				for j := i - 1; j <= i+1 && !isAdjacent; j++ {
					if j < 0 || j >= len(line) {
						continue
					}

					k1, k2 := i1, i2

					if k1 > 0 {
						k1--
					}

					if k2 < len(line) {
						k2++
					}

					isAdjacent = isAdjacent || hasSymbol(data[j][k1:k2])
				}

				if isAdjacent {
					ch <- mustParseInt(line[i1:i2])
				}
			}
		}
	}()

	return ch
}
