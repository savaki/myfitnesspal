package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

type Options struct {
	Username string
	Password string
}

func Opts(c *cli.Context) *Options {
	return &Options{
		Username: c.String("username"),
		Password: c.String("password"),
	}
}

var flags = []cli.Flag{
	cli.StringFlag{"username", "", "myfitnesspal username", "MYFITNESSPAL_USERNAME"},
	cli.StringFlag{"password", "", "myfitnesspal password", "MYFITNESSPAL_PASSWORD"},
}

func main() {
	app := cli.NewApp()
	app.Name = "myfitnesspal"
	app.Usage = "cli interface to myfitnesspal"
	app.Author = "Matt Ho"
	app.Commands = []cli.Command{
		diaryCommand,
	}
	app.Run(os.Args)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
