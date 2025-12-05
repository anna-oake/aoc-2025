package main

import (
	"sort"
	"strconv"
	"strings"
)

var d5freshFrom, d5freshTo []uint64

func (*methods) D5P1(input string) string {
	rows := strings.Split(input, "\n")
	var halfway bool
	var cnt int
	for _, r := range rows {
		if r == "" {
			halfway = true
			continue
		}
		if !halfway {
			parts := strings.Split(r, "-")
			from, _ := strconv.ParseUint(parts[0], 10, 64)
			to, _ := strconv.ParseUint(parts[1], 10, 64)
			d5freshFrom = append(d5freshFrom, from)
			d5freshTo = append(d5freshTo, to)
			continue
		}
		num, _ := strconv.ParseUint(r, 10, 64)
		for i, from := range d5freshFrom {
			if num >= from && num <= d5freshTo[i] {
				cnt++
				break
			}
		}
	}
	return strconv.Itoa(cnt)
}

type d5range struct {
	from uint64
	to   uint64
}

var d5ranges []*d5range

func (*methods) D5P2(input string) string {
	rows := strings.Split(input, "\n")
	for _, r := range rows {
		if r == "" {
			break
		}
		parts := strings.Split(r, "-")
		from, _ := strconv.ParseUint(parts[0], 10, 64)
		to, _ := strconv.ParseUint(parts[1], 10, 64)
		d5ranges = append(d5ranges, &d5range{from, to})
	}

	sort.Slice(d5ranges, func(i, j int) bool {
		return d5ranges[i].from < d5ranges[j].from
	})

	for i := range len(d5ranges) {
		r1 := d5ranges[i]
		if r1 == nil {
			continue
		}
		for j := i + 1; j < len(d5ranges); j++ {
			r2 := d5ranges[j]
			if r2 == nil {
				continue
			}
			if r2.from >= r1.from && r2.from <= r1.to {
				if r2.to >= r1.to {
					r1.to = r2.to
				}
				d5ranges[j] = nil
			}
		}
	}

	var cnt uint64

	for _, r := range d5ranges {
		if r != nil {
			cnt += r.to - r.from + 1
		}
	}

	return strconv.FormatUint(cnt, 10)
}
