package collector

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// TODO: handle url from UK https://www.indeed.co.uk/jobs
// TODO: handle url from Portugal https://pt.indeed.com/ofertas
// TODO: handle url from USA https://www.indeed.com/jobs

var baseURL = "https://fr.indeed.com/jobs"

// StartCollector collects data needed
func StartCollector(term string, location string) string {

	var jobCount string

	queryUrl := fmt.Sprintf("?q=%v&l=%v&radius=0", term, location)

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
