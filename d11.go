package main

import (
	"slices"
	"strconv"
	"strings"
)

type d11device struct {
	ID          string
	Connections []string
}

var d11devs = make(map[string][]string)
var d11cache = make(map[string]int64)

func (*methods) D11P1(input string) string {
	rows := strings.Split(input, "\n")
	for _, r := range rows {
		parts := strings.Split(r, ": ")
		id := parts[0]
		conns := strings.Split(parts[1], " ")
		d11devs[id] = conns
	}

	return strconv.FormatInt(d11countPaths("you", "out"), 10)
}

func (*methods) D11P2(input string) string {
	rows := strings.Split(input, "\n")
	for _, r := range rows {
		parts := strings.Split(r, ": ")
		id := parts[0]
		conns := strings.Split(parts[1], " ")
		d11devs[id] = conns
	}

	return strconv.FormatInt(d11countPaths2("svr", "out", false, false), 10)
}

func d11countPaths(from, to string) int64 {
	conns, ok := d11devs[from]
	if !ok {
		return 0
	}
	if slices.Contains(conns, to) {
		return 1
	}
	var count int64
	for _, c := range conns {
		count += d11countPaths(c, to)
	}
	return count
}

func d11countPaths2(from, to string, dac, fft bool) int64 {
	if from == "dac" {
		dac = true
	}
	if from == "fft" {
		fft = true
	}

	key := from + to + strconv.FormatBool(dac) + strconv.FormatBool(fft)

	mem, ok := d11cache[key]
	if ok {
		return mem
	}
	conns, ok := d11devs[from]
	if !ok {
		d11cache[key] = 0
		return 0
	}
	if slices.Contains(conns, to) {
		if dac && fft {
			d11cache[key] = 1
			return 1
		} else {
			d11cache[key] = 0
			return 0
		}
	}
	var count int64
	for _, c := range conns {
		count += d11countPaths2(c, to, dac, fft)
	}
	d11cache[key] = count
	return count
}
