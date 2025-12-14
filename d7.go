package main

import (
	"strconv"
	"strings"
)

var d7splitters = make(map[int]bool)
var d7splits int
var d7visited = make(map[int]bool)
var d7counts = make(map[int]int)
var d7w, d7h int

func (*methods) D7P1(input string) string {
	var start coords

	rows := strings.Split(input, "\n")

	d7w = len(rows[0])
	d7h = len(rows)

	for y, r := range rows {
		for x, c := range r {
			p := coords{x, y}
			if c == 'S' {
				start = p
			}
			if c == '^' {
				d7splitters[p.getIdx(d7w)] = true
			}
		}
	}

	d7recurse(start)

	return strconv.Itoa(d7splits)
}

func (*methods) D7P2(input string) string {
	var start coords

	rows := strings.Split(input, "\n")

	d7w = len(rows[0])
	d7h = len(rows)

	for y, r := range rows {
		for x, c := range r {
			p := coords{x, y}
			if c == 'S' {
				start = p
			}
			if c == '^' {
				d7splitters[p.getIdx(d7w)] = true
			}
		}
	}

	return strconv.Itoa(d7count(start))
}

func d7recurse(now coords) {
	for {
		now = now.move(DOWN)
		if !now.inBounds(d7w, d7h) {
			return
		}
		if !d7splitters[now.getIdx(d7w)] {
			continue
		}
		if d7visited[now.getIdx(d7w)] {
			return
		}
		d7visited[now.getIdx(d7w)] = true

		d7splits++

		d7recurse(now.move(LEFT))
		d7recurse(now.move(RIGHT))
		return
	}
}

func d7count(now coords) int {
	for {
		now = now.move(DOWN)
		if !now.inBounds(d7w, d7h) {
			return 1
		}
		if !d7splitters[now.getIdx(d7w)] {
			continue
		}
		mem, ok := d7counts[now.getIdx(d7w)]
		if ok {
			return mem
		}

		left := now.move(LEFT)
		right := now.move(RIGHT)

		lc, ok := d7counts[left.getIdx(d7w)]
		if !ok {
			lc = d7count(left)
			d7counts[left.getIdx(d7w)] = lc
		}
		rc, ok := d7counts[right.getIdx(d7w)]
		if !ok {
			rc = d7count(right)
			d7counts[right.getIdx(d7w)] = rc
		}

		d7counts[now.getIdx(d7w)] = lc + rc
		return lc + rc
	}
}
