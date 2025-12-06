package main

import (
	"strconv"
	"strings"
)

func (*methods) D6P1(input string) string {
	rows := strings.Split(input, "\n")
	var problems [][]int
	var ops []bool
	for _, r := range rows {
		columns := strings.Split(r, " ")
		var i int
		for _, c := range columns {
			c = strings.TrimSpace(c)
			if c == "" {
				continue
			}
			if len(problems) <= i {
				problems = append(problems, []int{})
			}
			if c == "*" {
				ops = append(ops, true)
			} else if c == "+" {
				ops = append(ops, false)
			} else {
				num, _ := strconv.Atoi(c)
				problems[i] = append(problems[i], num)
			}
			i++
		}
	}
	var total int64
	for i, p := range problems {
		op := ops[i]
		result := int64(p[0])
		for j, num := range p {
			if j == 0 {
				continue
			}
			if op {
				result *= int64(num)
			} else {
				result += int64(num)
			}
		}
		total += result
	}
	return strconv.FormatInt(total, 10)
}

func (*methods) D6P2(input string) string {
	rows := strings.Split(input, "\n")
	w := len(rows[0]) - 1
	var problems [][]int64
	var ops []bool
	var done bool
	var problem []int64
	for i := w; i >= 0; i-- {
		var col string
		for _, r := range rows {
			if len(r) <= i {
				break
			}
			cell := strings.TrimSpace(string(r[i]))
			if cell == "" {
				continue
			}
			if cell == "*" {
				ops = append(ops, true)
				done = true
			} else if cell == "+" {
				ops = append(ops, false)
				done = true
			} else {
				col += cell
			}
		}
		num, _ := strconv.ParseInt(col, 10, 64)
		if num == 0 {
			continue
		}
		problem = append(problem, num)
		if done {
			problems = append(problems, problem)
			done = false
			problem = nil
		}
	}

	var total int64
	for i, p := range problems {
		op := ops[i]
		result := int64(p[0])
		for j, num := range p {
			if j == 0 {
				continue
			}
			if op {
				result *= int64(num)
			} else {
				result += int64(num)
			}
		}
		total += result
	}
	return strconv.FormatInt(total, 10)
}
