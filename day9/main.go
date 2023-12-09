package main

import (
	h "advcode2025/helper"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("day9/input.txt")
	if err != nil {
		panic(err)
	}

	res1, res2 := 0, 0
	lines := strings.Split(string(content), "\n")

	sequences := h.Map[string, []int](lines, func(s string) []int {
		return h.Map[string, int](strings.Split(s, " "), func(s string) int {
			return h.Must(strconv.Atoi(s))
		})
	})

	for _, seq := range sequences {
		value := predictNextNumber(seq)
		res1 += value
	}

	log.Println("calibrate (simple):", res1)
	log.Println("calibrate (advanced):", res2)
}

func diffs(numbers []int) []int {
	res := make([]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		res[i] = numbers[i+1] - numbers[i]
	}
	return res
}

func zeros(numbers []int) bool {
	return h.All(numbers, func(n int) bool { return n == 0 })
}

func predictNextNumber(numbers []int) int {
	lasts := make([]int, 0)

	for !zeros(numbers) {
		lasts = append(lasts, h.Last(numbers))
		numbers = diffs(numbers)
	}

	sum := 0
	for _, n := range lasts {
		sum += n
	}

	return sum
}
