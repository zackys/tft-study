package cnv

import (
)

type Rmln struct {
	FilterCmd

	col int
}

func NewRmln(col int) Rmln {
	cmd := Rmln{
		col: col,
	}

	cmd.Filter = func(line string) string {
		ret := ""
		if len(line) > cmd.col {
			ret = line[cmd.col:]
		}
		return ret
	}
	return cmd
}
