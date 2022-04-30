package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

func readPageData(pageURL string) ([]byte, error) {
	var ba []byte
	response, err := http.Get(pageURL)
	if err != nil {
		return ba, err
	}

	ba, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return ba, err
	}

	return ba, nil
}

func parsePageData(pageData []byte) {
	z := html.NewTokenizer(bytes.NewReader(pageData))

	var isList bool
	var isLink bool

	counter := 1

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			t := z.Token()
			if t.Data == "li" {
				isList = true
			}
			if t.Data == "a" {
				isLink = true
			}
		case html.TextToken:
			t := z.Token()
			if isList && isLink {
				fmt.Printf("@%d %s", counter, t.Data)
				fmt.Println()
				counter++
			}
			isList = false
			isLink = false
		}
	}
}
