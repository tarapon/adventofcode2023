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
	res1 = graph.perimeter() / 2
	res2 = graph.square() - res1 + 1

	log.Println("calibrate (simple):", res1)
	log.Println("calibrate (advanced):", res2)
}

type vertex struct {
	i, j int
}

type node struct {
	i, j int
	next *node
	prev *node
}

func (n *node) vertex() vertex {
	return vertex{i: n.i, j: n.j}
}

func (n *node) vertices() []vertex {
	res := make([]vertex, 0)

	cur := n
	for cur.next != n {
		res = append(res, cur.vertex())
		cur = cur.next
	}

	res = append(res, cur.vertex())

	return res
}

func (n *node) square() int {
	s1, s2 := 0, 0
	v := n.vertices()

	for k := 0; k < len(v)-1; k++ {
		s1 += v[k].i * v[k+1].j
		s2 += v[k+1].i * v[k].j
	}

	s1 += v[len(v)-1].i * v[0].j
	s2 += v[0].i * v[len(v)-1].j

	if s1 > s2 {
		return (s1 - s2) / 2
	} else {
		return (s2 - s1) / 2
	}
}

func (n *node) perimeter() int {
	cur := n
	k := 1

	for cur.next != n {
		cur = cur.next
		k++
	}

	return k
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
