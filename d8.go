package main

import (
	"sort"
	"strconv"
	"strings"
)

type d8pair struct {
	box1         string
	box2         string
	distance     int
	wallDistance int
}

func (*methods) D8P1(input string) string {
	var circuits []*map[string]bool
	connections := make(map[string]*map[string]bool)

	rows := strings.Split(input, "\n")
	var points []coords3d
	for _, r := range rows {
		p := strings.Split(r, ",")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])
		z, _ := strconv.Atoi(p[2])
		points = append(points, coords3d{x, y, z})
	}

	var pairs []d8pair
	for i, p := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			pairs = append(pairs, d8pair{p.String(), p2.String(), p.distance(p2), p.x * p2.x})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	for i, pair := range pairs {
		if i == 10 {
			break
		}
		c1, _ := connections[pair.box1]
		c2, _ := connections[pair.box2]
		var c *map[string]bool
		if c1 == nil && c2 != nil {
			c = c2
		} else if c1 != nil && c2 == nil {
			c = c1
		} else if c1 == nil && c2 == nil {
			c = &map[string]bool{}
			circuits = append(circuits, c)
		} else {
			if c1 == c2 {
				continue
			} else {
				c = c1
				for k := range *c2 {
					(*c)[k] = true
					connections[k] = c
				}
				*c2 = nil
			}
		}
		(*c)[pair.box1] = true
		(*c)[pair.box2] = true
		connections[pair.box1] = c
		connections[pair.box2] = c
	}

	var t1, t2, t3 int
	for _, c := range circuits {
		if c == nil {
			continue
		}
		l := len(*c)
		if l > t1 {
			t3 = t2
			t2 = t1
			t1 = l
		} else if l > t2 {
			t3 = t2
			t2 = l
		} else if l > t3 {
			t3 = l
		}
	}

	return strconv.Itoa(t1 * t2 * t3)
}

func (*methods) D8P2(input string) string {
	stray := make(map[string]bool)
	var circuits []*map[string]bool
	connections := make(map[string]*map[string]bool)

	rows := strings.Split(input, "\n")
	var points []coords3d
	for _, r := range rows {
		p := strings.Split(r, ",")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])
		z, _ := strconv.Atoi(p[2])
		points = append(points, coords3d{x, y, z})
		stray[r] = true
	}

	var pairs []d8pair
	for i, p := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			pairs = append(pairs, d8pair{p.String(), p2.String(), p.distance(p2), p.x * p2.x})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	var lastPair d8pair

	for _, pair := range pairs {
		c1, _ := connections[pair.box1]
		c2, _ := connections[pair.box2]
		var c *map[string]bool
		if c1 == nil && c2 != nil {
			c = c2
		} else if c1 != nil && c2 == nil {
			c = c1
		} else if c1 == nil && c2 == nil {
			c = &map[string]bool{}
			circuits = append(circuits, c)
		} else {
			if c1 == c2 {
				continue
			} else {
				c = c1
				for k := range *c2 {
					(*c)[k] = true
					connections[k] = c
				}
				*c2 = nil
			}
		}
		(*c)[pair.box1] = true
		(*c)[pair.box2] = true
		connections[pair.box1] = c
		connections[pair.box2] = c
		delete(stray, pair.box1)
		delete(stray, pair.box2)
		if len(stray) == 0 {
			lastPair = pair
			break
		}
	}

	return strconv.Itoa(lastPair.wallDistance)
}
