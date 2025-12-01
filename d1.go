package main

import (
	"strconv"
	"strings"
)

func (*methods) D1P1(input string) string {
	lines := strings.Split(input, "\n")
	pos := 50
	var cnt int
	for _, line := range lines {
		line = strings.TrimPrefix(line, "R")
		line = strings.ReplaceAll(line, "L", "-")
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		pos += num
		pos %= 100
		if pos > 99 {
			pos = pos - 100
		} else if pos < 0 {
			pos = pos + 100
		}
		if pos == 0 {
			cnt++
		}
	}

	return strconv.Itoa(cnt)
}
