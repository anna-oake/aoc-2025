package main

import (
	"strconv"
	"strings"
)

func (*methods) D3P1(input string) string {
	banks := strings.Split(input, "\n")
	var total int
	for _, bank := range banks {
		var maxj, secj int
		var maxidx, secidx int
		for i, cell := range bank {
			j := int(cell) - 48
			if j > maxj && i < len(bank)-1 {
				maxj = j
				maxidx = i
			}
		}
		for i := maxidx + 1; i < len(bank); i++ {
			j := int(bank[i]) - 48
			if j > secj {
				secj = j
				secidx = i
			}
		}
		if maxidx > secidx {
			total += secj*10 + maxj
		} else {
			total += maxj*10 + secj
		}
	}
	return strconv.Itoa(total)
}

func (*methods) D3P2(input string) string {
	banks := strings.Split(input, "\n")
	var total int64
	for _, bank := range banks {
		var on []int
		maxCells := 12
		var startidx int
		for {
			var maxj int
			stopidx := len(bank) - maxCells
			for i := startidx; i <= stopidx; i++ {
				j := int(bank[i]) - 48
				if j > maxj {
					maxj = j
					startidx = i + 1
				}
			}
			on = append(on, maxj)
			maxCells--
			if maxCells == 0 {
				break
			}
		}
		var jS string
		for _, j := range on {
			jS += strconv.Itoa(j)
		}
		jolts, _ := strconv.ParseInt(jS, 10, 64)
		total += jolts
	}
	return strconv.FormatInt(total, 10)
}
