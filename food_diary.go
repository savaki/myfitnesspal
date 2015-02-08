package myfitnesspal

import (
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

	macros := map[string]*Macros{}
	doc.Find(".total").Each(func(i int, s *goquery.Selection) {
		if label, entry, err := parseMacros(s); err == nil {
			macros[label] = entry
		}
	})

	return &DiaryEntry{
		Totals:    macros["Totals"],
		Goal:      macros["Your Daily Goal"],
		Remaining: macros["Remaining"],
	}, nil
}

func parseMacros(s *goquery.Selection) (label string, entry *Macros, err error) {
	values := s.Find("td").Map(func(i int, s *goquery.Selection) string {
		return strings.TrimSpace(s.Text())
	})

	atoi := func(str string) int {
		strWithoutCommas := strings.Replace(str, ",", "", -1)
		i, e := strconv.Atoi(strWithoutCommas)
		if err != nil {
			err = e
		}
		return i
	}

	label = values[0]
	entry = &Macros{
		Calories: atoi(values[1]),
		Carbs:    atoi(values[2]),
		Fat:      atoi(values[3]),
		Protein:  atoi(values[4]),
		Sodium:   atoi(values[5]),
		Sugar:    atoi(values[6]),
	}

	return
}
