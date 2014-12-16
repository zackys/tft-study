package cnv

import (
	"bufio"
	"os"
)

type End struct {
	fname string
	//fout  *os.File
}

func (cmd End) GoEnd(in <-chan string) <-chan string{
	out := make(chan string)
	go func() {
		var fout *os.File
		var err error

		//println(3)
		if len(cmd.fname) < 1 {
			fout = os.Stdout
			//println(4)
		} else {
			fout, err = os.Create(cmd.fname)
			if err != nil {
				panic(err)
			}
			defer fout.Close()
		}

		//println(5)

		writer := bufio.NewWriter(fout)
		for str := range in {
			//println(str)
			writer.WriteString(str)
		}
		writer.Flush()
		out <- "finished."
		//println(6)
	}()

	return out
}

func NewEnd(fname string) End {
	cmd := End{}
	//cmd.fname = "./test2.txt"
	cmd.fname = fname
	return cmd
}
