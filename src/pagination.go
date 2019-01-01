package main

import (
	"fmt"
	"strconv"
	"time"
)

type nav struct {
	current, start, end int
}

func (n *nav) prev() int {
	if n.current > n.start {
		return n.current - 1
	}
	return 0
}
func (n *nav) next() int {
	if n.current < n.end {
		return n.current + 1
	}
	return 0
}
func NewNav() *nav {
	n := new(nav)
	year, week := time.Now().ISOWeek()
	current := fmt.Sprintf(
		"%d%d",
		year,
		week,
	)
	n.start = 20191
	n.end, _ = strconv.Atoi(current)
	return n
}

func (n *nav) setCurrent(date string) {
	n.current, _ = strconv.Atoi(date)
}
