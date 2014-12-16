package cnv

import (
)

type Filter interface {
	GoFilter(in <-chan string) <-chan string
}

type FilterCmd struct {
	Filter func(string) string
}

func (cmd FilterCmd) GoFilter(in <-chan string) <-chan string {
	out := make(chan string)
	println("*1")
	go func() {
		for str := range in {
			line := cmd.Filter(str)
			//println(line)
			out <- line
			//out <- cmd.Filter(str)
		}
		close(out)
	}()
	return out
}