package main

import (
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

var filters map[string]GoFilterer = map[string]GoFilterer{}

var aa []*cli.App = []*cli.App{
	func() *cli.App {
		addln := cli.NewApp()
		addln.Name = "addln"
		addln.Usage = "filter for adding LineNumber"
		addln.Flags = []cli.Flag{
			cli.IntFlag{
				Name:  "start",
				Value: 100000,
				Usage: "start of LineNumper",
			},
			cli.IntFlag{
				Name:  "step",
				Value: 10,
				Usage: "step of LineNumber",
			},
		}
		addln.Action = func(c *cli.Context) {
			start := c.Int("start")
			step := c.Int("step")
			filters["addln"] = NewAddlnFilter(start, step)
		}

		return addln
	}(),

	func() *cli.App {
		rmln := cli.NewApp()
		rmln.Name = "rmln"
		rmln.Usage = "filter as adding LineNumber"
		rmln.Flags = []cli.Flag{
			cli.IntFlag{
				Name:  "col",
				Value: 6,
				Usage: "language for the greeting",
			},
		}
		rmln.Action = func(c *cli.Context) {
			col := c.Int("col")
			filters["rmln"] = NewRmlnFilter(col)
		}

		return rmln
	}(),
}

func main() {

	subApps := map[string]*cli.App{}
	for _, subApp := range aa {
		subApps[subApp.Name] = subApp
	}

	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "filter, f",
			Value: "",
			Usage: "language for the greeting",
		},
		cli.BoolFlag{
			Name:  "j",
			Usage: "出力コード：SJIS",
		},
		cli.BoolFlag{
			Name:  "J",
			Usage: "入力コード：SJIS",
		},
		cli.BoolFlag{
			Name:  "Lu",
			Usage: "出力時改行コード：Unix(LF)",
		},
		cli.BoolFlag{
			Name:  "Lw",
			Usage: "出力時改行コード：Windows(CRLF)",
		},
		cli.BoolFlag{
			Name:  "Lm",
			Usage: "出力時改行コード：Mac(CR)",
		},
	}
	app.Action = func(c *cli.Context) {

		var infilename, outfilename string
		if len(c.Args()) > 0 {
			infilename = c.Args()[0]
		}
		if len(c.Args()) > 1 {
			outfilename = c.Args()[1]
		}

		var compose <-chan string
		compose = NewInput(infilename).GoStart()

		if c.Bool("J") {
			compose = NewS2W8().GoFilter(compose)
		}

		ppp := c.String("filter")
		pp := strings.Split(ppp, "|")
		for _, f := range pp {
			ff := strings.Trim(f, " ")

			println("filter: ", ff)
			subCmd := strings.Split(ff, " ")[0]
			subApp, ok := subApps[subCmd]
			if ok {
				subApp.Run(strings.Split(ff, " "))
			}
		}

		for _, filter := range filters {
			compose = filter.GoFilter(compose)
		}

		if c.Bool("Lu") {
			println("Lu!")
			compose = NewLu().GoFilter(compose)
		} else if c.Bool("Lw") {
			println("Lw!")
			compose = NewLw().GoFilter(compose)
		} else if c.Bool("Lm") {
			println("Lw!")
			compose = NewLw().GoFilter(compose)
		}

		if c.Bool("j") {
			println("j!")
			compose = NewW82S().GoFilter(compose)
		}

		compose = NewOutput(outfilename).GoEnd(compose)
		<-compose
	}

	app.Run(os.Args)

}
