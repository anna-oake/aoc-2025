package main

import (
	"strconv"
	"strings"
)

// i really tried to be super smart but eventually gave up. sorry.

func (*methods) D2P1(input string) string {
	ranges := strings.Split(input, ",")
	var invalid int64
	for _, r := range ranges {
		r = strings.TrimSpace(r)
		ends := strings.Split(r, "-")
		from, _ := strconv.ParseInt(ends[0], 10, 64)
		to, _ := strconv.ParseInt(ends[1], 10, 64)
		for i := from; i <= to; i++ {
			s := strconv.FormatInt(i, 10)
			if len(s)%2 != 0 {
				continue
			}
			if s[:len(s)/2] == s[len(s)/2:] {
				invalid += i
			}
		}
	}
	return strconv.FormatInt(invalid, 10)
}

func (*methods) D2P2(input string) string {
	ranges := strings.Split(input, ",")
	var total int64
	for _, r := range ranges {
		r = strings.TrimSpace(r)
		ends := strings.Split(r, "-")
		from, _ := strconv.ParseInt(ends[0], 10, 64)
		to, _ := strconv.ParseInt(ends[1], 10, 64)
		for i := from; i <= to; i++ {
			if i < 10 {
				continue
			}
			s := strconv.FormatInt(i, 10)
			invalid := true
			for n := 1; n <= len(s); n++ {
				invalid = true
				if len(s)%n == 0 {
					l := len(s) / n
					if n == 1 {
						l = 1
					}
					pat := s[:l]
					for j := l; j <= len(s)-l; j += l {
						if pat != s[j:j+l] {
							invalid = false
							break
						}
					}
					if invalid {
						break
					}
				}
			}
			if invalid {
				total += i
			}
		}
	}
	return strconv.FormatInt(total, 10)
}
