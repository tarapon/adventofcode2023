package main

import (
	h "advcode2025/helper"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("day10/input.txt")
	if err != nil {
		panic(err)
	}

	res1, res2 := 0, 0
	lines := strings.Split(string(content), "\n")
	m := parseInput(lines)

	graph := toList(m)
	res1 = longestPath(graph)

	log.Println("calibrate (simple):", res1)
	log.Println("calibrate (advanced):", res2)
}

func longestPath(head *node) int {
	k := 1
	prev, next := head.prev, head.next
	for prev != next {
		k++
		prev, next = prev.prev, next.next
	}

	return k
}

type node struct {
	i, j int
	next *node
	prev *node
}

func parseInput(s []string) [][]byte {
	return h.Map(s, func(s string) []byte { return []byte(s) })
}

func findStart(m [][]byte) (int, int) {
	for i, row := range m {
		for j, val := range row {
			if val == 'S' {
				return i, j
			}
		}
	}

	panic("no start found")
}

func toList(m [][]byte) *node {
	head := &node{}

	head.i, head.j = findStart(m)

	// determine next node
	if head.i > 0 {
		if h.Contains([]byte{'|', '7', 'F'}, m[head.i-1][head.j]) {
			head.next = &node{i: head.i - 1, j: head.j, prev: head}
		}
	}

	if head.next == nil && head.i+1 < len(m) {
		if h.Contains([]byte{'|', 'L', 'J'}, m[head.i+1][head.j]) {
			head.next = &node{i: head.i + 1, j: head.j, prev: head}
		}
	}

	if head.next == nil && head.j > 0 {
		if h.Contains([]byte{'-', 'L', 'F'}, m[head.i][head.j-1]) {
			head.next = &node{i: head.i, j: head.j - 1, prev: head}
		}
	}

	if head.next == nil && head.j+1 < len(m[head.i]) {
		if h.Contains([]byte{'-', 'J', '7'}, m[head.i][head.j+1]) {
			head.next = &node{i: head.i, j: head.j + 1, prev: head}
		}
	}

	cur := head.next

	// follow path until we reach the head again
	for cur.i != head.i || cur.j != head.j {
		switch m[cur.i][cur.j] {
		case '|':
			if cur.prev.i < cur.i {
				cur.next = &node{i: cur.i + 1, j: cur.j, prev: cur}
			} else {
				cur.next = &node{i: cur.i - 1, j: cur.j, prev: cur}
			}
		case '-':
			if cur.prev.j < cur.j {
				cur.next = &node{i: cur.i, j: cur.j + 1, prev: cur}
			} else {
				cur.next = &node{i: cur.i, j: cur.j - 1, prev: cur}
			}
		case 'L':
			if cur.prev.i < cur.i {
				cur.next = &node{i: cur.i, j: cur.j + 1, prev: cur}
			} else {
				cur.next = &node{i: cur.i - 1, j: cur.j, prev: cur}
			}
		case 'J':
			if cur.prev.i < cur.i {
				cur.next = &node{i: cur.i, j: cur.j - 1, prev: cur}
			} else {
				cur.next = &node{i: cur.i - 1, j: cur.j, prev: cur}
			}
		case '7':
			if cur.prev.i == cur.i {
				cur.next = &node{i: cur.i + 1, j: cur.j, prev: cur}
			} else {
				cur.next = &node{i: cur.i, j: cur.j - 1, prev: cur}
			}
		case 'F':
			if cur.prev.i == cur.i {
				cur.next = &node{i: cur.i + 1, j: cur.j, prev: cur}
			} else {
				cur.next = &node{i: cur.i, j: cur.j + 1, prev: cur}
			}
		default:
			panic("unknown direction")
		}

		cur = cur.next
	}

	cur.prev.next, head.prev = head, cur.prev

	return head
}
