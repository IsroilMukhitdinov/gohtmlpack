package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

func readWebData(url string) ([]byte, error) {
	var ba []byte
	response, err := http.Get(url)
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

func parseWebData(ba []byte) {
	z := html.NewTokenizer(bytes.NewReader(ba))

	var isHeader2 bool
	counter := 1

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			t := z.Token()
			isHeader2 = t.Data == "h2"
		case html.TextToken:
			t := z.Token()
			if isHeader2 {
				fmt.Printf("[%d] %s", counter, t.Data)
				fmt.Println()
				counter++
			}
			isHeader2 = false
		}
	}
}
