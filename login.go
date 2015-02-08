package myfitnesspal

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/publicsuffix"
)

func login(username, password string) (*http.Client, error) {
	client, err := newHttpClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.Get(LoginUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var authenticityToken string
	var utf8 string

	doc.Find("input[name=authenticity_token]").First().Each(func(i int, s *goquery.Selection) {
		authenticityToken, _ = s.Attr("value")
	})

	doc.Find("input[name=utf8]").First().Each(func(i int, s *goquery.Selection) {
		utf8, _ = s.Attr("value")
	})

	params := url.Values{}
	params.Set("utf8", utf8)
	params.Set("authenticity_token", authenticityToken)
	params.Set("username", username)
	params.Set("password", password)
	params.Set("remember_me", "1")
	fmt.Println(params.Encode())
	resp, err = client.Post(LoginUrl, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !isLoggedIn(resp.Body) {
		return nil, ErrNotLoggedIn
	}

	return client, nil
}

func isLoggedIn(r io.Reader) bool {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return false
	}

	var loggedIn bool
	doc.Find(".user-2").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Attr("title")
		loggedIn = len(title) > 0
	})

	return loggedIn
}

func newHttpClient() (*http.Client, error) {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		return nil, err
	}

	return &http.Client{Jar: jar}, nil
}
