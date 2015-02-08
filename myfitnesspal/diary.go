package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"github.com/savaki/myfitnesspal"
)

var diaryCommand = cli.Command{
	Name:  "food-diary",
	Usage: "retrieve food diary information",
	Flags: append(flags, []cli.Flag{
		cli.StringFlag{"date", time.Now().Format(myfitnesspal.DateFormat), "the date to retrieve YYYY-MM-DD e.g. 2015-02-01", ""},
	}...),
	Action: diaryAction,
}

func diaryAction(c *cli.Context) {
	opts := Opts(c)

	client, err := myfitnesspal.New(opts.Username, opts.Password)
	check(err)

	timeString := c.String("date")
	date, err := time.Parse(myfitnesspal.DateFormat, timeString)
	check(err)

	entry, err := client.FoodDiary(date)
	check(err)

	data, err := json.MarshalIndent(entry, "", "  ")
	check(err)

	fmt.Println(string(data))
}
