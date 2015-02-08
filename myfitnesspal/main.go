package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/savaki/myfitnesspal"
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

func main() {
	app := cli.NewApp()
	app.Name = "myfitnesspal"
	app.Usage = "cli interface to myfitnesspal"
	app.Author = "Matt Ho"
	app.Action = Run
	app.Flags = []cli.Flag{
		cli.StringFlag{"username", "", "myfitnesspal username", "MYFITNESSPAL_USERNAME"},
		cli.StringFlag{"password", "", "myfitnesspal password", "MYFITNESSPAL_PASSWORD"},
	}
	app.Run(os.Args)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Run(c *cli.Context) {
	opts := Opts(c)

	client, err := myfitnesspal.New(opts.Username, opts.Password)
	check(err)

	entry, err := client.FoodDiary(time.Now())
	check(err)

	data, err := json.MarshalIndent(entry, "", "  ")
	check(err)

	fmt.Println(string(data))
}
