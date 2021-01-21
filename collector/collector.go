package collector

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var baseURL = "https://fr.indeed.com/"

// StartCollector collects data needed
func StartCollector(term string, location string) string {

	var jobCount string

	queryUrl := fmt.Sprintf("jobs?q=%v&l=%v&radius=0", term, location)

	fmt.Println("searching: ", queryUrl)

	collector := colly.NewCollector(
		colly.AllowedDomains("www.indeed.com", "indeed.com", "fr.indeed.com"),
	)

	collector.OnHTML("#searchCountPages", func(element *colly.HTMLElement) {
		e := element.Text
		str := strings.TrimSpace(e)
		strLen := len(str)
		jobCount = str[9 : strLen-7]
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting : ", request.URL.String())
	})

	collector.Visit(baseURL + queryUrl)

	return jobCount
}
