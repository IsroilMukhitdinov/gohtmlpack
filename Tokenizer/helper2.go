package main

import (
	// "fmt"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func readPage(pageName string) (string, error) {
	response, err := http.Get(pageName)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func parseAndShow(text string) {
	z := html.NewTokenizer(strings.NewReader(text))

	var isTableData bool
	var counter int

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			t := z.Token()
			if t.Data == "td" {
				isTableData = true
			}
		case html.TextToken:
			t := z.Token()
			if isTableData {
				fmt.Printf("%s ", t.Data)
				counter++
			}
			if isTableData && counter%3 == 0 {
				fmt.Println()
			}
			isTableData = false
		}
	}
}
