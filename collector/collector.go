package collector

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// TODO: handle url from UK https://www.indeed.co.uk/jobs
// TODO: handle url from Portugal https://pt.indeed.com/ofertas
// TODO: handle url from USA https://www.indeed.com/jobs

// StartCollector collects data needed
func StartCollector(term string, location string, country string) string {

	var baseURL string
	var removeIndex int

	countryReceiver := strings.ToLower(country)

	switch countryReceiver {
	case "fr":
		baseURL = "https://fr.indeed.com/jobs"
		removeIndex = len("emplois")
	case "uk":
		baseURL = "https://www.indeed.co.uk/jobs"
		removeIndex = len("jobs")
	case "pt":
		baseURL = "https://pt.indeed.com/ofertas"
		removeIndex = len("ofertas")
	case "usa":
		baseURL = "https://www.indeed.com/jobs"
		removeIndex = len("jobs")
	}

	var jobCount string

	queryUrl := fmt.Sprintf("?q=%v&l=%v&radius=0", term, location)

	collector := colly.NewCollector(
		colly.AllowedDomains("www.indeed.com", "indeed.com", "fr.indeed.com", "pt.indeed.com", "www.indeed.co.uk", "indeed.co.uk"),
	)

	collector.OnHTML("#searchCountPages", func(element *colly.HTMLElement) {
		e := element.Text
		str := strings.TrimSpace(e)
		strLen := len(str)
		jobCount = str[9 : strLen-removeIndex]
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting : ", request.URL.String())
	})

	collector.Visit(baseURL + queryUrl)

	return jobCount
}
