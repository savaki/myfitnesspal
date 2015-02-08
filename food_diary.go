package myfitnesspal

import (
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) FoodDiary(date time.Time) (*DiaryEntry, error) {
	// set the desired date
	dateString := date.Format(DateFormat)
	params := url.Values{}
	params.Set("date", dateString)

	// construct the request url
	uri, err := url.Parse(FoodDiaryUrl + c.username)
	if err != nil {
		return nil, err
	}
	uri.RawQuery = params.Encode()

	// execute the request
	resp, err := c.client.Get(uri.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return parseFoodDiary(resp.Body)
}

func parseFoodDiary(r io.Reader) (*DiaryEntry, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	entries := MacrosArray{}
	var section string = ""
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")
		if class == "meal_header" {
			section = findCellTexts(s)[0]
			return
		} else if class == "total" {
			section = "Totals"
		}

		if entry, err := parseMacros(s); err == nil {
			// bottom is the summary row, we don't want this
			// label is blank for filler columns that don't contain data
			if class != "bottom" && entry.Label != "" {
				entry.Section = section
				entries = append(entries, entry)
			}
		}
	})

	return &DiaryEntry{
		Breakfast: entries.FindAll("Breakfast"),
		Lunch:     entries.FindAll("Lunch"),
		Dinner:    entries.FindAll("Dinner"),
		Snacks:    entries.FindAll("Snacks"),
		Totals:    entries.Find("Totals", "Totals"),
		Goal:      entries.Find("Totals", "Your Daily Goal"),
		Remaining: entries.Find("Totals", "Remaining"),
	}, nil
}

func findCellTexts(s *goquery.Selection) []string {
	return s.Find("td").Map(func(i int, s *goquery.Selection) string {
		return strings.TrimSpace(s.Text())
	})
}

func parseMacros(s *goquery.Selection) (macros *Macros, err error) {
	values := findCellTexts(s)

	atoi := func(str string) int {
		strWithoutCommas := strings.Replace(str, ",", "", -1)
		i, e := strconv.Atoi(strWithoutCommas)
		if err != nil {
			err = e
		}
		return i
	}

	if len(values) < 7 {
		err = fmt.Errorf("no macros found")
		return
	}

	macros = &Macros{
		Label:    values[0],
		Calories: atoi(values[1]),
		Carbs:    atoi(values[2]),
		Fat:      atoi(values[3]),
		Protein:  atoi(values[4]),
		Sodium:   atoi(values[5]),
		Sugar:    atoi(values[6]),
	}

	return
}
