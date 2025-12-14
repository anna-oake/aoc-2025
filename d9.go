package main

import (
	"strconv"
	"strings"
)

func (*methods) D9P1(input string) string {
	rows := strings.Split(input, "\n")
	var tiles []coords32
	for _, r := range rows {
		pair := strings.Split(r, ",")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		tiles = append(tiles, coords32{x: x, y: y})
	}
	var maxArea int
	for i, t := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			t2 := tiles[j]
			area := (abs(t2.x-t.x) + 1) * (abs(t2.y-t.y) + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return strconv.Itoa(maxArea)
}
