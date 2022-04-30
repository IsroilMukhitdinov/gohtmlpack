package main

import "log"

func main() {
	// fileName := "index.html"
	// text, err := readHtmlFromFile(fileName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data := parse(text)
	// fmt.Println(data)

	// webPage := "http://webcode.me/countries.html"
	// text, err = readPage(webPage)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// parseAndShow(text)

	// url := "https://www.alexedwards.net/blog"
	// data, err := readWebData(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// parseWebData(data)

	// pageURL := "https://www.alexedwards.net/blog"
	// pageData, err := readPageData(pageURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// parsePageData(pageData)

	pageName := "https://www.alexedwards.net/blog"
	complexPageData, err := readComplexPageData(pageName)
	if err != nil {
		log.Fatal(err)
	}
	parseComplexPageData(complexPageData)

}
