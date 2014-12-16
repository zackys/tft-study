package cnv

import (
	"fmt"
)

type L struct {
	FilterCmd

	ls string

	cnt int
}

func newL(ls string) L {
	cmd := L{
		ls: ls,

//		FilterCmd: FilterCmd{
//			Filter: func(line string) string {
//				ret := fmt.Sprintf("%s%s", line, ls)
//				return ret
//			},
//		},
	}

	cmd.Filter = func(line string) string {
		//ret := fmt.Sprintf("%s%s", line, cmd.ls)
		var ret string
		if cmd.cnt == 0 {
			ret = fmt.Sprintf("%s", line)
			cmd.cnt++
		} else {
			ret = fmt.Sprintf("%s%s", cmd.ls, line)
		}
		return ret
	}

	return cmd
}

func NewLw() L {
	return newL("\r\n")
}

func NewLu() L {
	return newL("\n")
}

func NewLm() L {
	return newL("\r")
}
