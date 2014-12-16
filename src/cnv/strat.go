package cnv

import (
	"bufio"
	"os"
	"io"
)

type Start struct {
	fname string
	fin   *os.File
}

var DelimType int = Lu
var DelimByte byte = '\n'

const (
	L0 int = 0
	Lu int = 1
	Lm int = 2
	Lw int = 3
)

func (cmd Start) GoStart() <-chan string {
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

		println(1)
		//		scanner := bufio.NewScanner(fin)
		//		for scanner.Scan() {
		//			line := scanner.Text()
		//			//println(line)
		//			out <- line
		//		}
		reader := bufio.NewReader(fin)
		delim := DelimByte
		delimLen := 0

		var strCnt int
		bytes, err := reader.ReadBytes(delim)
		println("**")
		if err == nil {
			println("***")
			switch bytesLen := len(bytes); {
			case bytesLen >= 2:
				if bytes[len(bytes)-2] == '\r' {
					println("delim = '\\r\\n'")
					DelimType = Lw
					delimLen = 2
					strCnt = bytesLen - delimLen
				} else {
					println("delim = '\\n'")
					DelimType = Lu
					delimLen = 1
					strCnt = bytesLen - delimLen
				}
			case bytesLen == 1:
				println("delim = '\\n'")
				DelimType = Lu
				delimLen = 1
				strCnt = bytesLen - delimLen
			case bytesLen == 0:
				strCnt = 0
				DelimType = L0
			}

			for err == nil {
				out <- string(bytes[:strCnt])

				bytes, err = reader.ReadBytes(delim)
				strCnt = len(bytes) - delimLen
				if (strCnt < 0) {
					strCnt = 0
				}
			}

			if err != io.EOF {
				panic(err)
			} else {
				out <- string(bytes[:strCnt])
			}

			close(out)
		} else {
			panic(err)
		}
		println(2)

	}()
	return out
}

func NewStart(fname string) Start {
	cmd := Start{}
	cmd.fname = fname
	return cmd
}
