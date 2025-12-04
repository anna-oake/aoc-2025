package main

import (
	"strconv"
	"strings"
)

func (*methods) D4P1(input string) string {
	rows := strings.Split(input, "\n")
	w := len(rows[0])
	h := len(rows)
	var accessible int
	for y, row := range rows {
		for x := range w {
			if row[x] != '@' {
				continue
			}
			c := coords32{
				x: x,
				y: y,
			}
			var cnt int
			for d := range 8 {
				nc := c.move(d)
				if nc.inBounds(w, h) && rows[nc.y][nc.x] == '@' {
					cnt++
				}
			}
			if cnt < 4 {
				accessible++
			}
		}
	}
	return strconv.Itoa(accessible)
}

func (*methods) D4P2(input string) string {
	rows := strings.Split(input, "\n")
	w := len(rows[0])
	h := len(rows)
	var total int
	for {
		var newrows []string
		var accessible int
		for y, row := range rows {
			newrow := []rune(row)
			for x := range w {
				if row[x] != '@' {
					continue
				}
				c := coords32{
					x: x,
					y: y,
				}
				var cnt int
				for d := range 8 {
					nc := c.move(d)
					if nc.inBounds(w, h) && rows[nc.y][nc.x] == '@' {
						cnt++
					}
				}
				if cnt < 4 {
					newrow[x] = '.'
					accessible++
				}
			}
			newrows = append(newrows, string(newrow))
		}
		rows = newrows
		total += accessible
		if accessible == 0 {
			break
		}
	}
	return strconv.Itoa(total)
}
