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

	res1, res2 := 0, 0

	ch1 := extractNumbers(strings.Split(string(content), "\n"))
	ch2 := extractNumberPairs(strings.Split(string(content), "\n"))

	for num := range ch1 {
		res1 += num
	}

	for num := range ch2 {
		res2 += num
	}

	log.Println("Parts (simple):", res1)
	log.Println("Parts (advanced):", res2)
}

func mustParseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return val
}

func isSymbol(c int32) bool {
	return c != '.' && !isDigit(c)
}

func isDigit(c int32) bool {
	return c >= '0' && c <= '9'
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

func findNumbersAround(line string, pos int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		if isDigit(rune(line[pos])) {
			k1, k2 := pos, pos

			for k1 >= 0 && isDigit(rune(line[k1])) {
				k1--
			}

			for k2 < len(line) && isDigit(rune(line[k2])) {
				k2++
			}

			ch <- mustParseInt(line[k1+1 : k2])
			return
		}

		if pos > 0 && isDigit(rune(line[pos-1])) {
			for num := range findNumbersAround(line, pos-1) {
				ch <- num
			}
		}

		if pos < len(line)-1 && isDigit(rune(line[pos+1])) {
			for num := range findNumbersAround(line, pos+1) {
				ch <- num
			}
		}

	}()

	return ch
}

func extractNumberPairs(data []string) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i, line := range data {
			for j, c := range line {
				if !isSymbol(c) {
					continue
				}

				var nums []int

				for k := i - 1; k <= i+1; k++ {
					if k < 0 || k >= len(data) {
						continue
					}

					for n := range findNumbersAround(data[k], j) {
						nums = append(nums, n)
					}
				}

				if len(nums) == 2 {
					ch <- nums[0] * nums[1]
				}
			}
		}
	}()

	return ch
}
