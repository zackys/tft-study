package main

import (
	"fmt"
)

////

type AddlnFilter struct {
	Filter

	start int
	step  int

	current int
}

func NewAddlnFilter(start, step int) *AddlnFilter {
	fltr := &AddlnFilter{
		start: start,
		step:  step,

		current: start,
	}

	fltr.Apply = func(line string) string {
		ret := fmt.Sprintf("%06d%s", fltr.current, line)
		fltr.current += fltr.step
		return ret
	}
	return fltr
}

////

type RmlnFilter struct {
	Filter

	col int
}

func NewRmlnFilter(col int) *RmlnFilter {
	fltr := &RmlnFilter{
		col: col,
	}

	fltr.Apply = func(line string) string {
		ret := ""
		if len(line) > fltr.col {
			ret = line[fltr.col:]
		}
		return ret
	}
	return fltr
}