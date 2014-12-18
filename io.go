package main

import (
	"bufio"
	"io"
	"os"
)

type LS string

const (
	L0 LS = ""
	Lu LS = "\n"
	Lm LS = "\r"
	Lw LS = "\r\n"
)

var AssmptDelim byte = '\n'
var ls = Lu
var lsLen int

func checkLs(bytes []byte) {
	switch bytesLen := len(bytes); {
	case bytesLen >= 2:
		if bytes[len(bytes)-2] == '\r' {
			println("delim = '\\r\\n'")
			ls = Lw
		} else {
			println("delim = '\\n'")
			ls = Lu
		}
	case bytesLen == 1:
		println("delim = '\\n'")
		ls = Lu
	case bytesLen == 0:
		ls = L0
	}

	lsLen = len(string(ls))
}

type Input struct {
	fname string
	fin   *os.File
}

func (cmd *Input) GoStart() <-chan string {
	out := make(chan string)
	go func() {
		var fin *os.File
		var err error

		println("[", cmd.fname, "]")
		if len(cmd.fname) < 1 {
			fin = os.Stdin
			println(0)
		} else {
			fin, err = os.Open(cmd.fname)
			if err != nil {
				panic(err)
			}
			defer fin.Close()
		}

//		scanner := bufio.NewScanner(fin)
//		for scanner.Scan() {
//			line := scanner.Text()
//			//println(line)
//			out <- line
//		}
		reader := bufio.NewReader(fin)
		delim := AssmptDelim

		bytes, err := reader.ReadBytes(delim)
		//println("**")
		if err == nil {
			//println("***")
//			switch bytesLen := len(bytes); {
//			case bytesLen >= 2:
//				if bytes[len(bytes)-2] == '\r' {
//					println("delim = '\\r\\n'")
//					delimValue = Lw
//				} else {
//					println("delim = '\\n'")
//					delimValue = Lu
//				}
//			case bytesLen == 1:
//				println("delim = '\\n'")
//				delimValue = Lu
//			case bytesLen == 0:
//				delimValue = L0
//			}
			checkLs(bytes)

			for err == nil {
				out <- string(bytes)

				bytes, err = reader.ReadBytes(delim)
			}

			if err != io.EOF {
				panic(err)
			} else {
				if len(bytes) > 0 {
					out <- string(bytes)
				}
			}

			close(out)
		} else {
			panic(err)
		}
		//println(2)

	}()
	return out
}

func NewInput(fname string) *Input {
	cmd := &Input{}
	cmd.fname = fname
	return cmd
}

////

type L struct {
	Filter

	ls string
	format string
}
func trimLs(line string) (string, bool) {
	bytes := []byte(line)
	lsExists := true
	if n := len(bytes); n >= lsLen {
		line = string(bytes[:(n-lsLen)])
	} else {
		lsExists = false
	}

	return line, lsExists
}

func newL(ls string) *L {
	fltr := &L{
		ls: ls,
	}

	fltr.Apply = func(line string) string {
		var ret string
		trimedLine, lsExists := trimLs(line)
		if lsExists {
			ret = trimedLine + fltr.ls
		} else {
			ret = trimedLine
		}
		return ret
	}

	return fltr
}

func NewLw() *L {
	return newL("\r\n")
}

func NewLu() *L {
	return newL("\n")
}

func NewLm() *L {
	return newL("\r")
}

type Output struct {
	fname string
	//fout  *os.File
}

func (cmd *Output) GoEnd(in <-chan string) <-chan string{
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

func NewOutput(fname string) *Output {
	cmd := &Output{}
	//cmd.fname = "./test2.txt"
	cmd.fname = fname
	return cmd
}