package main

import (
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

func readHtmlFromFile(fileName string) (string, error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func parse(text string) []string {
	z := html.NewTokenizer(strings.NewReader(text))

	var data []string
	var isList bool

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return data
		case html.StartTagToken:
			t := z.Token()
			isList = t.Data == "li"
		case html.TextToken:
			t := z.Token()
			if isList {
				data = append(data, t.Data)
			}
			isList = false
		}

	}
}
