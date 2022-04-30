package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

func readComplexPageData(pageName string) ([]byte, error) {

	var ba []byte
	response, err := http.Get(pageName)
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

func parseComplexPageData(pageData []byte) {
	z := html.NewTokenizer(bytes.NewReader(pageData))

	var isHeader bool
	var isList bool
	var isLink bool

	var counter int

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			t := z.Token()
			if t.Data == "h2" {
				isHeader = true
			}
			if t.Data == "li" {
				isList = true
			}
			if t.Data == "a" {
				isLink = true
			}
		case html.TextToken:
			t := z.Token()
			if isHeader {
				fmt.Printf("@%s", t.Data)
				fmt.Println()
				counter = 1
				isHeader = false
			}
			if isList && isLink {
				fmt.Printf("\t@%d %s", counter, t.Data)
				fmt.Println()
				counter++
				isList = false
				isLink = false
			}

		}
	}
}
