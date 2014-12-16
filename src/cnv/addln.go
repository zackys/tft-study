package cnv

import (
	"fmt"
)

type Addln struct {
	FilterCmd

	start int
	step  int

	current int
}

func NewAddln(start, step int) *Addln {
	cmd := &Addln{
		start: start,
		step:  step,
	}

	cmd.current = cmd.start
	cmd.Filter = func(line string) string {
		ret := fmt.Sprintf("%06d%s", cmd.current, line)
		cmd.current += cmd.step
		return ret
	}
	return cmd
}
