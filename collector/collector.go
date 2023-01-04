package collector

import (
	"count-jobs/models"
	"count-jobs/utils"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// StartCollector collects data needed
// TODO: refactor
func StartCollector(term string, location string, country string) []byte {

	var baseURL string
	var removeIndex int
	var startIndex int

	countryReceiver := strings.ToLower(country)

	switch countryReceiver {
	case "fr":
		baseURL = "https://fr.indeed.com/jobs"
		removeIndex = len("emplois")
		startIndex = 9
	case "uk":
		baseURL = "https://uk.indeed.com/jobs"
		removeIndex = len("jobs")
		startIndex = 9
	case "pt":
		baseURL = "https://pt.indeed.com/ofertas"
		removeIndex = len("ofertas")
		startIndex = 12
	case "usa":
		baseURL = "https://www.indeed.com/jobs"
		removeIndex = len("jobs")
		startIndex = 9
	}

	var jobCount string

	queryURL := fmt.Sprintf("?q=%v&l=%v&radius=0", term, location)

	collector := colly.NewCollector(
		colly.AllowedDomains("www.indeed.com", "indeed.com", "fr.indeed.com", "https://fr.indeed.com", "pt.indeed.com", "https://pt.indeed.com", "https://uk.indeed.com", "uk.indeed.com", "www.indeed.co.uk", "indeed.co.uk"),
	)

	collector.OnHTML("#searchCountPages", func(element *colly.HTMLElement) {
		e := element.Text
		str := strings.TrimSpace(e)
		strLen := len(str)
		count := str[startIndex : strLen-removeIndex]
		jobCount = strings.TrimSpace(count)
	})

	// Set error handler
	collector.OnError(func(r *colly.Response, err error) {
		utils.Logger.Error("collector", err)
	})

	collector.OnRequest(func(request *colly.Request) {
		utils.Logger.Infof("Visiting %s ", request.URL.String())
	})

	collector.Visit(baseURL + queryURL)

	if jobCount == "" {
		utils.Logger.Error("There is no positions for this job 🙁")
	}

	j := models.Job{
		Tech:     term,
		Count:    jobCount,
		Location: location,
	}

	job, err := json.Marshal(j)
	if err != nil {
		utils.Logger.Error(err)
	}

	return job
}
