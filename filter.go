package main

import (

)

////

type GoFilterer interface {
	GoFilter(in <-chan string) <-chan string
}

type Filter struct {
	Apply func(string) string
}

func (fltr *Filter) GoFilter(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for str := range in {
			line := fltr.Apply(str)
			//println(line)
			out <- line
			//out <- cmd.Filter(str)
		}
		close(out)
	}()
	return out
}
