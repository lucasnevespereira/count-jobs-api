package collector

import (
	"fmt"

	"github.com/gocolly/colly"
)

var baseURL = "https://www.indeed.com/"

// StartCollector collects data needed
func StartCollector() {

	location := "Paris"
	term := "PHP"

	queryUrl := fmt.Sprintf("jobs?q=%v&l=%v+%2875%29&radius=0", term, location)

	fmt.Println("searching: ", queryUrl)

	collector := colly.NewCollector(
		colly.AllowedDomains("www.indeed.com", "indeed.com"),
	)

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting : ", request.URL.String())
	})

	collector.Visit(baseURL + queryUrl)

}
